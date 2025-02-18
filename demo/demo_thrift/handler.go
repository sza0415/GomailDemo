package main

import (
	"context"
	"github.com/sza0415/GomailDemo/demo/demo_thirft/biz/service"
	api "github.com/sza0415/GomailDemo/demo/demo_thirft/kitex_gen/api"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the EchoImpl interface.
func (s *EchoImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	resp, err = service.NewEchoService(ctx).Run(req)

	return resp, err
}
