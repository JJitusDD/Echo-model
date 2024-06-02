package response_definition

type Response struct {
	Meta Meta        `json:"meta" mapstructure:"meta"`
	Data interface{} `json:"data,omitempty" mapstructure:"data"`
}

type Meta struct {
	Code    int         `json:"code" mapstructure:"code"`
	Msg     string      `json:"msg" mapstructure:"msg"`
	Message interface{} `json:"message,omitempty" mapstructure:"message"`
}
