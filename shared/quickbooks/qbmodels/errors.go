package qbmodels

type QBOAuthErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type QBFaultErrorResponse struct {
	Fault QBFault `json:"Fault"`
	Time  string  `json:"time,omitempty"`
}

type QBFault struct {
	ErrorList []QBError `json:"Error"`
	Type      string    `json:"type"`
}

type QBError struct {
	Message       string `json:"Message"`              
	Detail        string `json:"Detail,omitempty"`     
	Code          string `json:"code"`                 
	Element       string `json:"element,omitempty"`    
	Severity      string `json:"severity"`             
	InnerErrorMsg string `json:"InnerError,omitempty"` 
}