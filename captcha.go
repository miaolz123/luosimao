package luosimao

import (
	"fmt"
	"math/rand"
	"time"
)

// MakeCaptcha : make a captcha
func MakeCaptcha(length int) string {
	captcha := ""
	for i := 0; i < length; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		captcha += fmt.Sprint(r.Intn(10))
	}
	return captcha
}
