package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/fatih/color"
)

type stdout struct {
	resultSet
}

func (out *stdout) print(...io.Writer) {
	for i, s := range out.searches {
		if i > 0 {
			fmt.Println()
		}

		dash := strings.Repeat("-", (20 - len(s.Keyword)))
		color.Red(s.Keyword + " " + dash)

		for _, e := range s.Entries {
			fmt.Printf("%-13s %s  ", e.Time, e.Title)
			color.Green(e.Price)
		}
	}
}
