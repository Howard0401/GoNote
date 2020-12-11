package middleware

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync/atomic"
)

//Set RequestKeyID

type ctxKeyRequestID int

//the key holds th unique request ID in request context
const RequestIDKey ctxKeyRequestID = 0

var (
	//const prefix
	prefix string
	reqID  int64
)

//init requestID
func init() {
	//the host name reported by the kernel
	hostName, err := os.Hostname()

	if hostName == "" || err != nil {
		hostName = "localhost"
	}
	//看不懂
	var buf [12]byte
	var b64 string
	for len(b64) < 10 {
		_, _ = rand.Read(buf[:])
		b64 = base64.StdEncoding.EncodeToString(buf[:])
		b64 = strings.NewReplacer("+", "", "/", "").Replace(b64)
	}

	prefix = fmt.Sprintf("%s/%s", hostName, b64[0:10])
}

func AddRequestID(h http.Handler) http.Handler {
	// The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		myid := atomic.AddInt64(&reqID, 1)
		ctx := r.Context()
		// ServeHTTP should write reply headers and data to the ResponseWriter and then return. Returning signals that the request is finished
		ctx = context.WithValue(ctx, RequestIDKey, fmt.Sprintf("%s-%06d", prefix, myid))
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetReqID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if reqID, ok := ctx.Value(RequestIDKey).(string); ok {
		return reqID
	}
	return ""
}
