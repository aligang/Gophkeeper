package config

import (
	"encoding/json"
	"google.golang.org/protobuf/encoding/protojson"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

func getServerConfigFromYaml(filePath string) *ServerConfig {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0400)

	if err != nil {
		panic(err)
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	configMap := map[string]any{}
	err = yaml.Unmarshal(bytes, configMap)
	if err != nil {
		panic(err)
	}
	jsonBytes, err := json.Marshal(configMap)
	s := &ServerConfig{}
	err = protojson.Unmarshal(jsonBytes, s)
	if err != nil {
		panic(err)
	}
	return s
}
