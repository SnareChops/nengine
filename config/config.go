package config

import (
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"

	"gopkg.in/yaml.v2"
)

type Config[T any] struct {
	Contents T
	Final    *T
}

func NewConfigFromFile[T any](file string, def T) (*Config[T], error) {
	reader, err := os.Open(file)
	if err != nil && err != os.ErrNotExist {
		return nil, err
	}
	if err != os.ErrNotExist {
		return NewConfig(reader, def)
	}
	var contents T
	return &Config[T]{Contents: contents, Final: &def}, nil
}

func NewConfig[T any](reader io.Reader, def T) (*Config[T], error) {
	var contents T
	err := yaml.NewDecoder(reader).Decode(&contents)
	if err != nil {
		return nil, err
	}
	return &Config[T]{Contents: contents, Final: &def}, nil
}

func (c *Config[T]) Load() error {
	// Reflect over properties present in contents and only override them if not "empty"
	con := reflect.ValueOf(c.Contents)
	final := reflect.ValueOf(c.Final).Elem()
	if con.Kind() != reflect.Struct || final.Kind() != reflect.Struct {
		return errors.New("Config cannot use non-struct types")
	}
	finalVal := reflect.ValueOf(&final).Elem()
	for i := 0; i < final.NumField(); i++ {
		fmt.Printf("Field %d\n", i)
		field := con.Field(i)
		fmt.Printf("Field: %s\n", field)
		finalField := finalVal.Field(i)
		fmt.Printf("Is zero %t\n", field.IsZero())
		if !field.IsZero() {
			finalField.Set(field)
		}
	}
	return nil
}

func (c Config[T]) Get() T {
	return *c.Final
}
