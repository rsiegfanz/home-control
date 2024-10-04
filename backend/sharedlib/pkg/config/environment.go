package config

import (
	"flag"
	"log"
)

func IsProd() bool {
	log.Println("checking env")
	env := flag.String("env", "", "set the environment (prod/dev)")
	flag.Parse()
	log.Println(*env)
	if *env == "prod" {
		return true
	}

	return false
}
