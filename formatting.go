package ppp

// IpAgentList contains all iphones from iphone X to 16.
var IpAgentList = []string{
	//"iPhone 6", "iPhone 6 plus",
	//"iPhone 7", "iPhone 7 plus",
	//"iPhone 8", "iPhone 8 plus",
	//"iPhone X", "iPhone XR",
	//"iPhone XS", "iPhone XS Max",
	"iPhone 11", "iPhone 11 Pro", "iPhone 11 Pro Max",
	//"iPhone SE (2nd generation)",
	"iPhone 12 mini", "iPhone 12", "iPhone 12 Pro", "iPhone 12 Pro Max",
	"iPhone 13 mini", "iPhone 13", "iPhone 13 Pro", "iPhone 13 Pro Max",
	"iPhone 14 mini", "iPhone 14", "iPhone 14 Pro", "iPhone 14 Pro Max",
	"iPhone 15", "iPhone 15 Plus", "iPhone 15 Pro", "iPhone 15 Pro Max",
	//"iPhone 16", "iPhone 16 Pro", "iPhone 16 Pro Max",
}

var StealthFlags = []string{
	"--incognito",
	"--accept-lang=en-US",
	"--lang=en-US",
	"--no-pings",
	"--mute-audio",
	"--no-first-run",
	"--no-default-browser-check",
	"--disable-cloud-import",
	"--disable-gesture-typing",
	"--disable-offer-store-unmasked-wallet-cards",
	"--disable-offer-upload-credit-cards",
	"--disable-print-preview",
	"--disable-voice-input",
	"--disable-wake-on-wifi",
	"--disable-cookie-encryption",
	"--ignore-gpu-blocklist",
	"--enable-async-dns",
	"--enable-simple-cache-backend",
	"--enable-tcp-fast-open",
	"--prerender-from-omnibox=disabled",
	"--enable-web-bluetooth",
	"--disable-features=AudioServiceOutOfProcess,IsolateOrigins,site-per-process,TranslateUI,BlinkGenPropertyTrees",
	"--aggressive-cache-discard",
	"--disable-extensions",
	"--disable-ipc-flooding-protection",
	"--disable-blink-features=AutomationControlled",
	"--test-type",
	"--enable-features=NetworkService,NetworkServiceInProcess,TrustTokens,TrustTokensAlwaysAllowIssuance",
	"--disable-component-extensions-with-background-pages",
	"--disable-default-apps",
	"--disable-breakpad",
	"--disable-component-update",
	"--disable-domain-reliability",
	"--disable-sync",
	"--disable-client-side-phishing-detection",
	"--disable-hang-monitor",
	"--disable-popup-blocking",
	"--disable-prompt-on-repost",
	"--metrics-recording-only",
	"--safebrowsing-disable-auto-update",
	"--password-store=basic",
	"--autoplay-policy=no-user-gesture-required",
	"--use-mock-keychain",
	"--force-webrtc-ip-handling-policy=disable_non_proxied_udp",
	"--webrtc-ip-handling-policy=disable_non_proxied_udp",
	"--disable-session-crashed-bubble",
	"--disable-crash-reporter",
	"--disable-dev-shm-usage",
	"--force-color-profile=srgb",
	"--disable-translate",
	"--disable-background-networking",
	"--disable-background-timer-throttling",
	"--disable-backgrounding-occluded-windows",
	"--disable-infobars",
	"--hide-scrollbars",
	"--disable-renderer-backgrounding",
	"--font-render-hinting=none",
	"--disable-logging",
	"--enable-surface-synchronization",
	"--run-all-compositor-stages-before-draw",
	"--disable-threaded-animation",
	"--disable-threaded-scrolling",
	"--disable-checker-imaging",
	"--disable-new-content-rendering-timeout",
	"--disable-image-animation-resync",
	"--disable-partial-raster",
	"--blink-settings=primaryHoverType=2,availableHoverTypes=2,primaryPointerType=4,availablePointerTypes=4",
	"--disable-layer-tree-host-memory-pressure",
}
