package mail

import (
	"log"
	"os"

	"github.com/resend/resend-go/v2"
)

type Client struct {
	AgencyEmail         string
	InternalAgencyEmail string
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

func (m *Service) sendEmail(from, to, pdfFilePath string) {
	f, err := os.ReadFile(pdfFilePath)
	log.Println(len(f))
	if err != nil {
		log.Println(err)
		return
	}
	attachment := &resend.Attachment{ // TODO: check if the pdf is less than 40MB
		Content:  f,
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

	_, err = m.resendClient.Emails.Send(params)
	if err != nil {
		log.Println("Failed to send email to", to, ":", err)
	}
}

func (m *Service) SendEmails(pdfFilePath, userEmail string) {
	go m.sendEmail(m.mailClient.InternalAgencyEmail, m.mailClient.AgencyEmail, pdfFilePath) // to travel agency
	go m.sendEmail(m.mailClient.AgencyEmail, userEmail, pdfFilePath)                        // to user
}
