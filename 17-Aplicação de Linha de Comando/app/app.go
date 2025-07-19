package app

import "github.com/urfave/cli"

//Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação da Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidores na internet"

	return app
}
