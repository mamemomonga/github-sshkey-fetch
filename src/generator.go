package main

import (
	"log"
	"text/template"
	"bytes"
	"time"
	"strings"
//	"github.com/davecgh/go-spew/spew"
)

func genKeys(sshkeys []PubKeyT) string {
	type TplT struct {
		Pubkey   string
		Username string
		DateTime string
	}
	tp := template.Must(template.New("T").Parse(`{{ range .}}# -----------------------------------
#  GitHub User: {{ .Username }}  
#  Created At: {{ .DateTime }}
{{ .Pubkey }}
# -----------------------------------

{{ end }}`))
	var m []TplT
	for _,v := range(sshkeys) {
		m = append(m,TplT{
			Pubkey:   strings.TrimSpace(v.Pubkey),
			Username: v.Username,
			DateTime: time.Now().Format("2006-01-02 15:04:06"),
		})
	}

	var buf bytes.Buffer
	err := tp.Execute(&buf,m)
	if err != nil {
		log.Fatal(err)
	}

	return buf.String()
}

func genAuthKeys(sshkeys []PubKeyT) string {
	pubkeys := genKeys(sshkeys)
	t := template.Must(template.New("T").Parse(`mkdir -m 0700 -p ~/.ssh
cat >> ~/.ssh/authorized_keys << 'EOS'
{{ . }}EOS
chmod 600 ~/.ssh/authorized_keys
`))
	var buf bytes.Buffer
	err := t.Execute(&buf,pubkeys)
	if err != nil {
		log.Fatal(err)
	}
	return buf.String()
}
