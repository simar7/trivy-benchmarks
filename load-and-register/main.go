package main

import (
	"fmt"
	"io"

	"github.com/aquasecurity/trivy/pkg/iac/rego"
	"github.com/aquasecurity/trivy/pkg/iac/rules"
)

func LoadRegisterLoop() {
	rego.LoadAndRegister()

	for _, rule := range rules.GetRegistered() {
		fmt.Fprint(io.Discard, rule.AVDID)
	}

	rules.Reset()
}

func main() {
	LoadRegisterLoop()
}
