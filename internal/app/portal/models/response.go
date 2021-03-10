package models

type Response struct {
	Code        int
	ContentType string
	Data        []byte
}
