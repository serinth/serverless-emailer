package api

type Address struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}

type SendEmailRequest struct {
	To      []*Address `json:"to,omitempty"`
	From    *Address   `json:"from,omitempty"`
	CC      []*Address `json:"cc,omitempty"`
	BCC     []*Address `json:"bcc,omitempty"`
	Subject *string   `json:"subject,omitempty"`
	Content *string   `json:"content,omitempty"`
}
