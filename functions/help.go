package functions

import "fmt"

func Help() {
	fmt.Println("Usage:")
	fmt.Println("  -u,  --url           Single target URL (e.g., https://example.com)")
	fmt.Println("  -l,  --list          File containing a list of target URLs")
	fmt.Println("  -c,  --cookie        HTTP cookie to include in requests (e.g., 'SESSIONID=abc123')")
	fmt.Println("  -m,  --malicious     Malicious Origin to use in CORS requests (default: http://corscan.go)")
	fmt.Println("  -hf, --hide-fails    Only show successful bypass attempts")
	fmt.Println("  -h,  --help          Display this help message")

}
