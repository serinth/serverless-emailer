package api

type SendEmailRequest struct {
	To      *string   `json:"to,omitempty"`
	From    *string   `json:"from,omitempty"`
	CC      []*string `json:"cc,omitempty"`
	BCC     []*string `json:"bcc,omitempty"`
	Subject *string   `json:"subject,omitempty"`
	Content *string   `json:"content,omitempty"`
}
