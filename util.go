package main

import (
	"strings"

	"github.com/funkygao/gocli"
	"github.com/funkygao/golib/color"
)

func displayRules(ui cli.Ui, rules []byte) {
	for _, rule := range strings.Split(string(rules), "\n") {
		if len(strings.TrimSpace(rule)) == 0 {
			continue
		}

		ui.Outputf("%s %s", color.Green("[ ]"), rule)
	}
}
