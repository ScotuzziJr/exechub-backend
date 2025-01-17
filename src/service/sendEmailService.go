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
	var htmlContent = fmt.Sprintf(`
    <body>
        <p>Olá, %s</p>
        <p>Seja muito bem-vindo à Exechub! Estamos extremamente felizes em tê-lo conosco e queremos que saiba que nossa equipe está à disposição para apoiar você em cada passo dessa jornada.</p>
        <p>Na Exechub, acreditamos que juntos podemos alcançar grandes resultados, e estamos ansiosos para colaborar com você, oferecendo as melhores soluções e estratégias para o seu crescimento e sucesso.</p>
        <p>Se tiver qualquer dúvida ou precisar de assistência, não hesite em nos contatar. Estamos aqui para ajudar!</p>
        <p>Mais uma vez, seja bem-vindo à Exechub. Vamos fazer grandes coisas juntos!</p>
        <p>Atenciosamente,</p>
        <p>Equipe Exechub</p>

        <p>Conecte-se conosco nas redes sociais!</p>
        <ul>
            <li><a href="https://www.linkedin.com/company/exechub">LinkedIn</a></li>
            <li><a href="https://twitter.com/exechub">Twitter</a></li>
        </ul>
    </body>`, name)

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
		HTMLContent: htmlContent,
	
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

