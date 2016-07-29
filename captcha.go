package luosimao

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// MakeCaptcha : make a captcha
func MakeCaptcha(length int) string {
	captcha := ""
	if length < 1 {
		return captcha
	}
	max, _ := strconv.ParseInt(fmt.Sprintf("%.f", math.Pow10(length)), 10, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	captcha = fmt.Sprint(r.Int63n(max))
	for len(captcha) < length {
		captcha = "0" + captcha
	}
	return captcha[0:length]
}
