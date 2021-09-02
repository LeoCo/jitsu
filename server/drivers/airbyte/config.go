package airbyte

import (
	"errors"
	"github.com/jitsucom/jitsu/server/airbyte"
	"strings"
)

//Config is a dto for airbyte configuration serialization
type Config struct {
	DockerImage            string            `mapstructure:"docker_image" json:"docker_image,omitempty" yaml:"docker_image,omitempty"`
	Config                 interface{}       `mapstructure:"config" json:"config,omitempty" yaml:"config,omitempty"`
	Catalog                interface{}       `mapstructure:"catalog" json:"catalog,omitempty" yaml:"catalog,omitempty"`
	InitialState           interface{}       `mapstructure:"initial_state" json:"initial_state,omitempty" yaml:"initial_state,omitempty"`
	StreamTableNames       map[string]string `mapstructure:"stream_table_names" json:"stream_table_names,omitempty" yaml:"stream_table_names,omitempty"`
	StreamTableNamesPrefix string            `mapstructure:"stream_table_name_prefix" json:"stream_table_name_prefix,omitempty" yaml:"stream_table_name_prefix,omitempty"`
}

//Validate returns err if configuration is invalid
func (ac *Config) Validate() error {
	if ac == nil {
		return errors.New("Airbyte config is required")
	}

	if ac.DockerImage == "" {
		return errors.New("Airbyte docker_image is required")
	}

	if ac.Config == nil {
		return errors.New("Airbyte config is required")
	}

	if ac.StreamTableNames == nil {
		ac.StreamTableNames = map[string]string{}
	}

	ac.DockerImage = strings.TrimPrefix(ac.DockerImage, airbyte.DockerImageRepositoryPrefix)

	return nil
}