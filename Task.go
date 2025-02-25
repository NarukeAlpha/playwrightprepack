// Package Playwrightprepack pre-sets a lot of the default configurations I usually
// use to scrape or test websites.
package playwrightprepack

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/playwright-community/playwright-go"
)

// PlaywrightInit Returns a browser with predefined settings. one browser per proxy.
//
// Proxy - instantly pass the proxy from the pre-built collection, alternatively can be nil
// platform - 1 for webkit - 2 for firefox - any other number defaults to chromium
// headless - boolean for headless
// playwright - playwright instance to create the browser from
func PlaywrightInit(prx *playwright.Proxy, plt int8, hdl bool, pw *playwright.Playwright) (playwright.BrowserContext, error) {
	var dev = pw.Devices[IpAgentList[rand.Intn(len(IpAgentList)-1)]]
	var platform playwright.Browser
	var err error

	switch plt {
	case 1:
		platform, err = pw.WebKit.Launch(playwright.BrowserTypeLaunchOptions{
			Headless: playwright.Bool(hdl),
			Args:     StealthFlags,
		})
	case 2:
		platform, err = pw.Firefox.Launch(playwright.BrowserTypeLaunchOptions{
			Headless: playwright.Bool(hdl),
			Args:     StealthFlags,
		})
	default:
		StealthFlags = append(StealthFlags, "--disable-gpu") // Chromium-specific
		platform, err = pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
			Headless: playwright.Bool(hdl),
			Args:     StealthFlags,
		})
		if err != nil {
			return nil, errors.Join(fmt.Errorf("could not launch Chromium browser"), err)
		}
	}
	platform, err = pw.WebKit.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	if err != nil {
		return nil, errors.Join(fmt.Errorf("could not launch browser"), err)
	}
	bsr, err := platform.NewContext(playwright.BrowserNewContextOptions{
		//RecordHarContent: playwright.HarContentPolicyAttach,
		//RecordHarMode: playwright.HarModeFull,
		//RecordHarPath: playwright.String("test.har"),
		ColorScheme:       playwright.ColorSchemeDark,
		DeviceScaleFactor: playwright.Float(dev.DeviceScaleFactor),
		HasTouch:          playwright.Bool(dev.HasTouch),
		//IgnoreDefaultArgs: []string{
		//	"--enable-automation",
		//},
		IsMobile:          playwright.Bool(dev.IsMobile),
		JavaScriptEnabled: playwright.Bool(true),
		Proxy:             prx,
		UserAgent:         playwright.String(dev.UserAgent),
		Viewport:          dev.Viewport,
		Permissions: []string{
			"geolocation",
			"notifications",
		},
	})

	if err != nil {
		log.Println(dev)
		return nil, errors.Join(fmt.Errorf("could not launch new context browser"), err)
	}

	script := playwright.Script{
		Content: playwright.String(`
    const defaultGetter = Object.getOwnPropertyDescriptor(
      Navigator.prototype,
      "webdriver"
    ).get;
    defaultGetter.apply(navigator);
    defaultGetter.toString();
    Object.defineProperty(Navigator.prototype, "webdriver", {
      set: undefined,
      enumerable: true,
      configurable: true,
      get: new Proxy(defaultGetter, {
        apply: (target, thisArg, args) => {
          Reflect.apply(target, thisArg, args);
          return false;
        },
      }),
    });
    const patchedGetter = Object.getOwnPropertyDescriptor(
      Navigator.prototype,
      "webdriver"
    ).get;
    patchedGetter.apply(navigator);
    patchedGetter.toString();
  `),
	}
	err = bsr.AddInitScript(script)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("could not add JS script"), err)
	}

	log.Printf("Browser Launched, user agent: %v \n", dev)
	log.Println()
	return bsr, nil
}

// ProxyLoad Returns a Slice of all proxies from a given csv file
func ProxyLoad(dir string) ([]*playwright.Proxy, error) {
	var pps []*playwright.Proxy
	f, err := os.Open(dir)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("could not open file provided \nBad directory?"), err)
	}

	csvr := csv.NewReader(f)
	for i := 0; true; i++ {
		r, err := csvr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, errors.Join(fmt.Errorf("CSV reader failed - err : %v", err), err)
		}
		s := strings.Split(r[0], ":")
		srv := s[0] + ":" + s[1]
		usr := s[2]
		pss := s[3]

		var p = playwright.Proxy{
			Server:   srv,
			Username: &usr,
			Password: &pss,
		}
		pps = append(pps, &p)

	}
	err = f.Close()
	if err != nil {
		log.Fatalf("failed to close file - err: %v", err)
	}
	return pps, nil
}
