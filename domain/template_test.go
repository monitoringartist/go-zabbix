package domain

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"path/filepath"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestGroup(t *testing.T) {
	content, err := ioutil.ReadFile(filepath.Join("testdata", t.Name()+".xml"))
	if err != nil {
		t.Fatalf("failed reading: %s", err)
	}
	xGroup := &Group{}
	err = xml.Unmarshal([]byte(content), &xGroup)
	if err != nil {
		t.Fatalf("XML unmarshal error: %s", err)
	}

	content, err = ioutil.ReadFile(filepath.Join("testdata", t.Name()+".json"))
	if err != nil {
		t.Fatalf("failed reading: %s", err)
	}
	sJSON, err := json.MarshalIndent(xGroup, "", "    ")
	if err != nil {
		t.Fatalf("failed reading: %s", err)
	}
	sJSON = append(sJSON, byte('\n'))
	res := bytes.Compare(sJSON, content)
	if res != 0 {
		t.Fatalf("failed JSON comparison")
	}

	content, err = ioutil.ReadFile(filepath.Join("testdata", t.Name()+".yaml"))
	if err != nil {
		t.Fatalf("failed reading: %s", err)
	}
	sYaml, err := yaml.Marshal(xGroup)
	if err != nil {
		t.Fatalf("failed reading: %s", err)
	}
	res = bytes.Compare(sYaml, content)
	if res != 0 {
		t.Fatalf("failed YAML comparison")
	}
}
