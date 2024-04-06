package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config map[string]any

func GetConfig(p string) Config {
	path := fmt.Sprintf("./configs/%s", p)

	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var data Config
	if err := json.Unmarshal(content, &data); err != nil {
		panic(err)
	}

	return data
}

func (c Config) Get(key string) any {
	value, ok := c[key]
	if !ok {
		panic("unable to find key")
	}

	return value
}

func (c Config) Config(key string) Config {
	value, ok := c[key]
	if !ok {
		panic("unable to find key")
	}

	return value.(Config)
}

func (c Config) String(key string) string {
	value, ok := c[key]
	if !ok {
		panic("unable to find key")
	}

	strVal, ok := value.(string)
	if !ok {
		panic("value is not a string")
	}

	return strVal
}

func (c Config) Int(key string) int {
	value, ok := c[key]
	if !ok {
		panic("unable to find key")
	}

	intVal, ok := value.(int)
	if !ok {
		panic("value is not an int")
	}

	return intVal
}

func (c Config) Float(key string) float64 {
	value, ok := c[key]
	if !ok {
		panic("unable to find key")
	}

	floatVal, ok := value.(float64)
	if !ok {
		panic("value is not a float64")
	}

	return floatVal
}

func (c Config) Bool(key string) bool {
	value, ok := c[key]
	if !ok {
		panic("unable to find key")
	}

	boolVal, ok := value.(bool)
	if !ok {
		panic("value is not a bool")
	}

	return boolVal
}
