package functions

import (
	"fmt"
	"net/http"
	"os"
)

func Error(message string, hideFails bool) {
	if !hideFails {
		fmt.Println("\033[38;5;196m" + message + "\033[0m")
	}
}

func Info(message string) {
	fmt.Println("\033[38;5;226m" + message + "\033[0m")
}

func NewRequest(client *http.Client, target, origin, cookie string) (*http.Response, error) {
	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		Error("[-] Error making request", HideFailsEnabled)
	}
	req.Header.Set("Origin", origin)
	if cookie != "" {
		req.Header.Set("Cookie", cookie)
	}
	return client.Do(req)
}

func GeneratePoC(filename, content string) {
	content += "\n"
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		Error("[-] Error creating payload file: "+filename, HideFailsEnabled)
		return
	}
	Success("[+] Payload file created: " + filename)
}

func Success(message string) {
	fmt.Println("\033[38;5;226m" + message + "\033[0m")
}
