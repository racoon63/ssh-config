package commands

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/urfave/cli"
)

var sshConfigPath string = os.Getenv("HOME") + "/.ssh/config"

// List prints all found hosts in ~/.ssh/config
func List(c *cli.Context) error {
	file, err := os.Open(sshConfigPath)
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if m, err := regexp.MatchString(`^Host\s\S+`, scanner.Text()); m {
			checkError(err)
			hosts := strings.Split(scanner.Text(), " ")
			fmt.Println(strings.Join(hosts[1:], " "))
		}
	}
	return nil
}
