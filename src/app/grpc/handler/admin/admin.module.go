package admin

import "gtimekeeper/src/app/grpc/handler/admin/greet"

type ModuleAdmin struct{}

func (m *ModuleAdmin) AdminGrpc() []any {
	return []any{
		&greet.HandlerGreet{},
	}
}
