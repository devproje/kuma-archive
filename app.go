package main

import (
	"errors"
	"fmt"
	"os"
	"syscall"

	"git.wh64.net/devproje/kuma-archive/config"
	"git.wh64.net/devproje/kuma-archive/internal/routes"
	"git.wh64.net/devproje/kuma-archive/internal/service"
	"github.com/devproje/commando"
	"github.com/devproje/commando/option"
	"github.com/devproje/commando/types"
	"github.com/gin-gonic/gin"
	"golang.org/x/term"
)

var (
	hash    = "unknown"
	branch  = "unknown"
	version = "unknown"
)

func main() {
	command := commando.NewCommando(os.Args[1:])
	cnf := config.Get()

	ver := service.NewVersion(version, branch, hash)
	command.Root("daemon", "run file server", func(n *commando.Node) error {
		fmt.Printf("Kuma Archive %s\n", version)
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

		// init auth module
		service.NewAuthService()

		gin := gin.Default()
		routes.New(gin, ver, apiOnly)

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

	command.Root("version", "show system version info", func(n *commando.Node) error {
		fmt.Printf("Kuma Archive version %s\n", ver.String())
		return nil
	})

	command.ComplexRoot("account", "file server account manager", []commando.Node{
		command.Then("create", "create account", func(n *commando.Node) error {
			var username, password string

			fmt.Print("new username: ")
			if _, err := fmt.Scanln(&username); err != nil {
				return fmt.Errorf("failed to read username: %v", err)
			}

			fmt.Print("new password: ")
			bytePassword, err := term.ReadPassword(int(syscall.Stdin))
			if err != nil {
				return fmt.Errorf("failed to read password: %v", err)
			}
			password = string(bytePassword)
			fmt.Println()

			fmt.Print("type new password one more time: ")
			checkByte, err := term.ReadPassword(int(syscall.Stdin))
			if err != nil {
				return fmt.Errorf("failed to read password: %v", err)
			}
			check := string(checkByte)
			fmt.Println()

			if password != check {
				return errors.New("password check is not correct")
			}

			auth := service.NewAuthService()
			if err := auth.Create(&service.Account{Username: username, Password: password}); err != nil {
				return err
			}

			fmt.Printf("Account for %s created successfully\n", username)
			return nil
		}),
	})

	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
