package goslurm

import (
	"context"
	"net/http"
	"os"
)

func genAuthHeaders(user, token string) func(ctx context.Context, req *http.Request) error {
	tmpUser := user
	if user == "" {
		envUser, ok := os.LookupEnv("SLURM_USER")
		if ok {
			tmpUser = envUser
		}
	}

	tmpToken := token
	if token == "" {
		envToken, ok := os.LookupEnv("SLURM_JWT")
		if ok {
			tmpToken = envToken
		}
	}

	return func(ctx context.Context, req *http.Request) error {
		req.Header.Add("X-SLURM-USER-NAME", tmpUser)
		req.Header.Add("X-SLURM-USER-TOKEN", tmpToken)
		return nil
	}
}
