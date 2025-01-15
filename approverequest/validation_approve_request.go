package approverequest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// sendValidationPost envia o código e o timestamp para o servidor
func SendValidationPost(code, domain string, timestamp int64) error {
	// Monta os dados da requisição

	data := map[string]interface{}{
		"code":      code,
		"timestamp": timestamp,
	}

	// Converte os dados para JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("erro ao gerar JSON: %v", err)
	}

	// URL do endpoint (domínio e rota)
	url := "https://" + domain + "/approve-request/validate-code"

	// Envia a requisição POST com o JSON gerado
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("erro ao enviar requisição: %v", err)
	}
	defer resp.Body.Close()

	// Exibe o status da resposta (opcional)
	fmt.Printf("Status da resposta: %s\n", resp.Status)
	return nil
}
