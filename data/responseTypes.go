package data

import "encoding/json"

var MissionResponse []string = []string{"structuredResponse", "introduction", "testing_methodology", "conclusion"}
var SVOptions []string = []string{"yes", "vulnerable", "out_of_threshold", "not_exploitable", "n/a"}
var MissionOptions []string = []string{"yes", "no", "n/a"}

type Mission struct {
	StructuredResponse string `json:"structuredResponse"`
	Introduction       string `json:"introduction"`
	TestingMethodology string `json:"testing_methodology"`
	Conclusion         string `json:"conclusion"`
}

type ConnectionStatus struct {
	ConnectedBy           string          `json:"connected_by"`
	Status                string          `json:"status"`
	Slug                  string          `json:"slug"`
	Env                   string          `json:"env"`
	PendingSlug           string          `json:"pending_slug"`
	ListingIsRegisterable json.RawMessage `json:"listing_is_registerable"`
	ID                    int             `json:"id"`
	GatewayConnected      bool            `json:"gateway_connected"`
	IP                    string          `json:"ip_address"`
	EnvType               string          `json:"env_type"`
}
