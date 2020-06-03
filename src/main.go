package main

import (
	"os"
	"fmt"
	"log"
	"flag"
	"github.com/mamemomonga/github-sshkey-fetch/src/buildinfo"
)

func main() {

	var (
		flg_u = flag.String("u","","GitHub Username")
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
	username := *flg_u
	log.Printf("Username: %s",username)
	sshkey,err := fetch(fmt.Sprintf("https://github.com/%s.keys",username))
	if err != nil {
		log.Fatal(err)
	}
	if(*flg_p) {
		fmt.Println(generator(sshkey,username))

	} else {
		fmt.Println(sshkey)

	}
}


