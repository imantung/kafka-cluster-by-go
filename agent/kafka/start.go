package kafka

import (
	"fmt"
	"kafka-cluster-by-go/agent"
	"os/exec"

	"github.com/BaritoLog/go-boilerplate/execkit"
	"github.com/urfave/cli"
)

func Start(c *cli.Context) (err error) {
	agency := agent.NewAgency()
	agency.Prepare()

	agentPath := agency.AgentPath(AgentName)
	settings := agency.Get(AgentName).Settings.(map[string]interface{})

	// configFile := settings["config_file"].(string)

	sh := fmt.Sprintf("%s/%s/bin/kafka-server-start.sh",
		agentPath, settings["installer_name"].(string))
	config := fmt.Sprintf("%s/%s",
		agentPath, settings["config_file"].(string))

	err = execkit.Run(nil, exec.Command(sh, "-daemon", config))
	return

}
