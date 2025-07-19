package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação da Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidores na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome dos servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}
	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}

//Terminal:
//go run main.go ip --host amazon.com.br

//Retorno:
//Ponto de partida
//72.21.203.171
//52.94.225.243
//54.239.26.87

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host) //NS: name server
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor)
	}
}

//Terminal:
//go run main.go servidores --host amazon.com.br

//Retorno:
//Ponto de partida
//&{ns1.amzndns.org.}
//&{ns2.amzndns.co.uk.}
//&{ns2.amzndns.org.}
//&{ns1.amzndns.co.uk.}
//&{ns1.amzndns.net.}
//&{ns2.amzndns.com.}
//&{ns2.amzndns.net.}
//&{ns1.amzndns.com.}
