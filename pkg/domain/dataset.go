package domain

type Site string

const (
	Brax     Site = "brax"
	Mandrill      = "mandrill"
	Zendesk       = "zendesk"
	HubSpot       = "hubspot"
	Hotjar        = "hotjar"
	Petrus        = "petrus"
)

type Dataset struct {
	SiteId                  string    `json:"siteId"`
	TimeAgo                 int       `json:"TimeAgo"`
	TimeStep                int       `json:"TimeStep"`
	OutliersDetectionMethod string    `json:"OutliersDetectionMethod"`
	MetricesList            []string  `json:"MetricesList"`
	Results                 []float64 `json:"Results"`
	Outliers                []float64 `json:"Outliers"`
}

type Datasets struct {
	Datasets []Dataset `json:"Datasets"`
}
