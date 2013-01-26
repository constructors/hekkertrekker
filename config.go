package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

const (
	configFile = ".hekkertrekker"
)

var generalConfig struct {
	Token              string
	NewBranchCommitMsg string
	DeliverCommitMsg   string
	DoneCommitMsg      string
	DoneLabel          string
	Name               string
}
var repositoryConfig struct {
	ProjectID     int
	StagingBranch string
}

func initConfig() {
	var p string

	p = os.ExpandEnv(path.Join("$HOME", configFile))
	if contents, err := ioutil.ReadFile(p); err != nil {
		bye("%v\n", err)
	} else {
		if err := json.Unmarshal(contents, &generalConfig); err != nil {
			bye("%v\n", err)
		}
	}

	p = path.Join(hgRoot(), configFile)
	if contents, err := ioutil.ReadFile(p); err != nil {
		bye("%v\n", err)
	} else {
		if err := json.Unmarshal(contents, &repositoryConfig); err != nil {
			bye("%v\n", err)
		}
	}
}
