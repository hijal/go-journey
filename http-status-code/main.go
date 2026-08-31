package main

import "fmt"

const (
	statusOk            = 200
	statusBadRequest    = 400
	statusUnauthorized  = 401
	statusNotFound      = 404
	statusInternalError = 500
)

const (
	msgOk            = "OK"
	msgBadRequest    = "Bad Request"
	msgUnauthorized  = "Unauthorized"
	msgNotFound      = "Not Found"
	msgInternalError = "Internal Server Error"
)

func statusText(code int) string {
	switch code {
	case statusOk:
		return msgOk
	case statusBadRequest:
		return msgBadRequest
	case statusUnauthorized:
		return msgUnauthorized
	case statusNotFound:
		return msgNotFound
	case statusInternalError:
		return msgInternalError
	default:
		return "Unknown Status"
	}
}

func main() {
	codes := []int{
		statusOk, statusUnauthorized, statusNotFound, 503,
	}

	for _, code := range codes {
		fmt.Printf("%d -> %s\n", code, statusText(code))
	}
}
