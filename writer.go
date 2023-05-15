package timeout

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

type TimeoutWriter struct {
	gin.ResponseWriter
	// header
	H http.Header
	// body
	Body           *bytes.Buffer
	TimeoutOptions // TimeoutOptions in options.go

	Code        int
	Mu          sync.Mutex
	TimedOut    bool
	WroteHeader bool
}

func (tw *TimeoutWriter) Write(b []byte) (int, error) {
	tw.Mu.Lock()
	defer tw.Mu.Unlock()
	if tw.TimedOut {
		return 0, nil
	}

	return tw.Body.Write(b)
}

func (tw *TimeoutWriter) WriteHeader(code int) {
	checkWriteHeaderCode(code)
	tw.Mu.Lock()
	defer tw.Mu.Unlock()
	if tw.TimedOut {
		return
	}
	tw.writeHeader(code)
}

func (tw *TimeoutWriter) SetContentType(typ string) {
	tw.Header().Set("Content-Type", typ)
}

func (tw *TimeoutWriter) writeHeader(code int) {
	tw.WroteHeader = true
	tw.Code = code
}

func (tw *TimeoutWriter) WriteHeaderNow() {}

func (tw *TimeoutWriter) Header() http.Header {
	return tw.H
}

func checkWriteHeaderCode(code int) {
	if code < 100 || code > 999 {
		panic(fmt.Sprintf("invalid WriteHeader code %v", code))
	}
}
