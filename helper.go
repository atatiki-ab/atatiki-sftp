package main

import (
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func getSftpClient() (*sftp.Client, error) {
	sshConfig := &ssh.ClientConfig{
		User: settings.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(settings.Password),
			ssh.PasswordCallback(func() (string, error) {
				return settings.Password, nil
			}),
			ssh.KeyboardInteractive(func(user, instruction string, questions []string, echos []bool) ([]string, error) {
				// Just send the password back for all questions ;)
				answers := make([]string, len(questions))
				for i := range answers {
					answers[i] = settings.Password
				}

				return answers, nil
			}),
		},
		Timeout: 10 * time.Second,
	}
	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	host := settings.Host
	//fmt.Println(settings)
	client, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return sftp.NewClient(client)
}
