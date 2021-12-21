package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	ProjectName string   `json:"project_name"`
	GroupId     string   `json:"group_id"`
	ArtifactId  string   `json:"artifact_id"`
	Version     string   `json:"version"`
	JavaVersion string   `json:"java_version"`
	Modules     []string `json:"modules"`
}

func LoadConfig(file string) (*Config, error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
