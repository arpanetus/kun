package config

import "gopkg.in/yaml.v3"
import "fmt"

type KunConfig struct {
	Latitude  float64 `yaml:"latitude"`
	Longitude float64 `yaml:"longitude"`
}

func DefaultConfig() *KunConfig {
	// default location is Almaty :D.
	return &KunConfig{
		Latitude:  43.2384923,
		Longitude: 76.943262,
	}
}

func (c *KunConfig) Marshal() ([]byte, error) {
	m, err := yaml.Marshal(c)
	if err != nil {
		return []byte{}, fmt.Errorf("cannot serialize config: %s", err)
	}

	return m, nil
}

func (c *KunConfig) Unmarshal(m []byte) error {
	err := yaml.Unmarshal(m, c)
	if err != nil {
		return fmt.Errorf("cannot deserialize config: %s", err)
	}
	return nil
}
