package main

import (
	"fmt"
	"strings"

	"github.com/funkygao/gocli"
)

type devops struct {
	Ui  cli.Ui
	Cmd string
}

func (s *devops) Run(args []string) (exitCode int) {
	rules, _ := templateDevopsBytes()
	displayRules(s.Ui, rules)
	return
}

func (s *devops) Synopsis() string {
	return "DevOps Checklist"
}

func (s *devops) Help() string {
	help := fmt.Sprintf(`
Usage: %s

    %s		
		`, s.Cmd, s.Synopsis())
	return strings.TrimSpace(help)
}
