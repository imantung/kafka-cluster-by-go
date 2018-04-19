package general

import (
	"kafka-cluster-by-go/agent"
	"os"

	"github.com/urfave/cli"
)

func Environments(c *cli.Context) (err error) {
	agency := agent.NewAgency()
	agency.PrintEnv(os.Stdout)
	return
}
