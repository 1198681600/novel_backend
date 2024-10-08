package global

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func CreateHttpServer() *http.Server {
	port := 7899
	if value, err := strconv.ParseInt(os.Getenv("HTTP_PORT"), 10, 32); err == nil && value != 0 {
		port = int(value)
	}
	return &http.Server{Addr: fmt.Sprintf(":%d", port)}
}
