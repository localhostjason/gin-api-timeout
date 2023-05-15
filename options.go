package timeout

import (
	"encoding/json"
	"net/http"
	"time"
)

type CallBackFunc func(*http.Request)
type Option func(*TimeoutWriter)

type TimeoutOptions struct {
	CallBack      CallBackFunc
	DefaultMsg    string
	Timeout       time.Duration
	ErrorHttpCode int
}

func WithTimeout(d time.Duration) Option {
	return func(t *TimeoutWriter) {
		t.Timeout = d
	}
}

// WithErrorHttpCode Optional parameters
func WithErrorHttpCode(code int) Option {
	return func(t *TimeoutWriter) {
		t.ErrorHttpCode = code
	}
}

// WithDefaultMsg Optional parameters
func WithDefaultMsg(s map[string]interface{}) Option {
	return func(t *TimeoutWriter) {
		data, err := json.Marshal(s)
		if err != nil {
			return
		}
		t.DefaultMsg = string(data)
	}
}

// WithCallBack Optional parameters
func WithCallBack(f CallBackFunc) Option {
	return func(t *TimeoutWriter) {
		t.CallBack = f
	}
}
