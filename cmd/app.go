package main

import (
	"fmt"
	"os"

	"git.wh64.net/devproje/kuma-archive/internal/routes"
	"github.com/devproje/commando"
	"github.com/devproje/commando/option"
	"github.com/devproje/commando/types"
	"github.com/gin-gonic/gin"
)

func main() {
	command := commando.NewCommando(os.Args[1:])
	command.Root("daemon", "run file server", func(n *commando.Node) error {
		debug, err := option.ParseBool(*n.MustGetOpt("debug"), n)
		if err != nil {
			return err
		}

		if !debug {
			gin.SetMode(gin.ReleaseMode)
		}

		gin := gin.Default()
		routes.New(gin)

		if err := gin.Run(); err != nil {
			return err
		}

		return nil
	}, types.OptionData{
		Name:  "debug",
		Desc:  "service debugging mode",
		Short: []string{"d"},
		Type:  types.BOOLEAN,
	})

	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
