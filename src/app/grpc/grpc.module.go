package grpc

import (
	"gtimekeeper/src/app/grpc/handler/admin"
)

type ModuleGrpc struct{}

func (m *ModuleGrpc) BootGrpc() []any {
	return []any{
		&admin.ModuleAdmin{},
	}
}
