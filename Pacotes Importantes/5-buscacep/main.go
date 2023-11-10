package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

// func main() {
// 	for _, cep := range os.Args[1:] {
// 		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")

// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
// 		}

// 		defer req.Body.Close()

// 		res, err := io.ReadAll(req.Body)

// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
// 		}

// 		var data ViaCEP
// 		err = json.Unmarshal(res, &data)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
// 		}

// 		file, err := os.Create("cidade.txt")
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
// 		}
// 		defer file.Close()

//			_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s\n", data.Cep, data.Localidade, data.Uf))
//			fmt.Println("Arquivo criado com sucesso!")
//			fmt.Println(data)
//		}
//	}
func main() {
	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(":8080", nil)
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cep, err := BuscaCep(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cep)
}

func BuscaCep(cep string) (*ViaCEP, error) {
	resp, error := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if error != nil {
		return nil, error
	}
	defer resp.Body.Close()

	body, error := io.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}

	var c ViaCEP
	error = json.Unmarshal(body, &c)
	if error != nil {
		return nil, error
	}

	return &c, nil
}
