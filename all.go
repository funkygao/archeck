package main

import (
	"fmt"
	"strings"

	"github.com/funkygao/gocli"
)

type all struct {
	Ui  cli.Ui
	Cmd string
}

func (s *all) Run(args []string) (exitCode int) {
	for _, c := range []cli.Command{
		&availability{Ui: s.Ui, Cmd: s.Cmd},
		&devops{Ui: s.Ui, Cmd: s.Cmd},
		&resiliency{Ui: s.Ui, Cmd: s.Cmd},
		&scalability{Ui: s.Ui, Cmd: s.Cmd},
	} {
		s.Ui.Output(c.Synopsis())
		c.Run(args)
	}
	return
}

func (s *all) Synopsis() string {
	return "Run all checklists"
}

func (s *all) Help() string {
	help := fmt.Sprintf(`
Usage: %s

    %s		
		`, s.Cmd, s.Synopsis())
	return strings.TrimSpace(help)
}
