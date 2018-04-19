package agent

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/BaritoLog/go-boilerplate/oskit"
)

const (
	EnvBaritoPath     = "BARITO_PATH"
	DefaultBaritoPath = "/opt/barito"

	EnvAgencySource     = "AGENCY_SOURCE"
	DefaultAgencySource = "agency.json"
)

type Agency interface {
	Save() (err error)
	Prepare() (err error)
	Destroy() (err error)
	PrintEnv(w io.Writer)

	Add(name, version string, settings interface{}) *AgentRecord
	Get(name string) *AgentRecord
	Remove(name string) (err error)

	Path() string
	Source() string
	SourcePath() string
	AgentPath(name string) string
	Data() []*AgentRecord
}

type AgentRecord struct {
	Name     string      `json:"name"`
	Version  string      `json:"version"`
	Settings interface{} `json:"settings"`
}

type agency struct {
	data []*AgentRecord
}

// Read Agency
func NewAgency() Agency {

	return &agency{}
}

// Register Program
func (a *agency) Add(name, version string, settings interface{}) (record *AgentRecord) {

	for _, record := range a.data {
		if record.Name == name {
			record.Version = version
			record.Settings = settings
			return record
		}
	}

	record = &AgentRecord{
		Name:     name,
		Version:  version,
		Settings: settings,
	}
	a.data = append(a.data, record)
	return
}

func (a *agency) Remove(name string) (err error) {

	for i, record := range a.data {
		if record.Name == name {
			a.data[i] = a.data[len(a.data)-1]
			a.data[len(a.data)-1] = nil
			a.data = a.data[:len(a.data)-1]
			return
		}
	}

	return
}

// Get Program
func (a *agency) Get(name string) *AgentRecord {
	for _, record := range a.data {
		if record.Name == name {
			return record
		}
	}

	return nil
}

// Path
func (a *agency) Path() string {
	return oskit.Getenv(EnvBaritoPath, DefaultBaritoPath)
}

// Source
func (a *agency) Source() string {
	return oskit.Getenv(EnvAgencySource, DefaultAgencySource)
}

// SourcePath
func (a *agency) SourcePath() string {
	return fmt.Sprintf("%s/%s", a.Path(), a.Source())
}

func (a *agency) AgentPath(name string) string {
	return fmt.Sprintf("%s/%s", a.Path(), name)
}

// Data
func (a *agency) Data() []*AgentRecord {
	return a.data
}

// Prepare
func (a *agency) Prepare() (err error) {

	sourcePath := a.SourcePath()
	path := a.Path()

	raw, err := ioutil.ReadFile(sourcePath)
	if err == nil {
		json.Unmarshal(raw, &a.data)
	} else {
		Info("Can't found agency: %s\n", err.Error())
		Info("Create empty agency")

		os.MkdirAll(path, os.ModePerm)
		err = ioutil.WriteFile(sourcePath, []byte("[]"), 0644)
	}
	return
}

// Save Agency
func (a *agency) Save() (err error) {
	path := a.Path()
	source := a.Source()

	sourcePath := path + "/" + source

	os.MkdirAll(path, os.ModePerm)
	b, _ := json.MarshalIndent(a.Data(), "", "    ")
	err = ioutil.WriteFile(sourcePath, b, 0644)
	return
}

// Destroy Agency
func (a *agency) Destroy() (err error) {
	sourcePath := a.SourcePath()

	Info("Remove agency: %s\n", sourcePath)
	err = os.Remove(sourcePath)

	return
}

func (a *agency) PrintEnv(w io.Writer) {
	fmt.Fprintf(w, "%s=%s\n", EnvBaritoPath, DefaultBaritoPath)
	fmt.Fprintf(w, "%s=%s\n", EnvAgencySource, DefaultAgencySource)
}
