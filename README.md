# luosimao [![GoDoc](https://godoc.org/github.com/miaolz123/luosimao?status.svg)](https://godoc.org/github.com/miaolz123/luosimao) [![Build Status](https://travis-ci.org/miaolz123/luosimao.svg?branch=master)](https://travis-ci.org/miaolz123/luosimao)

### A SDK of luosimao.com for Golang

```go
package main

import (
	"log"

	"github.com/miaolz123/luosimao"
)

// Please replace this section
const (
	keySMS           = "3ax6xnjp29jd6fds4gc373sgvjxteol0"
	keyVoice         = "3ax6xnjp29jd6fds4gc373sgvjxteol0"
	testMobileNumber = 13600000000
)

func main() {
	client := luosimao.Client{KeySMS: keySMS, KeyVoice: keyVoice}
	err := client.SendSMS(testMobileNumber, "Hello World!【Google Inc.】")
	if err != nil {
		log.Println("Send SMS error:", err)
	}
	err = client.SendVoice(testMobileNumber, luosimao.MakeCaptcha(4))
	if err != nil {
		log.Println("Send voice error:", err)
	}
}
```
