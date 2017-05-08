package main

import (
	"fmt"
	"strings"

	"github.com/funkygao/gocli"
)

type tcc struct {
	Ui  cli.Ui
	Cmd string
}

func (s *tcc) Run(args []string) (exitCode int) {
	rules, _ := templateTccBytes()
	displayRules(s.Ui, rules)
	return
}

func (s *tcc) Synopsis() string {
	return "Try/Confirm/Cancel for distributed microservice consistency"
}

func (s *tcc) Help() string {
	help := fmt.Sprintf(`
Usage: %s

    %s		
		`, s.Cmd, s.Synopsis())
	return strings.TrimSpace(help)
}
