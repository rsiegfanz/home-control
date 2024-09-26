package config

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type Config interface{}

func LoadConfig[T Config]() (T, error) {
	var config T

	v := reflect.ValueOf(&config).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		tag := fieldType.Tag.Get("env")

		envVar := tag

		value, exists := os.LookupEnv(envVar)
		if !exists {
			return config, fmt.Errorf("missing environment variable %s", envVar)
		}

		if value == "" {
			return config, fmt.Errorf("missing value for variable %s", envVar)
		}

		err := setField(field, value)
		if err != nil {
			return config, fmt.Errorf("failed to set field %s: %w", fieldType.Name, err)
		}

	}

	return config, nil
}

func setField(field reflect.Value, value string) error {
	if !field.CanSet() {
		return errors.New("field can't be set")
	}

	switch field.Kind() {
	case reflect.String:
		field.SetString(value)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intVal, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		field.SetInt(intVal)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintVal, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		field.SetUint(uintVal)
	case reflect.Float32, reflect.Float64:
		floatVal, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		field.SetFloat(floatVal)
	case reflect.Bool:
		boolVal, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		field.SetBool(boolVal)
	default:
		return fmt.Errorf("invalid field type %s is not supported", field.Kind())
	}

	return nil
}
