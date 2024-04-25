package enderecos

import "strings"

// TipoDeEndereço vai ser exportada, entao ela deve começar com letra MAIÚSCULA
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]
	// O strings.Split vai dividir uma string em uma slice de acordo com o separador que nós colocamos em " "
	// JOÃO VICTOR G
	// ["JOÃO" "VICTOR" "G"]

	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}
	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavraDoEndereco) 
	}
	return "Tipo Inválido"
}
