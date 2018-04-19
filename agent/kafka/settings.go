package kafka

type Settings struct {
	InstallerName     string                 `json:"installer_name"`
	InstallerFileType string                 `json:"installer_file_type"`
	InstallerUrl      string                 `json:"installer_url"`
	Version           string                 `json:"version"`
	ConfigFile        string                 `json:"config_file"`
	ConfigParam       map[string]interface{} `json:"config_param"`
}

func RetrieveSettings() (setting *Settings, err error) {
	// TODO: retrieve from market server
	setting = &Settings{
		InstallerName:     "kafka_2.10-0.10.2.0",
		InstallerFileType: "tgz",
		InstallerUrl:      "http://localhost:4000/kafka_2.10-0.10.2.0.tgz",
		Version:           "0.10.2",
		ConfigFile:        "start_kafka.properties",
		ConfigParam: map[string]interface{}{
			"broker.id":                         0,
			"num.network.threads":               3,
			"num.io.threads":                    8,
			"socket.send.buffer.bytes":          102400,
			"socket.receive.buffer.bytes":       102400,
			"socket.request.max.bytes":          104857600,
			"log.dirs":                          "/tmp/kafka-logs",
			"num.partitions":                    1,
			"num.recovery.threads.per.data.dir": 1,
			"log.retention.hours":               168,
			"zookeeper.connect":                 "localhost:2181",
			"zookeeper.connection.timeout.ms":   6000,
		},
	}
	return
}
