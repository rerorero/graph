package worker

import (
	"fmt"
	"sort"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/gogo/protobuf/types"
	"github.com/golang/protobuf/proto"
	"github.com/google/go-cmp/cmp"
	"github.com/rerorero/prerogel/command"
	"github.com/rerorero/prerogel/plugin"
	"github.com/rerorero/prerogel/util"
	"github.com/sirupsen/logrus/hooks/test"
)

func Test_partitionActor_Receive_init(t *testing.T) {
	var mockMux sync.Mutex
	var initializedVertes []string
	var barrierAckCount int32
	var computeAckCount int32
	type fields struct {
		plugin      plugin.Plugin
		vertexProps *actor.Props
	}
	tests := []struct {
		name                  string
		fields                fields
		cmd                   []proto.Message
		wantRespond           []proto.Message
		wantInitializedVertex []string
	}{
		{
			name: "transition from init to superstep",
			fields: fields{
				plugin: &MockedPlugin{
					GetAggregatorsMock: func() []plugin.Aggregator {
						return nil
					},
				},
				vertexProps: actor.PropsFromFunc(func(c actor.Context) {
					mockMux.Lock()
					defer mockMux.Unlock()
					switch cmd := c.Message().(type) {
					case *command.LoadVertex:
						initializedVertes = append(initializedVertes, cmd.VertexId)
						c.Respond(&command.LoadVertexAck{
							VertexId: cmd.VertexId,
						})
					case *command.SuperStepBarrier:
						i := atomic.AddInt32(&barrierAckCount, 1)
						c.Send(c.Parent(), &command.SuperStepBarrierAck{VertexId: initializedVertes[i-1]})
					}
				}),
			},
			cmd: []proto.Message{
				&command.InitPartition{
					PartitionId: 123,
				},
				&command.LoadVertex{VertexId: "test1"},
				&command.LoadVertex{VertexId: "test2"},
				&command.LoadVertex{VertexId: "test3"},
				&command.SuperStepBarrier{},
			},
			wantRespond: []proto.Message{
				&command.InitPartitionAck{PartitionId: 123},
				&command.LoadVertexAck{VertexId: "test1"},
				&command.LoadVertexAck{VertexId: "test2"},
				&command.LoadVertexAck{VertexId: "test3"},
				&command.SuperStepBarrierPartitionAck{PartitionId: 123},
			},
			wantInitializedVertex: []string{"test1", "test2", "test3"},
		},
		{
			name: "superstep, compute, Ack",
			fields: fields{
				plugin: &MockedPlugin{
					GetAggregatorsMock: func() []plugin.Aggregator {
						return nil
					},
				},
				vertexProps: actor.PropsFromFunc(func(c actor.Context) {
					mockMux.Lock()
					defer mockMux.Unlock()
					switch cmd := c.Message().(type) {
					case *command.LoadVertex:
						initializedVertes = append(initializedVertes, cmd.VertexId)
						c.Respond(&command.LoadVertexAck{VertexId: cmd.VertexId})
					case *command.SuperStepBarrier:
						c.Send(c.Parent(), &command.SuperStepBarrierAck{VertexId: initializedVertes[barrierAckCount]})
						barrierAckCount++
					case *command.Compute:
						c.Send(c.Parent(), &command.ComputeAck{VertexId: initializedVertes[computeAckCount]})
						computeAckCount++
					}
				}),
			},
			cmd: []proto.Message{
				&command.InitPartition{
					PartitionId: 123,
				},
				&command.LoadVertex{VertexId: "test1"},
				&command.LoadVertex{VertexId: "test2"},
				&command.LoadVertex{VertexId: "test3"},
				&command.SuperStepBarrier{},
				&command.Compute{SuperStep: 0},
			},
			wantRespond: []proto.Message{
				&command.InitPartitionAck{PartitionId: 123},
				&command.LoadVertexAck{VertexId: "test1"},
				&command.LoadVertexAck{VertexId: "test2"},
				&command.LoadVertexAck{VertexId: "test3"},
				&command.SuperStepBarrierPartitionAck{PartitionId: 123},
				&command.ComputePartitionAck{
					PartitionId:      123,
					AggregatedValues: make(map[string]*types.Any),
				},
			},
			wantInitializedVertex: []string{"test1", "test2", "test3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initializedVertes = nil
			barrierAckCount = 0
			logger, _ := test.NewNullLogger()
			props := actor.PropsFromProducer(func() actor.Actor {
				return NewPartitionActor(tt.fields.plugin, tt.fields.vertexProps, logger)
			})

			context := actor.EmptyRootContext
			proxy := util.NewActorProxy(context, props, nil)

			for i, cmd := range tt.cmd {
				res, err := proxy.SendAndAwait(context, cmd, tt.wantRespond[i], time.Second)
				if (err != nil) != (tt.wantRespond[i] == nil) {
					t.Fatalf("i=%d: %v Ack %d", i, err, barrierAckCount)
				}
				if diff := cmp.Diff(tt.wantRespond[i], res); diff != "" {
					t.Errorf("i=%d: unexpected respond: %s", i, diff)
				}
			}
			sort.Strings(initializedVertes)
			sort.Strings(tt.wantInitializedVertex)
			if diff := cmp.Diff(initializedVertes, tt.wantInitializedVertex); diff != "" {
				t.Errorf("initializedVertex: %s", diff)
			}
			if int(barrierAckCount) != len(initializedVertes) {
				t.Errorf("inconsistent barrier count: %d vs %d", barrierAckCount, len(initializedVertes))
			}
		})
	}
}

func Test_partitionActor_Receive_superstep(t *testing.T) {
	logger, _ := test.NewNullLogger()
	vid := []plugin.VertexID{"test-1", "test-2", "test-3"}
	var called int32
	var messageAckCount int32
	plugin := &MockedPlugin{
		GetAggregatorsMock: func() []plugin.Aggregator {
			return nil
		},
	}
	vertexProps := actor.PropsFromFunc(func(c actor.Context) {
		switch cmd := c.Message().(type) {
		case *command.LoadVertex:
			c.Respond(&command.LoadVertexAck{VertexId: cmd.VertexId})
		case *command.SuperStepBarrier:
			i := atomic.AddInt32(&called, 1)
			c.Send(c.Parent(), &command.SuperStepBarrierAck{VertexId: string(vid[i-1])})
		case *command.Compute:
			i := atomic.AddInt32(&called, 1)
			c.Request(c.Parent(), &command.SuperStepMessage{
				Uuid:         fmt.Sprintf("uuid-%d", i),
				SuperStep:    cmd.SuperStep,
				SrcVertexId:  string(vid[i-1]),
				DestVertexId: "dummy",
				Message:      nil,
			})
		case *command.SuperStepMessageAck:
			i := atomic.AddInt32(&messageAckCount, 1)
			c.Send(c.Parent(), &command.ComputeAck{VertexId: string(vid[i-1]), Halted: false})
		}
	})

	partitionProps := actor.PropsFromProducer(func() actor.Actor {
		return NewPartitionActor(plugin, vertexProps, logger)
	})

	var receivedMessage int32
	computeAckCh := make(chan *command.ComputePartitionAck)
	defer close(computeAckCh)
	context := actor.EmptyRootContext
	proxy := util.NewActorProxy(context, partitionProps, func(ctx actor.Context) {
		switch cmd := ctx.Message().(type) {
		case *command.SuperStepMessage:
			atomic.AddInt32(&receivedMessage, 1)
			ctx.Respond(&command.SuperStepMessageAck{Uuid: cmd.Uuid})
		case *command.ComputePartitionAck:
			computeAckCh <- cmd
		}
	})

	// move state forward
	called = 0
	if _, err := proxy.SendAndAwait(context, &command.InitPartition{
		PartitionId: 123,
	}, &command.InitPartitionAck{}, time.Second); err != nil {
		t.Fatal(err)
	}
	for _, id := range vid {
		if _, err := proxy.SendAndAwait(context, &command.LoadVertex{VertexId: string(id)},
			&command.LoadVertexAck{}, time.Second); err != nil {
			t.Fatal(err)
		}
	}

	// step 0
	called = 0
	if _, err := proxy.SendAndAwait(context, &command.SuperStepBarrier{}, &command.SuperStepBarrierPartitionAck{}, time.Second); err != nil {
		t.Fatal(err)
	}

	called = 0
	messageAckCount = 0
	proxy.Send(context, &command.Compute{SuperStep: 0})
	if ack := <-computeAckCh; ack.PartitionId != 123 {
		t.Fatal("unexpected partition id")
	}

	// step 1
	called = 0
	if _, err := proxy.SendAndAwait(context, &command.SuperStepBarrier{}, &command.SuperStepBarrierPartitionAck{}, time.Second); err != nil {
		t.Fatal(err)
	}

	called = 0
	messageAckCount = 0
	proxy.Send(context, &command.Compute{SuperStep: 1})
	if ack := <-computeAckCh; ack.PartitionId != 123 {
		t.Fatal("unexpected partition id")
	}
}
