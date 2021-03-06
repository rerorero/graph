package plugin

import (
	"github.com/gogo/protobuf/types"
)

// Message is a message sent from the vertex to another vertex during super-step.
type Message interface{}

// VertexID is id of vertex
type VertexID string

// AggregatableValue is value to be aggregated by aggregator
type AggregatableValue interface{}

// ComputeContext provides information for vertices to process Compute()
type ComputeContext interface {
	SuperStep() uint64
	ReceivedMessages() []Message
	SendMessageTo(dest VertexID, m Message) error
	VoteToHalt()
	GetAggregated(aggregatorName string) (AggregatableValue, bool, error)
	PutAggregatable(aggregatorName string, v AggregatableValue) error
}

// Vertex is abstract of a vertex. thread safe.
type Vertex interface {
	Compute(computeContext ComputeContext) error
	GetID() VertexID
	GetValueAsString() string
}

// Aggregator is Pregel aggregator implemented by user
type Aggregator interface {
	Name() string
	Aggregate(v1 AggregatableValue, v2 AggregatableValue) (AggregatableValue, error)
	MarshalValue(v AggregatableValue) (*types.Any, error)
	UnmarshalValue(pb *types.Any) (AggregatableValue, error)
	ToString(v AggregatableValue) string
}

// Plugin provides an implementation of particular graph computation.
type Plugin interface {
	// TODO: either NewVertex() or NewPartitionVertices() is enough
	// TODO: improve to load vertices. each vertex loading should be concurrently
	NewVertex(id VertexID) (Vertex, error)
	NewPartitionVertices(partitionID uint64, numOfPartitions uint64, register func(v Vertex)) error
	Partition(vertex VertexID, numOfPartitions uint64) (uint64, error)
	MarshalMessage(msg Message) (*types.Any, error)
	UnmarshalMessage(pb *types.Any) (Message, error)
	GetCombiner() func(destination VertexID, messages []Message) ([]Message, error)
	GetAggregators() []Aggregator
}
