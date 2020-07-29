package utils

type AppError struct {
	Msg    string `json:"msg"`
	Status int    `json:"status"`
	Code   string `json:"code"`
}
