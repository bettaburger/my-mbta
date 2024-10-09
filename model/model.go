package model

type Alert struct {
	ActivePeriod []struct {
		End   string `json:"end"`
		Start string `json:"start"`
	} `json:"active_period"`
	Attributes struct {
		Banner               interface{} `json:"banner"`
		Cause                string      `json:"cause"`
		CreatedAt            string      `json:"created_at"`
		Description          interface{} `json:"description"`
		DurationCertainty    string      `json:"duration_certainty"`
		Effect               string      `json:"effect"`
		Header               string      `json:"header"`
		Image                interface{} `json:"image"`
		ImageAlternativeText interface{} `json:"image_alternative_text"`
		InformedEntity       []struct {
			Activities []string `json:"activities"`
			Route      string   `json:"route"`
			RouteType  int      `json:"route_type"`
		} `json:"informed_entity"`
		Lifecycle     string      `json:"lifecycle"`
		ServiceEffect string      `json:"service_effect"`
		Severity      int         `json:"severity"`
		ShortHeader   string      `json:"short_header"`
		Timeframe     interface{} `json:"timeframe"`
		UpdatedAt     string      `json:"updated_at"`
		URL           interface{} `json:"url"`
	} `json:"attributes"`
	ID    string `json:"id"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Type string `json:"type"`
}

type Data struct {
	Data [] Alert `json:"data"`
}