package zookeeper

type Settings struct {
	InstallerName   string                 `json:"installer_name"`
	InstallFileType string                 `json:"installer_file_type"`
	InstallerUrl    string                 `json:"installer_url"`
	Version         string                 `json:"version"`
	ConfigFile      string                 `json:"config_file"`
	ConfigParam     map[string]interface{} `json:"config_param"`
}

func RetrieveSettings() (setting *Settings, err error) {
	// TODO: retrieve from market server
	setting = &Settings{
		InstallerName:   "kafka_2.10-0.10.2.0",
		InstallFileType: "tgz",
		InstallerUrl:    "http://localhost:4000/kafka_2.10-0.10.2.0.tgz",
		Version:         "0.10.2",
		ConfigFile:      "start_zookeeper.properties",
		ConfigParam: map[string]interface{}{
			"dataDir":        "/tmp/zookeeper",
			"clientPort":     2181,
			"maxClientCnxns": 0,
		},
	}
	return
}
