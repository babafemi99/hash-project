package models

type Attributes struct {
	TraitType string `json:"trait_type"`
	Value     any    `json:"value"`
	MinValue  int    `json:"min_value,omitempty"`
	MaxValue  int    `json:"max_value,omitempty"`
}

type Data struct {
	ExampleData string `json:"example_data"`
	Hash        string `json:"hash"`
}

type Collections struct {
	Name       string       `json:"name"`
	Id         string       `json:"id"`
	Attributes []Attributes `json:"attributes"`
}

type NFT struct {
	Format           string       `json:"format"`
	ID               string       `json:"id"`
	FileName         string       `json:"file_name"`
	Name             string       `json:"name"`
	Description      string       `json:"description"`
	MintingTools     string       `json:"minting_tools"`
	Gender           string       `json:"gender"`
	SensitiveContent bool         `json:"sensitive_content"`
	SeriesNumber     int          `json:"series_number"`
	SeriesTotal      int          `json:"series_total"`
	Attributes       []Attributes `json:"attributes"`
	Collections      Collections  `json:"collections"`
	Data             Data         `json:"data"`
}
