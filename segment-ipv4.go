package main

import (
	"os"

	pwl "github.com/justjanne/powerline-go/powerline"
)

func segmentIPv4(p *powerline) []pwl.Segment {
	content := ""
	ipv4Public, _ := os.LookupEnv("PUBLIC_IPV4")
	ipv4Local, _ := os.LookupEnv("LOCAL_IPV4")

	if ipv4Public != "" {
		content += ipv4Public

		if ipv4Local != "" {
			content += " // "
		}
	}

	if ipv4Local != "" {
		content += "(" + ipv4Local + ")"
	}

	if content == "" {
		return []pwl.Segment{}
	}

	return []pwl.Segment{{
		Name:       "ipv4",
		Content:    content,
		Foreground: p.theme.IPv4Fg,
		Background: p.theme.IPv4Bg,
	}}
}
