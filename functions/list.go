package functions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SingleList(list string, hideFails bool, malicious string, cookie string) {
	file, err := os.Open(list)
	if err != nil {
		if !hideFails {
			fmt.Println("\033[38;5;196m" + "[-] Error reading file" + "\033[0m")
			return
		}
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		SingleUrl(line, hideFails, malicious, cookie)

	}

}
