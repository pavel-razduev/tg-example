package types

type BaseRequest struct {
	A string
}
type BaseResponse struct {
	Payload string `json:"payload"`
}
