package commands

import "log"

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
