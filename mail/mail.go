package mail

import (
	"log"

	"github.com/resend/resend-go/v2"
)

type Client struct {
	AgencyEmail         string
	InternalAgencyEmail string
	UserEmail           string
	PDFFilePath         string
	APIKey              string
}

type Service struct {
	mailClient   *Client
	resendClient *resend.Client
}

func (m *Client) Start() *Service {
	return &Service{
		mailClient:   m,
		resendClient: resend.NewClient(m.APIKey),
	}
}

func (m *Service) sendEmail(from, to string) {
	attachment := &resend.Attachment{ // TODO: check if the pdf is less than 40MB
		Path:     m.mailClient.PDFFilePath,
		Filename: "Orçamento.pdf",
	}

	params := &resend.SendEmailRequest{
		From:    from,
		To:      []string{to},
		Subject: "[ORÇAMENTO] Vizela Viagens",
		Html: `
		<p>Prezado,</p>
		<p>Em anexo segue o orçamento solicitado</p>
		<p>Cumprimentos,</p>
		<p>Vizela Viagens</p>
		`,
		Attachments: []*resend.Attachment{attachment},
	}

	_, err := m.resendClient.Emails.Send(params)
	if err != nil {
		log.Println("Failed to send email to", to)
	}
}

func (m *Service) SendEmails() {
	go m.sendEmail(m.mailClient.InternalAgencyEmail, m.mailClient.AgencyEmail) // to travel agency
	go m.sendEmail(m.mailClient.AgencyEmail, m.mailClient.UserEmail)           // to user
}
