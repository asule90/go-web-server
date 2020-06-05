package utils

import "log"

func Describe(i interface{}) {
	log.Printf("(%T, %v)\n", i, i)
}
