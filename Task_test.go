package playwrightprepack_test

import (
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/playwright-community/playwright-go"
	"playwrightprepack"
)

func init() {
	log.Println("Initializing")
	err := playwright.Install()
	if err != nil {
		log.Fatalf("Error installing playwright: %v", err)
	}
	os.Setenv("HEADLESS", "false")
	log.Println("starting Tests")
}

var pw *playwright.Playwright

func TestMain(m *testing.M) {
	os.Setenv("HEADLESS", "false")
	var err error
	pw, err = playwright.Run()
	if err != nil {
		log.Fatalf("Error running playwright: %v", err)
	}

	code := m.Run()
	pw.Stop()
	os.Exit(code)
}

var wbbrowser playwright.BrowserContext

func TestWebKit(t *testing.T) {
	log.Println("Running WebKit test")
	var errp error
	//proxy := "161.0.70.152:5741:tnzpwplz:156y8h5d4l6q"
	wbbrowser, errp = playwrightprepack.PlaywrightInit("", pw)
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
		page.Close()

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
		time.Sleep(5 * time.Second)
		log.Println("mandatory 5 second sleep while recaptcha loads")

		btn := page.GetByRole("button", playwright.PageGetByRoleOptions{
			Name: "Refresh score now!",
		})
		if btn == nil {
			t.Fatalf("could not find button to refresh captcha score")
		}
		btn.Click()
		log.Println("refreshing captcha score")
		err = page.WaitForLoadState(playwright.PageWaitForLoadStateOptions{
			State: playwright.LoadStateNetworkidle,
		})
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
			t.Fatalf("recaptcha score lower than 0.7")
		}
		log.Println("recaptcha score: ", tx)
		page.Close()

	})

	wbbrowser.Close()

}
