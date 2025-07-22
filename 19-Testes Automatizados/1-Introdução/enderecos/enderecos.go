package enderecos

import "strings"

//TipoDeEndereco verifica se o endereço tem um tipo válido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	enderecoTemUmTipoValio := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValio = true
		}
	}

	if enderecoTemUmTipoValio {
		return strings.Title(primeiraPalavraDoEndereco)
	}

	return "Tipo Inválido"
}
