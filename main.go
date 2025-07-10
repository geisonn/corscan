package main

import (
	"github.com/geisonn/corscan/functions"
	"flag"
	"fmt"
)

func main() {
	help := flag.Bool("help", false, "Display this help message")
	flag.BoolVar(help, "h", false, "Display this help message (shorthand)")

	url := flag.String("url", "", "Single target URL (e.g., https://example.com)")
	flag.StringVar(url, "u", "", "Single target URL (shorthand)")

	list := flag.String("list", "", "File containing a list of target URLs")
	flag.StringVar(list, "l", "", "List of target URLs (shorthand)")

	cookie := flag.String("cookie", "", "HTTP cookie to include in requests (e.g., 'SESSIONID=abc123')")
	flag.StringVar(cookie, "c", "", "HTTP cookie to include in requests (shorthand)")

	malicious := flag.String("malicious", "http://corscan.go", "Malicious Origin to use in CORS requests (e.g., https://evil.com)")
	flag.StringVar(malicious, "m", "http://corscan.go", "Malicious Origin to use in CORS requests (shorthand)")

	hideFails := flag.Bool("hide-fails", false, "Only show successful bypass attempts")
	flag.BoolVar(hideFails, "hf", false, "Only show successful bypass attempts (shorthand)")
	flag.Parse()

	fmt.Println("\033[33m" + `   _____ ____  _____   _____  _____          _   _ 
  / ____/ __ \|  __ \ / ____|/ ____|   /\   | \ | |
 | |   | |  | | |__) | (___ | |       /  \  |  \| |
 | |   | |  | |  _  / \___ \| |      / /\ \ | .   |
 | |___| |__| | | \ \ ____) | |____ / ____ \| |\  |
  \_____\____/|_|  \_\_____/ \_____/_/    \_\_| \_|
                                                   
       	    Created by: Geison Nunes                                              

` + "\033[0m")

	if *help {
		functions.Help()
		return
	}

	if *url != "" {
		functions.SingleUrl(*url, *hideFails, *malicious, *cookie)
	}

	if *list != "" {
		functions.SingleList(*list, *hideFails, *malicious, *cookie)
	}
}
