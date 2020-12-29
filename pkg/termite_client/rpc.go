package termite_client

import (
	"github.com/diaohaha/termite/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/memory"
	"log"
	"sync"
)

var termiteClientV1Once sync.Once
var termiteClient termite.TermiteService

func init() {
}

func newTestRegistry() registry.Registry {
	r := memory.NewRegistry()
	return r
}

func AddWorkFlow(cid string, project string, vflow string) error {
	service := "Termite"
	endpoint := "Termite.AddWorkFlow"
	address := "172.17.2.76"
	//address := "127.0.0.1"
	port := 17002

	r := newTestRegistry()
	c := client.NewClient(
		client.Registry(r),
	)
	//c := client.DefaultClient
	c.Options().Selector.Init(selector.Registry(r))

	req := c.NewRequest(service, endpoint, &termite.AddWorkFlowRequest{
		Cid:         cid,
		Project:     project,
		WorkflowKey: vflow,
	})

	// test calling remote address
	if err := c.Call(context.Background(), req, nil, client.WithAddress(fmt.Sprintf("%s:%d", address, port))); err != nil {
		log.Print("call with address error", err)
		return err
	}
	return nil
}
