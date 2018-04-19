package agent

import (
	"bytes"
	"os"
	"testing"

	. "github.com/BaritoLog/go-boilerplate/testkit"
)

func TestAgency(t *testing.T) {

	os.Setenv(EnvBaritoPath, ".")
	os.Setenv(EnvAgencySource, "agency.json")

	ag1 := &agency{
		data: []*AgentRecord{
			&AgentRecord{Name: "kafka", Version: "0.1.1"},
			&AgentRecord{Name: "elasticsearch", Version: "0.1.1"},
		},
	}

	err := ag1.Save()
	FatalIfError(t, err)

	ag2 := NewAgency()
	err = ag2.Prepare()
	lenData := len(ag2.Data())
	FatalIfError(t, err)
	FatalIf(t, lenData != 2, "agency data should be 2: %d", lenData)

	registered := ag2.Add("right_agent", "1.8", "settings")

	path := ag2.Path()
	right := ag2.Get("right_agent")
	wrong := ag2.Get("wrong_agent")

	FatalIf(t, path != ".", "agency path is wrong: %s", path)
	FatalIf(t, right != registered, "Right agent is not registered: %+v", right)
	FatalIf(t, wrong != nil, "Wrong agent should return nil: %+v", wrong)

	err = ag2.Destroy()
	FatalIfError(t, err)
}

func TestReadAgency_MissingSource(t *testing.T) {
	os.Setenv(EnvBaritoPath, ".")
	os.Setenv(EnvAgencySource, "new_agency.json")

	agency := NewAgency()
	agency.Prepare()

	_, err := os.Stat("./new_agency.json")
	FatalIfError(t, err)

	agency.Destroy()
}

func TestReadAgency_CreateEmptyAgencyFailed(t *testing.T) {
	os.Setenv(EnvBaritoPath, "/etc")
	os.Setenv(EnvAgencySource, "agency.json")

	ag := NewAgency()
	err := ag.Prepare()
	FatalIfWrongError(t, err, "open /etc/agency.json: permission denied")
}

func TestAgency_Register_SameName(t *testing.T) {
	ag := &agency{
		data: []*AgentRecord{
			&AgentRecord{Name: "kafka", Version: "0.1.1"},
		},
	}

	ag.Add("kafka", "0.1.2", "settings")
	rec := ag.Get("kafka")
	FatalIf(t, rec.Version != "0.1.2", "version should be updated: %s", rec.Version)
	FatalIf(t, rec.Settings != "settings", "settings should be updated: %s", rec.Settings)
}

func TestAgency_PrintEnv(t *testing.T) {
	os.Setenv(EnvBaritoPath, "/etc")
	os.Setenv(EnvAgencySource, "agency.json")

	ag := NewAgency()

	var buf bytes.Buffer
	ag.PrintEnv(&buf)

	get := buf.String()
	want := "BARITO_PATH=/opt/barito\nAGENCY_SOURCE=agency.json\n"

	FatalIf(t, get != want, "print wrong value/format: %s", get)
}

func TestAgency_AgentPath(t *testing.T) {
	os.Setenv(EnvBaritoPath, "/opt/barito2")
	ag := NewAgency()

	agentPath := ag.AgentPath("kafka")
	FatalIf(t, agentPath != "/opt/barito2/kafka", "generate wrong agentPath: %s", agentPath)
}

func TestAgency_Remove(t *testing.T) {
	os.Setenv(EnvBaritoPath, ".")
	os.Setenv(EnvAgencySource, "agency.json")

	ag := NewAgency()
	ag.Add("kafka", "1.1", "settings")

	record := ag.Get("kafka")
	FatalIf(t, record == nil, "record should be not nil")

	ag.Remove("kafka")

	record = ag.Get("kafka")
	FatalIf(t, record != nil, "record should be nil: %v", record)
}
