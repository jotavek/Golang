package main

import "fmt"

func main() {
	usuario := map[string]map[string]string{
		"pessoa": {
			"nome":      "João",
			"sobrenome": "Victor",
			"idade":     "22",
		},
		"escolaridade": {
			"escola": "IACI",
		},
	}
	// Estrutura do Delete é: delete(mapa, chave)

	delete(usuario, "escolaridade")
	fmt.Println(usuario)
	usuario["cor do cabelo"] = map[string]string{
		"cor": "preto",
	}
	fmt.Println(usuario)
}
