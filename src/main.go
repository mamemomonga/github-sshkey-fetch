package main

import (
	"os"
	"fmt"
	"log"
	"flag"
	"strings"
	"github.com/mamemomonga/github-sshkey-fetch/src/buildinfo"
)

type PubKeyT struct {
	Username string
	Pubkey   string
}

func main() {
	var (
		flg_u = flag.String("u","","GitHub usernames(comma separated)")
		flg_p = flag.Bool("g",false,"Generator")
		flg_v = flag.Bool("v",false,"Version")
	)
	flag.Parse()

	if(*flg_v) {
		fmt.Printf("github-sshkey-fetch %s-%s\n",buildinfo.Version, buildinfo.Revision)
		os.Exit(0)
	}

	if(*flg_u == "") {
		flag.PrintDefaults()
		return
	}

	usernames := strings.Split(*flg_u,",")

	var pubkeys []PubKeyT
	for _,u := range(usernames) {
		k,err := fetch(fmt.Sprintf("https://github.com/%s.keys",u))
		if err != nil {
			log.Fatal(err)
		}
		pubkeys = append(pubkeys, PubKeyT{ Username: u, Pubkey: k })
	}

	if(*flg_p) {
		fmt.Println(genAuthKeys(pubkeys))
	} else {
		fmt.Println(genKeys(pubkeys))
	}
}
