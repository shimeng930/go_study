package base

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Action = func(c *cli.Context) error {
		fmt.Println("BOOM!")
		return nil
	}
	app.Name = "GoTest"
	app.Usage = "hello"
	app.Version = "1.2"

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}