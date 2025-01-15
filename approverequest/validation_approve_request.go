package approverequest

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type CodeData struct {
	Code      string
	Timestamp int64
}

func GenerateCodeWithTimestamp(secretKey string, durationSeconds int64) CodeData {
	// Calcula o timestamp final: tempo atual (em segundos) + duração informada
	ts := time.Now().Unix() + durationSeconds

	// Cria o HMAC usando SHA256 com a chave secreta e o valor numérico do timestamp convertido para string
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(strconv.FormatInt(ts, 10)))
	code := hex.EncodeToString(h.Sum(nil))

	return CodeData{
		Code:      code,
		Timestamp: ts,
	}
}

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
		err = fmt.Errorf("erro ao gerar JSON: %v", err)
		log.Println(err.Error())
		return false, err
	}

	// Define a URL do endpoint (domínio e rota)
	url := "https://" + domain + "/approve-request/validate-code"

	// Envia a requisição POST com o JSON gerado
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		err = fmt.Errorf("erro ao enviar requisição: %v", err)
		log.Println(err.Error())
		return false, err
	}
	defer resp.Body.Close()

	// Exibe o status da resposta (opcional)
	log.Printf("Status da resposta: %s\n", resp.Status)

	// Decodifica a resposta do servidor (espera um JSON com o campo "valid")
	var result struct {
		Valid bool `json:"valid"`
	}

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		err = fmt.Errorf("erro ao decodificar a resposta JSON: %v", err)
		log.Println(err.Error())
		return false, err
	}

	return result.Valid, nil
}
