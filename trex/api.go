package trex

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Collector) getSummary() (*summary, error) {
	resp, err := http.Get(fmt.Sprintf("%s/summary", c.TrexAPIAddress))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var summary summary

	err = json.NewDecoder(resp.Body).Decode(&summary)
	return &summary, err
}
