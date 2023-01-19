package agent

import (
	"github.com/unravellingtechnologies/unravelling-cloud/cli/lib/health"
	"github.com/unravellingtechnologies/unravelling-cloud/cli/lib/server"
	"github.com/unravellingtechnologies/unravelling-cloud/cli/lib/terraform"
)

func Start(host string, port int, terraformEnabled bool, workingDir string) {
	if terraformEnabled {
		terraform.Prepare(server.Router)
	}

	health.Initialize(server.Router)

	server.StartServer(&server.Opts{Port: port, Host: host})
}
