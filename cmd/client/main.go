package main

import (
	"github.com/diaohaha/termite/dal"
	proto "github.com/diaohaha/termite/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"log"
)

func main() {
	fmt.Print("hello. client.")
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			dal.Env.MicroRegistry,
		}
	})
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("termite.client"),
		micro.Registry(reg),
	)
	service.Init()

	// Create new greeter client
	client := proto.NewTermiteService("termite", service.Client())
	log.Println("hello handler")
	client.AddWorkFlow(context.TODO(), &proto.AddWorkFlowRequest{WorkflowKey: "video_release", Cid: "1110"})
	//client.WorkStart(context.TODO(), &proto.WorkStartRequest{WorkId:280})
}
