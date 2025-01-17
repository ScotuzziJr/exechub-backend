package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type EmailRecipient struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type EmailSender struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type EmailBody struct {
	Sender 		EmailSender  		`json:"sender"`
	To 			[]EmailRecipient 	`json:"to"`
	Subject 	string				`json:"subject"`
	HTMLContent string				`json:"htmlContent"`
}


func SendEmail(email string, name string) {
	emailBody := EmailBody{
		Sender: EmailSender{
			Name: "ExecHub",	
			Email: "exechubbr@hotmail.com",
		},
		To: []EmailRecipient{
			{
				Email: email,
				Name: name,
			},
		},
		Subject: "Bem-vindo a ExecHub!",
		HTMLContent: `<html>
						<head>
						</head>
						<body>
							<p>Bem-vindo a ExecHub! Estamos felizes em tê-lo conosco.</p>
						</body>
					</html>`,
	
	}

	jsonData, err := json.Marshal(emailBody)

	if err != nil {
		fmt.Printf("Erro ao criar o Json", err)
		return
	}

	fmt.Printf("Json gerado: %s", jsonData)


	req, err := http.NewRequest("POST", "https://api.brevo.com/v3/smtp/email", bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Printf("Erro ao criar a requisição: %v\n", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", "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")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("Erro ao enviar a requisição: %v\n", err)
		return
	}

	if(resp.StatusCode != 201){
		fmt.Printf("EMAIL ENVIADO COM SUCESSO!")
	}
	
	defer resp.Body.Close()

	fmt.Printf("Status %s", resp.Status)
}

