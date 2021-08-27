package main

import (
	"os"
	"strings"

	pwl "github.com/justjanne/powerline-go/powerline"
)

func segmentServerName(p *powerline) []pwl.Segment {
	serverName, _ := os.LookupEnv("SERVER_NAME")
	if serverName == "" {
		return []pwl.Segment{}
	}
	return []pwl.Segment{{
		Name:       "server-name",
		Content:    strings.ToUpper(serverName),
		Foreground: p.theme.ServerNameFg,
		Background: p.theme.ServerNameBg,
	}}
}
