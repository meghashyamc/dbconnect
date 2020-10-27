package dbapi

const errorStatus = "error processing request to db"
const successStatus = "request to db successful"

const (
	scapedData = "scrapedData"
)
const Port = "9050"

type Dbpayload struct {
	Key   string `json:"key"`
	Field string `json:"field"`
	Value string `json:"value"`
}

type apiResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
