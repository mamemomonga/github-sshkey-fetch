package main

import (
	"log"
	"text/template"
	"bytes"
	"time"
)
func generator(sshkeys, username string) string {
	t := template.Must(template.New("T").Parse(`mkdir -m 0700 -p ~/.ssh
cat >> ~/.ssh/authorized_keys << 'EOS'
#  GitHub User: {{ .Username }}  
#  Created At: {{ .DateTime }}
{{ .SSHKeys }}
EOS
chmod 700 ~/.ssh/authorized_keys
`))
	m := struct {
		SSHKeys  string
		Username string
		DateTime string
	}{
		SSHKeys: sshkeys,
		Username: username,
		DateTime: time.Now().Format("2006-01-02 15:04:06"),
	}
	var buf bytes.Buffer
	err := t.Execute(&buf,m)
	if err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

