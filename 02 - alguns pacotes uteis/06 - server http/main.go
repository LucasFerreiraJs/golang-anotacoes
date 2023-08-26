package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ViaCep struct {
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

func main() {
	// startando um server
	// manipilando header

	// podemos adicionar ou pegar informações dos headers

	http.HandleFunc("/", BuscaCEPHandler)
	http.ListenAndServe(":8080", nil)
}

func BuscaCEPHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
	}

	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cep, error := BuscaCep(cepParam)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// * associa com a response
	json.NewEncoder(w).Encode(cep)

	// seria a mesma coisa que
	/* result, err := json.Marshal(cep)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(result)
	*/
	// w.Write([]byte("Hello World"))

}

func BuscaCep(cep string) (*ViaCep, error) {

	resp, error := http.Get("https://viacep.com.br/ws/" + cep + "/json")
	if error != nil {
		return nil, error
	}
	defer resp.Body.Close()

	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}

	var c ViaCep
	error = json.Unmarshal(body, &c) // &apontamento
	if error != nil {

		return nil, error
	}

	return &c, nil
}
