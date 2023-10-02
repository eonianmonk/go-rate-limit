package backend

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/eonianmonk/go-rate-limit/backend/cli"
	"github.com/eonianmonk/go-rate-limit/backend/http/responses"
	"github.com/gofiber/fiber/v2"
)

func startApp(maxRate int) {
	// can omit "--max-rate 50", because it is default value
	cli.Run([]string{
		"placeholder",
		"run",
		"--max-rate",
		strconv.Itoa(maxRate),
		"--cfg",
		"../config.yaml",
	})
}

func TestIntegration(t *testing.T) {
	maxRate := 50
	go startApp(maxRate)
	wantedRLExceededMessage := "rate limit exceeded"

	rateEp := "http://localhost:8080/api/v1/rate"
	var resp responses.GetRate

	for i := 0; i < maxRate; i++ {
		req := fiber.Get(rateEp)
		_, body, errs := req.Bytes()
		if len(errs) > 0 {
			t.Fatalf("failed to make req to svc: %s", errs)
		}
		err := json.Unmarshal(body, &resp)
		if err != nil {
			t.Fatalf("failed to unmarshal response body: %s", err.Error())
		}
		if int(resp.Rate) != i {
			t.Fatalf("unexpected value: got %d instead of %d", resp.Rate, i)
		}
	}

	req := fiber.Get(rateEp)
	_, body, errs := req.Bytes()
	if len(errs) > 0 {
		t.Fatalf("failed to make req to svc: %s", errs)
	}
	var errResp responses.Error
	err := json.Unmarshal(body, &errResp)
	if err != nil {
		t.Fatalf("failed to unmarshal error response: %s", err.Error())
	}
	if errResp.Error != wantedRLExceededMessage {
		t.Fatalf("got unexpected error: %s, wanted %s", errResp.Error, wantedRLExceededMessage)
	}

}
