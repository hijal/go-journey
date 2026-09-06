package main

import (
	"fmt"
	"strings"
)

func sendAlert(recipient string, channels ...string) string {
	if len(channels) == 0 {
		channels = []string{"email"}
	}

	return fmt.Sprintf("alert -> %s via %s", recipient, strings.Join(channels, ", "))
}

func main() {
	fmt.Println(sendAlert("oncall@team.dev", "sms", "slack"))
	fmt.Println(sendAlert("oncall@team.dev"))

	allChannels := []string{"pagerduty", "sms", "slack", "email"}
	fmt.Println(sendAlert("oncall@team.dev", allChannels...))
}
