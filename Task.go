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

type ProxyStruct struct {
	ip  string
	usr string
	pw  string
}

var IpAgentList = []string{
	//"iPhone 6", "iPhone 6 plus",
	//"iPhone 7", "iPhone 7 plus",
	//"iPhone 8", "iPhone 8 plus",
	"iPhone X", "iPhone XR",
	"iPhone XS", "iPhone XS Max",
	"iPhone 11", "iPhone 11 Pro", "iPhone 11 Pro Max",
	"iPhone SE (2nd generation)",
	"iPhone 12 mini", "iPhone 12", "iPhone 12 Pro", "iPhone 12 Pro Max",
	"iPhone 13 mini", "iPhone 13", "iPhone 13 Pro", "iPhone 13 Pro Max",
	"iPhone 14 mini", "iPhone 14", "iPhone 14 Pro", "iPhone 14 Pro Max",
	"iPhone 15", "iPhone 15 Plus", "iPhone 15 Pro", "iPhone 15 Pro Max",
}

func pp(p string) *playwright.Proxy {
	st := strings.Split(p, ":")
	fmt.Printf(" proxy string %v \n", st)
	var pwps playwright.Proxy
	if len(st) == 4 {
		pwps = playwright.Proxy{
			Server:   st[0] + ":" + st[1],
			Username: &st[2],
			Password: &st[3],
		}
		return &pwps
	}
	return nil
}
func PlaywrightInit(prx string, plt int8, hdl bool, pw *playwright.Playwright) (playwright.BrowserContext, error) {
	var dev = pw.Devices[IpAgentList[rand.Intn(len(IpAgentList)-1)]]
	var platform playwright.Browser
	var err error

	switch plt {
	case 1:
		platform, err = pw.WebKit.Launch(playwright.BrowserTypeLaunchOptions{
			Headless: playwright.Bool(hdl),
		})
		if err != nil {
			return nil, errors.Join(fmt.Errorf("could not launch WebKit browser"), err)
		}
	case 2:
		platform, err = pw.Firefox.Launch(playwright.BrowserTypeLaunchOptions{
			Headless: playwright.Bool(hdl),
		})
		if err != nil {
			return nil, errors.Join(fmt.Errorf("could not launch Firefox browser"), err)
		}
	default:
		platform, err = pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
			Headless: playwright.Bool(hdl),
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
		Proxy:             pp(prx),
		UserAgent:         playwright.String(dev.UserAgent),
		Viewport:          dev.Viewport,
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
