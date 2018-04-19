package zookeeper

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

	record := agency.Get(AgentName)
	if record == nil {
		err = fmt.Errorf("Can't find agent '%s'", AgentName)
		return
	}

	agentPath := agency.AgentPath(AgentName)
	settings := record.Settings.(map[string]interface{})

	sh := fmt.Sprintf("%s/%s/bin/zookeeper-server-start.sh",
		agentPath, settings["installer_name"].(string))
	config := fmt.Sprintf("%s/%s",
		agentPath, settings["config_file"].(string))

	err = execkit.Run(nil, exec.Command(sh, "-daemon", config))
	return

}
