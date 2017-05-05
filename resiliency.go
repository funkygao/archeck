package main

import (
	"fmt"
	"strings"

	"github.com/funkygao/gocli"
)

type resiliency struct {
	Ui  cli.Ui
	Cmd string
}

func (s *resiliency) Run(args []string) (exitCode int) {
	rules, _ := templateResiliencyBytes()
	s.Ui.Output(string(rules))
	return
}

func (s *resiliency) Synopsis() string {
	return "Resiliency Checklist"
}

func (s *resiliency) Help() string {
	help := fmt.Sprintf(`
Usage: %s

    %s		
		`, s.Cmd, s.Synopsis())
	return strings.TrimSpace(help)
}
