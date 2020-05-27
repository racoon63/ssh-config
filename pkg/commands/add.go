package commands

import (
	"fmt"
	"os"
)

// Add adds a host entry to the ssh config
func Add(name string, host string, user string, port string, keyPath string) error {

	fmt.Println(name)
	fmt.Println(host)
	fmt.Println(user)
	fmt.Println(port)
	fmt.Println(keyPath)

	var tmpl string = "\nHost " + name + "\n  HostName " + host

	if user != "" {
		tmpl = tmpl + "\n  User " + user
	}
	if port != "" {
		tmpl = tmpl + "\n  Port " + port
	}
	if keyPath != "" {
		tmpl = tmpl + "\n  IdentityFile " + keyPath
	}

	f, err := os.OpenFile("config", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkError(err)
	defer f.Close()
	_, err = f.WriteString(tmpl)
	checkError(err)

	return nil
}
