package rpc

import (
	"github.com/hatajoe/8am/app/interface/rpc/v1.0"
	"github.com/hatajoe/8am/app/registry"
	"google.golang.org/grpc"
)

func Apply(server *grpc.Server, ctn *registry.Container) {
	v1.Apply(server, ctn)
}
