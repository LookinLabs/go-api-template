package swagger

type StatusUnprocessableContent struct {
	Message string `json:"message" example:"Posted data is not valid"`
	Code    string `json:"code" example:"UNPROCESSABLE_CONTENT"`
}