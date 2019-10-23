package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

func info() {
	app.Name = "Web Lookup CLI"
	app.Usage = "Lookup IPs, CNAMEs, MX records and Name Servers for a specified host"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "ericwat",
			Email: "ericwat@ericwat.com",
		},
	}
	app.Copyright = "(c) 2019 Eric Watkins"
	app.Version = "1.0.0"
}

func commands() {

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "fullscreeninteractive.com",
			Usage: "domain for the host",
		},
		cli.StringFlag{
			Name:  "url",
			Value: "www.fullscreeninteractive.com",
			Usage: "url for the host",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "nameserver",
			Aliases: []string{"ns"},
			Flags:   flags,
			Usage:   "Looks up the nameserver for the specified host",
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("url"))
				if err != nil {
					return err
				}

				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},
		{
			Name:    "ip",
			Aliases: []string{"ip"},
			Flags:   flags,
			Usage:   "Looks up the IP addresses for the specified host",
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err != nil {
					fmt.Println(err)
				}

				for i := 0; i < len(ip); i++ {
					fmt.Println(ip[i])
				}
				return nil
			},
		},
		{
			Name:    "cname",
			Aliases: []string{"cn"},
			Flags:   flags,
			Usage:   "Looks up the CNAME for the specified host",
			Action: func(c *cli.Context) error {
				cname, err := net.LookupCNAME(c.String("host"))
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(cname)
				return nil
			},
		},
		{
			Name:    "mx",
			Aliases: []string{"mx"},
			Flags:   flags,
			Usage:   "Looks up the MX records for the specified host",
			Action: func(c *cli.Context) error {
				mx, err := net.LookupMX(c.String("host"))
				if err != nil {
					fmt.Println(err)
				}

				for i := 0; i < len(mx); i++ {
					fmt.Println(mx[i].Host, mx[i].Pref)
				}
				return nil
			},
		},
	}
}

func main() {
	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
