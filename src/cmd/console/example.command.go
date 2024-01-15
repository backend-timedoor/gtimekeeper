package console

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

type ExampleCommand struct{}

func (m *ExampleCommand) Signature() string {
	return "example"
}

func (m *ExampleCommand) Flags() []cli.Flag {
	return []cli.Flag{}
}

func (m *ExampleCommand) Handle(c *cli.Context) (err error) {
	fmt.Println("Hello World From Example Command!")

	return nil
}