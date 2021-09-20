package controllers

import "mime/multipart"

type Context interface {
	Param(string) string
	Bind(interface{}) error
	Status(int)
	JSON(int, interface{})
	Cookie(name string) (string, error)
	SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)
	FormFile(name string) (*multipart.FileHeader, error)
	GetPostForm(key string) (string, bool)
}
