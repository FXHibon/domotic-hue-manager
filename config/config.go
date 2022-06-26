package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type AppConfiguration struct {
	Api ApiConfiguration
}

type ApiConfiguration struct{ Uri, User string }

type InvalidPath struct {
	Path, Reason string
}

func (i InvalidPath) Error() string {
	return fmt.Sprintf("Invalid path %#v reason=%#v", i.Path, i.Reason)
}

type InvalidField struct {
	FieldName, Reason string
}

func (i InvalidField) Error() string {
	return fmt.Sprintf("Invalid field=%#v reason=%#v", i.FieldName, i.Reason)
}

func LoadConfigurationOrDie() AppConfiguration {
	path, valued := os.LookupEnv("CREDENTIALS_PATH")
	if !valued {
		path = ".credentials.json"
	}

	file, err := os.ReadFile(path)
	if err != nil {
		log.Panic(InvalidPath{path, err.Error()})
	}

	log.Printf("Read configuration file at '%v'", path)

	var conf AppConfiguration

	_ = json.NewDecoder(bytes.NewReader(file)).Decode(&conf)

	if conf.Api.User == "" {
		log.Panic(InvalidField{".api.user", "empty string not allowed"})
	}

	if conf.Api.Uri == "" {
		log.Panic(InvalidField{".api.uri", "empty string not allowed"})
	}

	return conf
}
