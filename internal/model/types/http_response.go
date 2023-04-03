package types

type JSONRespHeader struct {
	ServerProcessTime string `json:"process_time,omitempty"`
	Message           string `json:"message,omitempty"`
	Reason            string `json:"reason,omitempty"`
	Code              int    `json:"code,omitempty"`
}

type JSONResponse struct {
	Header      JSONRespHeader `json:"header"`
	Data        interface{}    `json:"data"`
	ErrorString string         `json:"error,omitempty"`
}
