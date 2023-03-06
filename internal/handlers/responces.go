package handlers

type ServerError struct {
	Message string `json:"message" example:"status server error"`
}

type HTTPError struct {
	Message string `json:"message" example:"status bad request"`
}

type HTTPSuccess struct {
	Message string `json:"message" example:"OK"`
}
