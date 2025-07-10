package functions

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func SingleUrl(target string, hideFails bool, malicious string, cookie string) {
	if !strings.HasPrefix(target, "http://") && !strings.HasPrefix(target, "https://") {
		target = "http://" + target
	}

	parsedURL, err := url.Parse(target)
	if err != nil {
		Error("[-] Failed to parse URL", hideFails)
		return
	}

	client := &http.Client{}

	resp, err := NewRequest(client, target, malicious, cookie)
	if err != nil {
		fmt.Printf("\033[38;5;27m"+"[*] Host: %s\n"+"\033[0m", parsedURL.Host)
		Error("[-] Error making request", hideFails)
		return
	}

	defer resp.Body.Close()

	fmt.Printf("\033[38;5;27m"+"[*] Host: %s\n"+"\033[0m", parsedURL.Host)
	_, allowCredentials := resp.Header["Access-Control-Allow-Credentials"]
	if allowCredentials {
		Info("[INFO] Access-Control-Allow-Credentials header present")
	}

	_, allowOrigin := resp.Header["Access-Control-Allow-Origin"]
	if allowOrigin {
		Info("[INFO] Access-Control-Allow-Origin header present")
	}

	if allowOrigin && allowCredentials {
		OriginNull(client, target, malicious, cookie, parsedURL, hideFails)
		MaliciousOrigin(client, target, malicious, cookie, parsedURL, hideFails)
	} else {
		Error("[-] Access-Control-Allow-Credentials and Access-Control-Allow-Origin headers are not present in the same request", hideFails)
	}
}
