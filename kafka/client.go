package kafka

type Client struct {
	Id string `json:"id"`
	Field1 int64 `json:"field1"`
}

func (c *Client) DoSomething(s string) bool {
	return s == ""
}
