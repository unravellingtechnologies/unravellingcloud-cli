/*
Copyright © 2022 Unravelling Technologies GmbH <unravelling@unravelling.io>
*/
package main

import (
	"github.com/unravellingtechnologies/unravelling-cloud/cli/cmd"
	_ "github.com/unravellingtechnologies/unravelling-cloud/cli/lib/logging"
)

var Version = "v0.1.0"

func main() {
	cmd.Execute()
}
