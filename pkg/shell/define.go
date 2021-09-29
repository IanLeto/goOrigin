package shell

import (
	"goOrigin/config"
)

var SSHConn *SSHSession

func InitSSH() {
	SSHConn = NewSSHClient(config.Conf)
}
