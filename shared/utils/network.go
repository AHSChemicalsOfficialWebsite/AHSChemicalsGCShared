package utils

import (
	"net/http"
	"net"
	"strings"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/shared/constants"
)

// CorsEnabledFunction handles Cross-Origin Resource Sharing (CORS) for HTTP requests.
//
// This function verifies if the request's Origin header is in the list of allowed origins.
// If allowed, it sets appropriate CORS headers on the HTTP response to enable cross-origin access.
//
// Additionally, it handles preflight (OPTIONS) requests by returning HTTP 204 No Content.
//
// Parameters:
//   - response: http.ResponseWriter used to write HTTP headers and responses.
//   - request: *http.Request representing the incoming HTTP request.
//
// Returns:
//   - true: If the request is an OPTIONS preflight request and has been handled.
//   - false: If the request is not an OPTIONS preflight request (normal processing should continue).
func CorsEnabledFunction(response http.ResponseWriter, request *http.Request) bool {
	allowedOrigins := map[string]bool{
		constants.CorsAllowOriginDebug:      true,
		constants.CorsAllowOriginStaging:    true,
		constants.CorsAllowOriginProduction: true,
	}

	var origin string
	if allowedOrigins[request.Header.Get("Origin")] {
		origin = request.Header.Get("Origin")
	}

	if allowedOrigins[origin] {
		response.Header().Set("Access-Control-Allow-Origin", origin)
		response.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
		response.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		response.Header().Set("Access-Control-Max-Age", "3600")
	}

	if request.Method == http.MethodOptions {
		response.WriteHeader(http.StatusNoContent) // 204 No Content
		return true
	}

	return false
}

// GetIp attempts to retrieve the IP address of the client making the HTTP request.
//
// This function checks for commonly used headers set by proxies and load balancers
// to extract the original client IP address in the following order:
//   1. "X-Forwarded-For": May contain a comma-separated list of IPs. The first one is typically the client IP.
//   2. "X-Real-Ip": Set by some reverse proxies to indicate the client’s IP.
//   3. request.RemoteAddr: Falls back to the address directly from the TCP connection.
//
// It returns the first valid, parsed IP address found in these sources, or an empty string if none is found.
//
// Parameters:
//   - request (*http.Request): The HTTP request from which to extract the client IP.
//
// Returns:
//   - string: The client's IP address, or an empty string if it cannot be determined.
//
func GetIp(request *http.Request) string {
	xff := request.Header.Get("X-Forwarded-For")
	if xff != "" {
		ips := strings.Split(xff, ",")
		ip := strings.TrimSpace(ips[0])
		if parsed := net.ParseIP(ip); parsed != nil {
			return parsed.String()
		}
	}

	xrip := request.Header.Get("X-Real-Ip")
	if xrip != "" {
		if parsed := net.ParseIP(xrip); parsed != nil {
			return parsed.String()
		}
	}

	ip, _, err := net.SplitHostPort(request.RemoteAddr)
	if err == nil {
		if parsed := net.ParseIP(ip); parsed != nil {
			return parsed.String()
		}
	}

	return ""
}