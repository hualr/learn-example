package main

import (
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.App{}
	app.Run(os.Args)
}
