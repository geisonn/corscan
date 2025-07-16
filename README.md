# Corscan – Smart Detection and Exploitation of CORS Misconfigurations
Corscan is a Go-based tool that automates the detection and exploitation of CORS misconfigurations. It performs intelligent analysis of Access-Control-Allow-Origin behavior, detects credentialed cross-origin access, and generates PoCs automatically — streamlining CORS testing for bug bounty and offensive security workflows.

## 🔍 Overview

While most tools detect CORS misconfigurations using static checks or basic reflection patterns, **Corscan** focuses on practical, real-world exploitation scenarios. It sends crafted requests with various `Origin` headers, observes server behavior with and without cookies, and automatically generates proof-of-concept payloads when vulnerable behavior is detected.

This targeted approach reduces noise, avoids unnecessary traffic, and highlights only what matters — exploitable and high-risk CORS configurations.

## ⚡ Features

- ✅ Detects permissive `Access-Control-Allow-Origin` behavior  
- ✅ Supports credentialed CORS detection via session cookies (`-c`)  
- ✅ Automatically generates PoC payloads for confirmed vulnerabilities  
- ✅ Designed for bug bounty and offensive security workflows  
- ✅ Lightweight, fast, and written in pure Go  

## 🛠 How to Use
```
└─# corscan -h         
   _____ ____  _____   _____  _____          _   _ 
  / ____/ __ \|  __ \ / ____|/ ____|   /\   | \ | |
 | |   | |  | | |__) | (___ | |       /  \  |  \| |
 | |   | |  | |  _  / \___ \| |      / /\ \ | .   |
 | |___| |__| | | \ \ ____) | |____ / ____ \| |\  |
  \_____\____/|_|  \_\_____/ \_____/_/    \_\_| \_|
                                                   
            Created by: Geison Nunes                                              


Usage:
  -u,  --url           Single target URL (e.g., https://example.com)
  -l,  --list          File containing a list of target URLs
  -c,  --cookie        HTTP cookie to include in requests (e.g., 'SESSIONID=abc123')
  -m,  --malicious     Malicious Origin to use in CORS requests (default: http://corscan.go)
  -hf, --hide-fails    Only show successful bypass attempts
  -h,  --help          Display this help message

```

## 📦 Installation
`go install github.com/geisonn/corscan@latest`
