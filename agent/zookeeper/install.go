package zookeeper

import (
	"fmt"
	"kafka-cluster-by-go/agent"
	"log"
	"os"

	"github.com/BaritoLog/go-boilerplate/execkit"
	"github.com/BaritoLog/go-boilerplate/execkit/linux"
	"github.com/BaritoLog/go-boilerplate/filekit"
	"github.com/urfave/cli"
)

const (
	AgentName = "zookeeper"
)

func Install(c *cli.Context) (err error) {

	agency := agent.NewAgency()
	agency.Prepare()

	log.Printf("Retrieve %s agent setting\n", AgentName)
	set, err := RetrieveSettings()
	if err != nil {
		return
	}

	agentPath := agency.AgentPath(AgentName)
	configPath := agentPath + "/" + set.ConfigFile

	log.Printf("Make agent directory '%s'\n", agentPath)
	os.MkdirAll(agentPath, os.ModePerm)

	log.Printf("Write properties file '%s'\n", configPath)
	err = filekit.WritePropertiesFile(configPath, set.ConfigParam)
	if err != nil {
		return
	}

	log.Printf("(Exec)\n\n")
	tmp := fmt.Sprintf("%s/%s.%s", agentPath, set.InstallerName, set.InstallFileType)
	err = execkit.Run(os.Stdout,
		linux.Download(set.InstallerUrl, tmp),
		linux.ExtractGzip(tmp, agentPath),
		linux.Remove(tmp),
	)
	if err != nil {
		return
	}

	log.Printf("Add '%s' to local agency\n", AgentName)
	agency.Add(AgentName, set.Version, set)
	agency.Save()

	// TODO: register to market

	return
}
