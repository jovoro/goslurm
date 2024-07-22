package goslurm

import "net/http"

type Response interface {
	StatusCode() int
}

func GenClient(server, user, token string, hc *http.Client) (*Client, error) {
	 return NewClient(server, WithHTTPClient(hc), WithRequestEditorFn(genAuthHeaders(user, token)))
}

func GenClientWithResponses(server, user, token string, hc *http.Client) (*ClientWithResponses, error) {
	 return NewClientWithResponses(server, WithHTTPClient(hc), WithRequestEditorFn(genAuthHeaders(user, token)))
}
