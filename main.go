package main

import (
	"log"
	"os"

	"kafka-cluster-by-go/agent"
	"kafka-cluster-by-go/agent/general"
	"kafka-cluster-by-go/agent/kafka"
	"kafka-cluster-by-go/agent/zookeeper"

	. "github.com/urfave/cli"
)

const (
	Name    = "Barito Agent"
	Version = "0.1"
)

func main() {

	agent.Verbose = true

	app := App{
		Name:    Name,
		Usage:   "Agent for installation and ops works",
		Version: Version,
		Commands: []Command{
			Command{
				Name:      "kafka",
				ShortName: "k",
				Subcommands: []Command{
					Command{Name: "install", Action: kafka.Install},
					Command{Name: "start", Action: kafka.Start},
					Command{Name: "stop", Action: kafka.Stop},
					Command{Name: "stop", Action: kafka.Status},
				},
			},
			Command{
				Name:      "zookeeper",
				ShortName: "z",
				Subcommands: []Command{
					Command{Name: "install", Action: zookeeper.Install},
					Command{Name: "start", Action: zookeeper.Start},
					Command{Name: "stop", Action: zookeeper.Stop},
					Command{Name: "status", Action: zookeeper.Status},
				},
			},
			Command{Name: "env", Usage: "print environment", Action: general.Environments},
			Command{Name: "ki", Usage: "kafka install", Action: kafka.Install},
			Command{Name: "ks", Usage: "kafka start", Action: kafka.Start},
			Command{Name: "kx", Usage: "kafka stop", Action: kafka.Stop},
			Command{Name: "ka", Usage: "kafka status", Action: kafka.Status},
			Command{Name: "zi", Usage: "zookeeper install", Action: zookeeper.Install},
			Command{Name: "zs", Usage: "zookeeper start", Action: zookeeper.Start},
			Command{Name: "zx", Usage: "zookeeper stop", Action: zookeeper.Stop},
			Command{Name: "za", Usage: "zookeeper status", Action: zookeeper.Status},
		},
	}

	err := app.Run(os.Args)
	FatalIfError(err)
}

func FatalIfError(err error) {
	if err != nil {
		// TODO: using https://github.com/golang/glog for logging
		log.Fatalf("\n\n\n\nFatal Error: %s", err.Error())
	}

}
