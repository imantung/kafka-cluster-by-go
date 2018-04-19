package kafka

import (
	"fmt"
	"kafka-cluster-by-go/agent"

	"github.com/BaritoLog/go-boilerplate/execkit"
	"github.com/BaritoLog/go-boilerplate/execkit/linux"
	"github.com/urfave/cli"
)

func Status(c *cli.Context) (err error) {

	agency := agent.NewAgency()
	agency.Prepare()

	record := agency.Get(AgentName)
	if record == nil {
		err = fmt.Errorf("Can't find agent '%s'", AgentName)
		return
	}

	agentPath := agency.AgentPath(AgentName)
	settings := record.Settings.(map[string]interface{})

	pid, err := execkit.Pid(
		"java",
		fmt.Sprintf("%s/%s", agentPath, settings["config_file"]),
	)

	if len(pid) > 0 {
		fmt.Printf("Kafka is started with pid=%s", pid)
	} else {
		fmt.Println("Kafka is stopped")
	}

	return
}

func Pid() (pid []byte, err error) {
	pid, err = linux.Bash(`ps ax | grep -i 'kafka' | grep java | grep -v grep | awk '{print $1}'`).Output()
	return

}
