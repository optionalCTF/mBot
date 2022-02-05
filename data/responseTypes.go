package data

var MissionResponse []string = []string{"structuredResponse", "introduction", "testing_methodology", "conclusion"}
var SVOptions []string = []string{"yes", "vulnerable", "out_of_threshold", "not_exploitable", "n/a"}
var MissionOptions []string = []string{"yes", "no", "n/a"}

type Mission struct {
	StructuredResponse string `json:"structuredResponse"`
	Introduction       string `json:"introduction"`
	TestingMethodology string `json:"testing_methodology"`
	Conclusion         string `json:"conclusion"`
}
