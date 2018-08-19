package v1

import (
	"github.com/hatajoe/8am/app/interface/rpc/v1.0/protocol"
	"github.com/hatajoe/8am/app/registry"
	"github.com/hatajoe/8am/app/usecase"
	"google.golang.org/grpc"
)

func Apply(server *grpc.Server, ctn *registry.Container) {
	protocol.RegisterUserServiceServer(server, NewUserService(ctn.Resolve("user-usecase").(usecase.UserUsecase)))
}
