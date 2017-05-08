package main

import (
	"fmt"
	"os"

	"github.com/funkygao/gocli"
)

//go:generate go-bindata -nomemcopy -pkg main template/...

var commands map[string]cli.CommandFactory

func init() {
	ui := &cli.ColoredUi{
		Ui: &cli.BasicUi{
			Writer:      os.Stdout,
			Reader:      os.Stdin,
			ErrorWriter: os.Stderr,
		},
		OutputColor: cli.UiColorNone,
		InfoColor:   cli.UiColorGreen,
		ErrorColor:  cli.UiColorRed,
		WarnColor:   cli.UiColorYellow,
	}
	cmd := os.Args[0]
	commands = map[string]cli.CommandFactory{
		"scalability": func() (cli.Command, error) {
			return &scalability{
				Ui:  ui,
				Cmd: cmd,
			}, nil
		},

		"resiliency": func() (cli.Command, error) {
			return &resiliency{
				Ui:  ui,
				Cmd: cmd,
			}, nil
		},

		"availability": func() (cli.Command, error) {
			return &availability{
				Ui:  ui,
				Cmd: cmd,
			}, nil
		},

		"tcc": func() (cli.Command, error) {
			return &tcc{
				Ui:  ui,
				Cmd: cmd,
			}, nil
		},

		"all": func() (cli.Command, error) {
			return &all{
				Ui:  ui,
				Cmd: cmd,
			}, nil
		},

		"devops": func() (cli.Command, error) {
			return &devops{
				Ui:  ui,
				Cmd: cmd,
			}, nil
		},
	}
}

func main() {
	app := os.Args[0]
	args := os.Args[1:]
	c := cli.NewCLI(app, "archeck (1.0)")
	c.Args = args
	c.Commands = commands
	exitCode, err := c.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	} else if c.IsVersion() {
		os.Exit(0)
	}

	os.Exit(exitCode)
}
