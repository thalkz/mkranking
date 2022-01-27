package api

type JsonResponse struct {
	Status string      `json:"status"`
	Error  error       `json:"error,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}
