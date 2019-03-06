package model

type Response struct {
	Total    int
	PageSize int
	Data     []UserDetail
}
