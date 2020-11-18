package main

import (
    "contest-check-in/server"
    "fmt"
    "github.com/urfave/cli/v2"
    "log"
    "os"
    "runtime"
)

func main() {
    main := &cli.App{
        Name:  "WeJudge contest check-in tools",
        Action: func(c *cli.Context) error {
            fmt.Println("WeJudge contest check-in tools")
            fmt.Printf("built: %s(%s)\n", runtime.GOOS, runtime.GOARCH)
            return nil
        },
        Commands: cli.Commands{
            {
                Name:      "server",
                Usage:     "run server",
                Action: func(context *cli.Context) error {
                    server.RunServerWorker("127.0.0.1", 8088, "")
                    return nil
                },
            },
        },
    }
    err := main.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}
