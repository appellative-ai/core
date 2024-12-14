package core

import "net/http"

const (
	IAMUser           = "X-User"
	IAMPassword       = "X-Password"
	IAMUri            = "X-Uri"
	IAMFrom           = "X-From"
	IAMCredentialsUri = "iam:credentials"
)

// IAMProvider - IAM provider interface
type IAMProvider interface {
	HttpExchange
}

func NewIAMRequest(from string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, IAMCredentialsUri, nil)
	req.Header = make(http.Header)
	req.Header.Set(IAMFrom, from)
	return req
}

func NewIAMResponseFromUri(uri string, status *Status) *http.Response {
	resp := &http.Response{Status: status.String(), StatusCode: status.HttpCode(), Header: make(http.Header)}
	resp.Header.Set(IAMUri, uri)
	return resp
}

func NewIAMResponseFromCredentials(user, password string, status *Status) *http.Response {
	resp := &http.Response{Status: status.String(), StatusCode: status.HttpCode(), Header: make(http.Header)}
	resp.Header.Set(IAMPassword, password)
	resp.Header.Set(IAMUser, user)
	return resp
}
