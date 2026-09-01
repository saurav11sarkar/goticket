package httpResponse

type Success struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Meta    any    `json:"meta,omitempty"`
	Data    any    `json:"data,omitempty"`
}
