package approverequest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// SendValidationPost envia o código e o timestamp para o servidor,
// acessando o endpoint https://{domain}/approve-request/validate-code.
// Retorna true se a validação for bem-sucedida ou false, juntamente
// com um erro, se ocorrer algum problema.
func SendValidationPost(code, domain string, timestamp int64) (bool, error) {
	// Monta os dados da requisição
	data := map[string]interface{}{
		"code":      code,
		"timestamp": timestamp,
	}

	// Converte os dados para JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return false, fmt.Errorf("erro ao gerar JSON: %v", err)
	}

	// Define a URL do endpoint (domínio e rota)
	url := "https://" + domain + "/approve-request/validate-code"

	// Envia a requisição POST com o JSON gerado
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return false, fmt.Errorf("erro ao enviar requisição: %v", err)
	}
	defer resp.Body.Close()

	// Exibe o status da resposta (opcional)
	fmt.Printf("Status da resposta: %s\n", resp.Status)

	// Decodifica a resposta do servidor (espera um JSON com o campo "valid")
	var result struct {
		Valid bool `json:"valid"`
	}

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, fmt.Errorf("erro ao decodificar a resposta JSON: %v", err)
	}

	return result.Valid, nil
}
