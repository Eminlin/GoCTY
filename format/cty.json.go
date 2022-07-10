package format

type CtyJson struct {
	Content []CtyJsonContent `json:""`
}

type CtyJsonContent struct {
	Country   string `json:"Country"`
	Continent string `json:"Continent"`
}
