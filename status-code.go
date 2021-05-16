package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

var customStatus = map[int]string{
	// Unofficial codes
	103: "Checkpoint",

	218: "This is fine",

	419: "Page Expired",
	430: "Request Header Fields Too Large",

	509: "Bandwidth limit exceeded",
	529: "Site is overloaded",
	598: "Network read timeout error",

	// Internet Information Services
	440: "Login Time-out",
	449: "Retry With",
	451: "Redirect",

	// Nginx
	444: "No Response",
	494: "Request header too large",
	495: "SSL Certificate Error",
	496: "SSL Certificate Required",
	497: "HTTP Request Sent to HTTPS port",
	499: "Client Closed Request",

	// Cloudflare
	520: "Web Server Returned an Unknown Error",
	521: "Web Server is Down",
	522: "Connection Timed Out",
	523: "Origin is Unreachable",
	524: "A Timeout Occurred",
	525: "SSL Handshake Failed",
	526: "Invalid SSL Certificate",
	527: "Railgun Error",
	530: "Cloudflare Error",

	// AWS ELB
	561: "Unauthorized",
}

// statusCodeText returns status code text for codes known to Go, non standard
// codes defined by this api, or the empty string for unknown codes.
func statusCodeText(code int) string {
	if text := http.StatusText(code); text != "" {
		return text
	} else if customStatus[code] != "" {
		return customStatus[code]
	}
	return ""
}

// statusCode returns a HTTP code + message for a user provided param
func statusCodeHandler(s *server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "statusCode")
		codeInt, err := strconv.Atoi(code)
		// See https://golang.org/src/net/http/server.go?s=33920:33955#L1089
		if err != nil || codeInt < 100 || codeInt > 999 {
			http.Error(w, "Invalid status code", http.StatusBadRequest)
			return
		}
		http.Error(w, statusCodeText(codeInt), codeInt)
	}
}
