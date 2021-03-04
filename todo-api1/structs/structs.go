package structs

//Response is an struct
type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}
