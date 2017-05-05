package main

import (
	"fmt"
	"strings"

	"github.com/funkygao/gocli"
)

type scalability struct {
	Ui  cli.Ui
	Cmd string
}

func (s *scalability) Run(args []string) (exitCode int) {
	rules, _ := templateScalabilityBytes()
	s.Ui.Output(string(rules))
	return
}

func (s *scalability) Synopsis() string {
	return "Scalability Checklist"
}

func (s *scalability) Help() string {
	help := fmt.Sprintf(`
Usage: %s

    %s		
		`, s.Cmd, s.Synopsis())
	return strings.TrimSpace(help)
}
