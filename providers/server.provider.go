package providers

import (
	"gtimekeeper/src/app/grpc"
	"gtimekeeper/src/app/http"

	"github.com/backend-timedoor/gtimekeeper-framework/app"
	"github.com/backend-timedoor/gtimekeeper-framework/base/contracts"
	"github.com/backend-timedoor/gtimekeeper-framework/base/server"
	"github.com/backend-timedoor/gtimekeeper-framework/base/server/servers"
)

type ServerServiceProvider struct{}

var (
	httpServer = &servers.Http{
		Modules: []any{
			&http.ModuleHttp{},
		},
	}
	grpcServer = &servers.Grpc{
		Modules: []any{
			&grpc.ModuleGrpc{},
		},
	}
)

func (p *ServerServiceProvider) Boot() {
	app.Server.RegisterCustomeValidation([]contracts.Validation{
		//
	})
}

func (p *ServerServiceProvider) Register() {
	s := []contracts.ServerHandle{
		grpcServer,
		httpServer,
	}

	app.Server = server.BootServer(s)
}
