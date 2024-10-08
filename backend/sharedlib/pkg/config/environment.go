package config

import (
	"flag"
)

var isProd = -1

func IsProd() bool {
	if isProd == -1 {
		env := flag.String("env", "", "set the environment (prod/dev)")
		flag.Parse()

		if *env == "prod" {
			isProd = 1
		} else {
			isProd = 0
		}
	}

	if isProd == 1 {
		return true
	}

	return false
}
