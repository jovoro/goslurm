package goslurm

import (
	"io"
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestAuthRaw(t *testing.T) {
	hc := http.Client{}

	// Raw client
	c, err := GenClient("http://arc.ft.uam.es:6820", "", "", &hc)
	if err != nil {
		t.Fatalf("err: %v\n", err)
	}

	resp, err := c.SlurmV0038Ping(context.TODO())
	if err != nil {
		t.Fatalf("err: %v\n", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected HTTP 200 but received %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("err: %v\n", err)
	}

	fmt.Printf("body: %s\n", body)
}

func TestAuthResp(t *testing.T) {
	hc := http.Client{}

	// Raw client
	c, err := GenClientWithResponses("http://arc.ft.uam.es:6820", "", "", &hc)
	if err != nil {
		t.Fatalf("err: %v\n", err)
	}

	resp, err := c.SlurmV0038PingWithResponse(context.TODO())
	if err != nil {
		t.Fatalf("err: %v\n", err)
	}
	if resp.StatusCode() != http.StatusOK {
		t.Fatalf("Expected HTTP 200 but received %d", resp.StatusCode())
	}

	fmt.Printf("body: %#v\n", resp.JSON200)
}
