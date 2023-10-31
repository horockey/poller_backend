package http_controller

import (
	"crypto/aes"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	authTokenHeader string        = "X-Auth-Token"
	maxTimeDiff     time.Duration = time.Second * 5
)

func (ctrl *httpController) checkAuthToken(req *http.Request) error {
	token := req.Header.Get(authTokenHeader)
	if token == "" {
		return errors.New("got empty auth token header")
	}

	cipher, err := aes.NewCipher(ctrl.hashSeed)
	if err != nil {
		return fmt.Errorf("creating new cipher: %w", err)
	}

	resBytes := make([]byte, len(token))
	cipher.Decrypt(resBytes, []byte(token))

	resTime, err := time.Parse(time.RFC3339, string(resBytes))
	if err != nil {
		return fmt.Errorf("parsing time: %w", err)
	}

	if diff := time.Now().Sub(resTime); diff > maxTimeDiff {
		return fmt.Errorf("timediff is too great: %s (max: %s)", diff.String(), maxTimeDiff.String())
	}

	return nil
}
