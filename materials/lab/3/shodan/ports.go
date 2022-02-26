package shodan

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (s *Client) ListPorts() ([]byte, error) {

	req, err := http.Get((fmt.Sprintf("%s/shodan/ports?key=%s", BaseURL, s.apiKey)))
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
