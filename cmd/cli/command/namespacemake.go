package command

import (
	"fmt"
	"strings"
	"time"

	"github.com/KvrocksLabs/kvrocks_controller/cmd/cli/context"
	"github.com/KvrocksLabs/kvrocks_controller/server/handlers"
	"github.com/KvrocksLabs/kvrocks_controller/util"
	"gopkg.in/urfave/cli.v1"
)

var MakeNsCommand = cli.Command{
	Name:      "mkns",
	Usage:     "make namespcae",
	ArgsUsage: "mkns ${namespace}",
	Action:    mknsAction,
	Description: `
    create namespce
    `,
}

func mknsAction(c *cli.Context) {
	if len(c.Args()) != 1 {
		fmt.Println("mkns only set one param(${namespace})")
		return
	}
	name := c.Args()[0]
	if strings.Contains(name, "/") {
		fmt.Println("namespcae can't contain '/'")
		return
	}
	ctx := context.GetContext()
	if ctx.Location != context.LocationRoot {
		fmt.Println("mkns need return root dir '/'")
		return
	}

	resp, err := util.HttpPost(handlers.GetNamespaceRootURL(ctx.Leader), handlers.CreateNamespaceParam{Namespace: name}, 5*time.Second)
	if HttpResponeException("make namespcae", resp, err) {
		return
	}
}