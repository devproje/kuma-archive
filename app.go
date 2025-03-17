package main

import (
	"fmt"
	"os"

	"git.wh64.net/devproje/kuma-archive/config"
	"git.wh64.net/devproje/kuma-archive/internal/routes"
	"github.com/devproje/commando"
	"github.com/devproje/commando/option"
	"github.com/devproje/commando/types"
	"github.com/gin-gonic/gin"
)

var (
	hash    = "unknown"
	version = "unknown"
)

func main() {
	fmt.Printf("Kuma Archive %s-%s\n", version, hash)
	command := commando.NewCommando(os.Args[1:])
	cnf := config.Get()

	command.Root("daemon", "run file server", func(n *commando.Node) error {
		debug, err := option.ParseBool(*n.MustGetOpt("debug"), n)
		if err != nil {
			return err
		}

		apiOnly, err := option.ParseBool(*n.MustGetOpt("api-only"), n)
		if err != nil {
			return err
		}

		if !debug {
			gin.SetMode(gin.ReleaseMode)
		}

		gin := gin.Default()
		routes.New(gin, apiOnly)

		fmt.Fprintf(os.Stdout, "binding server at: http://0.0.0.0:%d\n", cnf.Port)
		if err := gin.Run(fmt.Sprintf(":%d", cnf.Port)); err != nil {
			return err
		}

		return nil
	}, types.OptionData{
		Name:  "debug",
		Desc:  "service debugging mode",
		Short: []string{"d"},
		Type:  types.BOOLEAN,
	}, types.OptionData{
		Name: "api-only",
		Desc: "no serve frontend service",
		Type: types.BOOLEAN,
	})

	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
