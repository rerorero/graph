package main

import (
	"fmt"
	"github.com/AsynkronIT/protoactor-go/actor"
	"time"
)

type Hello struct{ Who string }
type HelloActor struct{}

func (state *HelloActor) Receive(context actor.Context) {
    switch msg := context.Message().(type) {
    case Hello:
        fmt.Printf("Hello %v\n", msg.Who)
    }
}

func main() {
    context := actor.EmptyRootContext
    props := actor.PropsFromProducer(func() actor.Actor { return &HelloActor{} })
    pid := context.Spawn(props)
    context.Send(pid, Hello{Who: "Roger"})
    time.Sleep(2 * time.Second)
}
