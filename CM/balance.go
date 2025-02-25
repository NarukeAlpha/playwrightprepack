package CM

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getBalance(apiKey string) (float32, error) {
	body := cKey{
		ClientKey: apiKey,
	}
	jb, err := json.Marshal(body)
	if err != nil {
		return 0, err
	}
	req, err := http.NewRequest("POST", "https://api.capmonster.cloud/getBalance", io.NopCloser(strings.NewReader(string(jb))))
	if err != nil {
		return 0, err
	}
	clt := http.Client{
		Timeout: 30000,
	}
	resp, err := clt.Do(req)
	if err != nil {
		return 0, errors.Join(err, errors.New(fmt.Sprint("Failed to perform request ", req)))
	}
	respj := Response{}
	err = json.NewDecoder(resp.Body).Decode(&respj)
	if err != nil {
		return 0, err
	}
	if respj.ErrorId == 0 && respj.Balance != 0 {
		return respj.Balance, nil
	} else {
		return 0, err
	}
}
