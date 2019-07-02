package viewmodels

type HttpServiceResponse struct {
	Messages *MessagesResponse `json:"messages,omitempty"`
	Data     interface{}       `json:"data,omitempty"`
}

type ErrorMessageResponse struct {
	ErrorMessages MessagesResponse `json:"error_messages"`
}

type MessagesResponse struct {
	MessageEn string `json:"en,omitempty"`
	MessageId string `json:"id,omitempty"`
}
