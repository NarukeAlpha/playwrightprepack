package ppp_test

import (
	"log"
	rand2 "math/rand"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	ppp "github.com/NarukeAlpha/playwrightprepack"
	"github.com/NarukeAlpha/playwrightprepack/CM"
	"github.com/playwright-community/playwright-go"
)

func init() {
	log.Println("Initializing")
	err := playwright.Install()
	if err != nil {
		log.Fatalf("Error installing playwright: %v", err)
	}
	log.Println("starting Tests")
}

var pw *playwright.Playwright

func TestMain(m *testing.M) {
	var err error
	pw, err = playwright.Run()
	if err != nil {
		log.Fatalf("Error running playwright: %v", err)
	}

	code := m.Run()
	os.Exit(code)
}

var wbbrowser playwright.BrowserContext

func delay(i int) {
	time.Sleep(time.Duration(i) * time.Second)
}

func TestProxyLoad(t *testing.T) {
	t.Log("testing proxy struct load")
	t.Run("ProxyLoad", func(t *testing.T) {
		var proxylist []*playwright.Proxy
		proxylist, err := ppp.ProxyLoad("./test.csv")
		if err != nil {
			t.Fatalf("Error loading proxy list: %v", err)
		} else if len(proxylist) != 32400 {
			t.Fatalf("Expected 32400 proxies, got %d", len(proxylist))
		}
		for i := 0; i < 500; i++ {
			randomint := rand2.Intn(len(proxylist))
			proxy := proxylist[randomint]
			//usr := "tusr"
			//psw := "tpw"
			usr := *proxy.Username
			psw := *proxy.Password
			if proxy.Server != "1.0.0.27:933" {
				t.Fatalf("Expected 1.0.0.27:933, got %v", proxy.Server)
			} else if usr != "tusr" {
				t.Fatalf("Expected tusr, got %v", proxy.Username)
			} else if psw != "tpw" {
				t.Fatalf("Expected tpw, got %v", proxy.Password)
			}
		}

	})
}

func TestWebKit(t *testing.T) {
	log.Println("Running WebKit test")
	var errp error
	//proxy := "161.0.70.152:5741:tnzpwplz:156y8h5d4l6q"
	wbbrowser, errp = ppp.PlaywrightInit(nil, 1, false, pw)
	if errp != nil {
		t.Fatalf("could not initialize playwright: %v", errp)
	}
	log.Println("Browser initialized")

	t.Run("SannySoft", func(t *testing.T) {
		page, err := wbbrowser.NewPage()
		if err != nil {
			t.Fatalf("could not create page: %v", err)
		}
		_, err = page.Goto("https://bot.sannysoft.com")
		if err != nil {
			t.Fatalf("could not navigate to page: %v", err)
		}
		err = page.WaitForLoadState(playwright.PageWaitForLoadStateOptions{
			State: playwright.LoadStateNetworkidle,
		})
		if err != nil {
			t.Fatalf("could not wait for load state: %v", err)
		}
		elmt, err := page.QuerySelector(`#webdriver-result`)
		if err != nil {
			t.Fatalf("could not find element: %v", err)
		}
		tx, err := elmt.InnerText()
		if err != nil {
			t.Fatalf("could not get text: %v", err)
		}
		if tx == "present (failed)" {
			t.Fatalf("webdriver detected, test failed")
		}
		elmt2, err2 := page.QuerySelector(`#advanced-webdriver-result`)
		if err != nil {
			t.Fatalf("could not find element: %v", err)
		}
		tx2, err2 := elmt2.InnerText()
		if err2 != nil {
			t.Fatalf("could not get text: %v", err2)
		}
		if tx2 == "present (failed)" {
			t.Fatalf("webdriver detected, test failed")
		}
		err = page.Close()
		if err != nil {
			t.Fatalf("could not close page: %v", err)
		}

	})

	t.Run("ReCaptcha", func(t *testing.T) {
		page, err := wbbrowser.NewPage()
		if err != nil {
			t.Fatalf("could not create page: %v", err)
		}
		_, err = page.Goto("https://antcpt.com/eng/information/demo-form/recaptcha-3-test-score.html")
		if err != nil {
			t.Fatalf("could not navigate to page: %v", err)
		}
		err = page.WaitForLoadState(playwright.PageWaitForLoadStateOptions{
			State: playwright.LoadStateNetworkidle,
		})
		if err != nil {
			t.Fatalf("could not wait for load state: %v", err)
		}
		delay(10)
		log.Println("mandatory 10 second sleep while recaptcha loads\noverkill but its safer")

		btn := page.GetByRole("button", playwright.PageGetByRoleOptions{
			Name: "Refresh score now!",
		})
		if btn == nil {
			t.Fatalf("could not find button to refresh captcha score")
		}
		err = btn.ScrollIntoViewIfNeeded()
		if err != nil {
			t.Fatalf("could not scroll to button: %v", err)
		}
		delay(2)
		err = btn.Click()
		if err != nil {
			t.Fatalf("could not click button: %v", err)
		}
		log.Println("refreshing captcha score")
		_ = page.WaitForLoadState(playwright.PageWaitForLoadStateOptions{
			State: playwright.LoadStateNetworkidle,
		})
		delay(2)
		elmt, err := page.QuerySelector(`#score`)
		if err != nil {
			t.Fatalf("could not find element: %v", err)
		}
		tx, err := elmt.InnerText()
		if err != nil {
			t.Fatalf("could not get text: %v", err)
		}
		st := strings.Split(tx, ".")
		if len(st) != 2 {
			t.Fatalf("could not get score: %v", tx)
		}
		sc, err := strconv.Atoi(st[1])
		if err != nil {
			t.Fatalf("could not convert score to int: %v", err)
		}
		if sc < 7 {
			t.Fatalf("recaptcha score lower than 0.7 \nScore is %v", sc)
		}
		log.Println("recaptcha score: ", tx)
		_ = page.Close()

	})
	//t.Run("UniqueFP", func(t *testing.T) {
	//	page, err := wbbrowser.NewPage()
	//	if err != nil {
	//		t.Fatalf("could not create page: %v", err)
	//	}
	//	_, err = page.Goto("https://amiunique.org/fingerprint")
	//	if err != nil {
	//		t.Fatalf("could not navigate to page: %v", err)
	//	}
	//	time.Sleep(10 * time.Minute)
	//})

	err := wbbrowser.Close()
	if err != nil {
		t.Fatalf("Coudln't close webkit browser")
	}
	err = pw.Stop()
	if err != nil {
		t.Fatalf("Coudln't stop playwright")
	}

}

func TestTurnstileCaptcha(t *testing.T) {

	apikey := os.Getenv("CapMonsterKey")
	proxy := os.Getenv("TestProxy")
	websiteKey := os.Getenv("WebsiteKey")
	websiteURL := os.Getenv("Website")
	if apikey == "" || proxy == "" || websiteKey == "" || websiteURL == "" {
		t.Fatal("OS Envs not set")
	}
	balance, err := CM.GetBalance(apikey)
	if err != nil || balance < 1 {
		t.Fatal(balance, err)
	}
	pxy := strings.Split(proxy, ":")
	pproxy := playwright.Proxy{
		Server:   pxy[0] + ":" + pxy[1],
		Username: &pxy[2],
		Password: &pxy[3],
	}
	captchabrowser, err := ppp.PlaywrightInit(&pproxy, 1, false, pw)
	if err != nil {
		t.Fatal(err)
	}
	page, err := captchabrowser.NewPage()
	if err != nil {
		t.Fatal(err)
	}
	if _, err = page.Goto(websiteURL); err != nil {
		t.Fatal(err)
	}
	if err = CM.HandleTurnstileCookie(apikey, page, websiteKey, pproxy); err != nil {
		t.Fatal(err)
	}
	time.Sleep(10 * time.Minute)

}
