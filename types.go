package main

type Config struct {
	Credentials `json:"credentials"`
	Pairs       `json:"pairs"`
	Slack       `json:"slack"`
}

type Credentials struct {
	ApiKey    string `json:"api_key"`
	ApiSecret string `json:"api_secret"`
}

type Pair struct {
	Symbol    string  `json:"symbol"`
	Threshold float64 `json:"threshold"`
}

type Pairs []Pair

type Slack struct {
	WebHookUrl string `json:"web_hook_url"`
	UserName   string `json:"user_name"`
	IconEmoji  string `json:"icon_emoji"`
}
