package config

import (
	"flag"
)

func IsProd() bool {
	env := flag.String("env", "", "set the environment (prod/dev)")
	flag.Parse()
	if *env == "prod" {
		return true
	}

	return false
}
