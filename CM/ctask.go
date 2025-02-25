package CM

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	pw "github.com/playwright-community/playwright-go"
)

func HandleTurnstileCookie(api string, page pw.Page, websitekey string, proxy pw.Proxy) error {

	apireq := page.Request()
	resp, err := apireq.Get(page.URL(), pw.APIRequestContextGetOptions{Timeout: pw.Float(3000)})
	if err != nil {
		return errors.Join(errors.New(fmt.Sprint("Failed initial request to pull website html")), err)
	}
	body, err := resp.Body()
	if err != nil {
		return err
	}

	encodedHTML := base64.StdEncoding.EncodeToString(body)
	ua, err := page.Evaluate("() => navigator.userAgent")
	uaString := ua.(string)

	prx := strings.Split(proxy.Server, ":")

	cargo := OuterTask{
		ClientKey: api,
		Task: Task{
			Type:               "TurnstileTask",
			WebsiteURL:         page.URL(),
			WebsiteKey:         websitekey,
			CloudflareTaskType: "cf_clearance",
			HtmlPageBase64:     encodedHTML,
			UserAgent:          uaString,
			ProxyType:          "http",
			ProxyAddress:       prx[0],
			ProxyPort:          prx[1],
			ProxyLogin:         *proxy.Username,
			ProxyPassword:      *proxy.Password,
		},
	}

	cargores, err := apireq.Post("https://api.capmonster.cloud/createTask",
		pw.APIRequestContextPostOptions{
			Data: cargo,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		})
	if err != nil {
		errors.Join(errors.New(fmt.Sprint("Failed creating task")), err)
	}
	taskResult := Response{}
	body, err = cargores.Body()
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&taskResult)

	if err != nil {
		return err
	} else if taskResult.ErrorId != 0 {
		return errors.Join(errors.New(fmt.Sprint("Ran out of captcha balance")), err)
	}
	spear := OuterTask{ClientKey: api, TaskId: taskResult.TaskId}
	for i := 0; i < 39; i++ {
		handle, err := apireq.Post("https://api.capmonster.cloud/getTaskResult", pw.APIRequestContextPostOptions{Data: spear})
		if err != nil {
			return err
		}
		handlebar := Response{}
		if body, err = handle.Body(); err != nil {
			return err
		}
		if err := json.Unmarshal(body, &handlebar); err != nil {
			return err
		}

		if handlebar.ErrorId != 0 {
			return errors.Join(errors.New(fmt.Sprint("Task ran into an issue")), err)
		} else if handlebar.Status == "processing" {
			time.Sleep(3000)
			continue
		}
		if handlebar.Status == "ready" {
			u, err := url.Parse(page.URL())
			if err != nil {
				return err
			}
			d := u.Hostname()
			p := "/"
			cookie := pw.OptionalCookie{
				Name:   "cf_clearance",
				Value:  handlebar.Solution.CfClearance,
				Domain: &d,
				Path:   &p,
			}
			lc := []pw.OptionalCookie{cookie}
			if err = page.Context().AddCookies(lc); err != nil {
				return errors.Join(errors.New("failed to add cookie to browser"), err)
			}

			break
		}
	}

	return nil
}
