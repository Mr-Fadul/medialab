package gstd

type ElementProperty struct {
	Code        int      `json:"code"`
	Description string   `json:"description"`
	Response    Response `json:"response"`
}
type Param struct {
	Description string `json:"description"`
	Type        string `json:"type"`
	Access      string `json:"access"`
}
type Response struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Param Param  `json:"param"`
}
