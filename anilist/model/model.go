package model

type Data struct {
	Errors []struct {
		Message   string `json:"message"`
		Status    int    `json:"status"`
		Locations []struct {
			Line   int `json:"line"`
			Column int `json:"column"`
		} `json:"locations"`
	} `json:"errors"`
	Data map[string]interface{} `json:"data"`
}
