package CM

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func GetBalance(apiKey string) (float32, error) {
	body := cKey{
		ClientKey: apiKey,
	}
	jb, err := json.Marshal(body)
	if err != nil {
		return 0, err
	}
	bodyreader := bytes.NewReader(jb)
	req, err := http.NewRequest("POST", "https://api.capmonster.cloud/getBalance", bodyreader)
	if err != nil {
		return 0, err
	}
	clt := http.Client{
		Timeout: time.Minute * 5,
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
