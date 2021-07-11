package service

import (
	"fmt"
	"net/http"
	"time"
)

func ResponseIP(r *http.Request) {
	fmt.Printf("%v / %v  %v\n", r.Method, r.RemoteAddr, time.Now().Format(time.ANSIC))
}
