package main

import (
	"fmt"
	"os"

	"github.com/devproje/commando"
	"github.com/gin-gonic/gin"
)

func main() {
	command := commando.NewCommando(os.Args[:1])
	command.Root("daemon", "run file server", func(n *commando.Node) error {
		gin := gin.Default()
		if err := gin.Run(); err != nil {
			return err
		}

		return nil
	})

	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
