package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar() vai retornar a aplicação de linha de comando pronta para ser executada
// os comandos da aplicação vão ser recebidas em app.Commands ele é um slice de comandos. Nesse caso temos dois comandos "ip" e "servidores"
// As flags seria como se fosse o parâmetro que a gente vai passar para esse comando funcionar

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPS e nomes de servidor de internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "hltv.org",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPS de endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome dos servidores na internet",
			Flags:  flags,
			Action: buscarServidor,
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

func buscarServidor(c *cli.Context) {
	host := c.String("host")
	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servidores {
		fmt.Println(server.Host)
	}
}
