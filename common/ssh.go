package common

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"os/exec"
)

type Config struct {
	Src              Server
	Dest             Server
	PiiDB            bool
	Prefix           bool
	WordPress        bool
	WordPressNetwork bool
}

func RemoteDBBackup(config Config) {
	err := remoteDBBackupPrepare(config)

	if err != nil {
		Error(err.Error())
		os.Exit(1)
	}
}

func remoteDBBackupPrepare(config Config) error {
	//data := populateStruct(srcServer, destServer)
	const bashScript = `
mkdir -p {{.Src.BackupDir}}
cd {{.Src.BackupDir}}
# Check the commands are available
if ! command -v {{.Src.PhpPath}} &> /dev/null
then
    echo "{{.Src.PhpPath}} could not be found"
    exit
fi
if ! command -v {{.Src.N98Path}} &> /dev/null
then
    echo "{{.Src.N98Path}} could not be found"
    exit
fi
{{if .PiiDB}}
{{.Src.PhpPath}} {{.Src.N98Path}} --quiet --no-interaction db:dump --no-tablespaces --compression="gzip" --strip="@log @sessions" --force {{.Src.BackupDir}}/latest-m2.sql.gz
{{- else}}
{{.Src.PhpPath}} {{.Src.N98Path}} --quiet --no-interaction db:dump --no-tablespaces --compression="gzip" --strip="@log @sessions @trade @sales @idx @aggregated @temp @newrelic_reporting $ignore_tables" --force {{.Src.BackupDir}}/latest-m2.sql.gz
{{- end}}
{{if .WordPress}}
cd wp
{{.Src.PhpPath}} {{.Src.N98Path}} --quiet --no-interaction db:dump --no-tablespaces --compression="gzip" --connection="wordpress" --force {{.Src.BackupDir}}/latest-wp.sql.gz
{{- end}}
`
	t := template.Must(template.New("bashScript").Parse(bashScript))
	return t.Execute(os.Stdout, config)
}

func Connect(srcServer Server, payload string) {
	sshRemoteMachine := fmt.Sprintf("%s@%s", srcServer.User, srcServer.Host)
	sshRemotePort := fmt.Sprintf("-p%d", srcServer.Port)
	sshCmd := exec.Command("ssh", sshRemotePort, sshRemoteMachine, payload)
	var out bytes.Buffer
	sshCmd.Stdout = &out
	sshErr := sshCmd.Run()
	if sshErr != nil {
		Error(sshErr.Error())
	}
	fmt.Fprintln(os.Stdout, out.String())
}
