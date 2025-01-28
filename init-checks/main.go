package main

import (
	"fmt"
	"io"

	_ "github.com/aquasecurity/trivy/pkg/iac/rego"
	"github.com/aquasecurity/trivy/pkg/iac/rules"
)

func InitLoop() {
	for _, rule := range rules.GetRegistered() {
		fmt.Fprint(io.Discard, rule.AVDID)
	}
	rules.Reset()
}

func main() {
	InitLoop()
}
