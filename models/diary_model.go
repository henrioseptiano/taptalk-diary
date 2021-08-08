package models

type ReqCreateDiary struct {
	DatePost string `json:"datePost"`
	Title    string `json:"title"`
	BodyText string `json:"bodyText"`
}

type ReqUpdateDiary struct {
	DatePost string `json:"datePost"`
	Title    string `json:"title"`
	BodyText string `json:"bodyText"`
}
