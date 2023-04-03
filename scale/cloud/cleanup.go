package cloud

import (
	"fmt"
	"net/http"
)

func (fleet *Fleet) Cleanup(host *Host) (err error) {
	if fleet.Entrypoint == "" {
		return fmt.Errorf("HTTP endpoint is not defined")
	}
	var (
		req  *http.Request
		resp *http.Response
	)
	req, err = http.NewRequest("POST", fmt.Sprintf("http://%s/unregister", fleet.Entrypoint), nil)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Host", host.Name)
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("unregister request failed for %s: %w", host.Name, err)
	}
	defer func() { _ = resp.Body.Close() }()
	return nil
}
