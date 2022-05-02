package trex

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getSummary(target string) (*summary, error) {
	resp, err := http.Get(fmt.Sprintf("%s/summary", target))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var summary summary

	err = json.NewDecoder(resp.Body).Decode(&summary)
	return &summary, err
}
