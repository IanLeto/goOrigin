package shell

import (
	"github.com/sirupsen/logrus"
	"goOrigin/config"
	"golang.org/x/crypto/ssh"
)

type SSHSession struct {
	*ssh.Session
	Target string
}

func NewSSHClient(conf *config.Config) *SSHSession {

	client, err := ssh.Dial("tcp", conf.SSH.Address, &ssh.ClientConfig{
		Config:            ssh.Config{},
		User:              conf.SSH.User,
		Auth:              []ssh.AuthMethod{ssh.Password(conf.SSH.Auth)},
		HostKeyCallback:   nil,
		BannerCallback:    nil,
		ClientVersion:     "",
		HostKeyAlgorithms: nil,
		Timeout:           10,
	})
	if err != nil {
		logrus.Errorf("run ssh client failed %s", err)
	}
	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		logrus.Errorf("create session fail %s", err)
	}
	return &SSHSession{session, conf.SSH.Address}
}

func (s *SSHSession) Exec(content string) ([]byte, error) {
	res, err := s.CombinedOutput(content)
	logrus.Debugf("exec %s fail at %s", err, s.Target)
	return res, err
}
