package admin

type ModuleAdmin struct{}

func (m *ModuleAdmin) AdminGrpc() []any {
	return []any{
		// &greet.HandlerGreet{},
	}
}
