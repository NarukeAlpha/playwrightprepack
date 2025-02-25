package CM

type cKey struct {
	ClientKey string `json:"clientKey"`
}

type Solution struct {
	CfClearance string `json:"cf_clearance,omitempty"`
	Token       string `json:"token,omitempty"`
}

type Response struct {
	ErrorId          int      `json:"errorId"`
	ErrorCode        string   `json:"errorCode,omitempty"`
	ErrorDescription string   `json:"errorDescription,omitempty"`
	Balance          float32  `json:"balance,omitempty"`
	TaskId           int32    `json:"taskId,omitempty"`
	Solution         Solution `json:"solution,omitempty"`
	Status           string   `json:"status,omitempty"`
}

type Task struct {
	Type               string `json:"type"`
	WebsiteURL         string `json:"websiteURL"`
	WebsiteKey         string `json:"websiteKey,omitempty"`
	CloudflareTaskType string `json:"cloudflareTaskType,omitempty"`
	HtmlPageBase64     string `json:"htmlPageBase64,omitempty"`
	UserAgent          string `json:"userAgent,omitempty"`
	ProxyType          string `json:"proxyType,omitempty"`
	ProxyAddress       string `json:"proxyAddress,omitempty"`
	ProxyPort          string `json:"proxyPort,omitempty"`
	ProxyLogin         string `json:"proxyLogin,omitempty"`
	ProxyPassword      string `json:"proxyPassword,omitempty"`
	PageAction         string `json:"pageAction,omitempty"`
	PageData           string `json:"pageData,omitempty"`
	Data               string `json:"data,omitempty"`
}

type OuterTask struct {
	ClientKey string `json:"clientKey"`
	Task      Task   `json:"task,omitempty"`
	TaskId    int32  `json:"taskId,omitempty"`
}
