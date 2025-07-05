package mail

import (
	"log"

	"github.com/resend/resend-go/v2"
)

type MailClient struct {
	AgencyEmail         string
	InternalAgencyEmail string
	UserEmail           string
	PDFFilePath         string
	APIKey              string
}

type MailService struct {
	mailClient   *MailClient
	resendClient *resend.Client
}

func (m *MailClient) Start() *MailService {
	return &MailService{
		mailClient:   m,
		resendClient: resend.NewClient(m.APIKey),
	}
}

func (m *MailService) sendEmail(from, to string) {
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

func (m *MailService) SendEmails() {
	go m.sendEmail(m.mailClient.InternalAgencyEmail, m.mailClient.AgencyEmail) // to travel agency
	go m.sendEmail(m.mailClient.AgencyEmail, m.mailClient.UserEmail)           // to user
}
