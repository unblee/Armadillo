package main

import (
	"log"
	"os"
	"sort"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "Armadillo"
	app.Usage = "Password management CLI tool"
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		{
			Name: "init",
			Usage: "armadillo init <- initialize Armadillo and setting master password.",
			Action: func(c *cli.Context) error {
				log.Printf("Initialize Armadillo.")
				return nil
			},
		},
		{
			Name: "login",
			Usage: "armadillo login <- login with Armadillo.",
			Action: func(c *cli.Context) error {
				log.Printf("Armadillo login")
				return nil
			},
		},
		{
			Name: "logout",
			Usage: "armadillo logout <- logout with Armadillo.",
			Action: func(c *cli.Context) error {
				log.Printf("Armadillo logout")
				return nil
			},
		},
		{
			Name: "create",
			Usage: "armadillo create [site_name] <- setting password for site.",
			Action: func(c *cli.Context) error {
				log.Printf("Setting passowrd for site.")
				return nil
			},
		},
		{
			Name: "update",
			Usage: "armadillo update <- update password.",
			Action: func(c *cli.Context) error {
				log.Printf("Update password.")
				return nil
			},
		},
		{
			Name: "show",
			Usage: "armadillo show <- show password.",
			Action: func(c *cli.Context) error {
				log.Printf("Show password.")
				return nil
			},
		},
	}

	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
