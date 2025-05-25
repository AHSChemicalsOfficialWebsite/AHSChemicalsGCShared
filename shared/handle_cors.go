package shared

import (
	"net/http"
)

func CorsEnabledFunction(response http.ResponseWriter, request *http.Request) bool {
	
	//Allowed origins
	allowedOrigins := map[string]bool{
		"http://localhost:3000": true, 
		"https://ahschemicalsdebug.web.app": true,
		"https://azurehospitalitysupply.com": true,
	}

	var origin string
	if allowedOrigins[request.Header.Get("Origin")] {
		origin = request.Header.Get("Origin")
	}

	if allowedOrigins[origin] {
		response.Header().Set("Access-Control-Allow-Origin", origin)
		response.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		response.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		response.Header().Set("Access-Control-Allow-Credentials", "true")
		response.Header().Set("Access-Control-Max-Age", "3600")
	}

	//Handle Pre flight request 
	if request.Method == http.MethodOptions {
		response.WriteHeader(http.StatusNoContent)
		return true 
	}

	return false
}