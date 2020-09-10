package main

import (
	"fmt"
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/think-divergent/twitter-fish-cannon/server"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.IntFlag{Name: "port, p", Usage: "port to listen on", Value: 8080, EnvVar: "PORT"},
	}
	app.ArgsUsage = "STATIC_DIR"
	app.Action = serve

	app.RunAndExitOnError()
}

func runFishCannon() {
	// TODO trying to follow this example here
	// https://github.com/dghubble/go-twitter/blob/master/examples/streaming.go
	// TODO figure out how to get api access credentials on twitter
	// TODO read these from environment variable or commandline arguments?
	config := oauth1.NewConfig("consumerKey", "consumerSecret")
	token := oauth1.NewToken("accessToken", "accessSecret")
	httpClient := config.Client(oauth1.NoContext, token)
}

func serve(ctx *cli.Context) error {
	port := ctx.Int("port")
	if port == 0 {
		fmt.Println("error: Specify port either via `-p` or via PORT env var")
		return ctx.App.Command("help").Run(ctx)
	}

	staticPath := ctx.Args().First()
	if staticPath == "" {
		fmt.Println("error: STATIC_DIR argument required")
		return ctx.App.Command("help").Run(ctx)
	}

	log.Printf("Sending port and static path to server module (%d, %s)", port, staticPath)
	return server.ListenAndServe(port, staticPath)
}
