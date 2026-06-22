package utils

import (
	"net/http"
	"net"
	"strings"
)

const (
	CorsAllowOriginDebug      = "http://localhost:5173"
	CorsAllowOriginStaging    = "https://ahschemicalsstaging.web.app"
	CorsAllowOriginProduction = "https://azurehospitalitysupply.com"
)

func CorsEnabledFunction(response http.ResponseWriter, request *http.Request) bool {
	allowedOrigins := map[string]bool{
		CorsAllowOriginDebug:      true,
		CorsAllowOriginStaging:    true,
		CorsAllowOriginProduction: true,
	}

	var origin string
	if allowedOrigins[request.Header.Get("Origin")] {
		origin = request.Header.Get("Origin")
	}

	if allowedOrigins[origin] {
		response.Header().Set("Access-Control-Allow-Origin", origin)
		response.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
		response.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Access-Key")
		response.Header().Set("Access-Control-Max-Age", "3600")
	}

	if request.Method == http.MethodOptions {
		response.WriteHeader(http.StatusNoContent) // 204 No Content
		return true
	}

	return false
}

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