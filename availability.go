package main

import (
	"fmt"
	"strings"

	"github.com/funkygao/gocli"
)

type availability struct {
	Ui  cli.Ui
	Cmd string
}

func (s *availability) Run(args []string) (exitCode int) {
	rules, _ := templateAvailabilityBytes()
	displayRules(s.Ui, rules)
	return
}

func (s *availability) Synopsis() string {
	return "Availability Checklist"
}

func (s *availability) Help() string {
	help := fmt.Sprintf(`
Usage: %s

    %s		
		`, s.Cmd, s.Synopsis())
	return strings.TrimSpace(help)
}
