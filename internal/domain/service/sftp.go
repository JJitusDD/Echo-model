package service

import (
	"fmt"
	"strings"

	"echo-model/internal/domain/model/request"
	"echo-model/internal/domain/utilties"
	"echo-model/pkg/agent/ssh"
	"github.com/labstack/echo/v4"
	"github.com/pkg/sftp"
)

// @Param request body request.SftpReadPathReq true "query params"
// @Success 200 {object} response.Response
// @Router /sftp/read-path [post]
// @tags SFTP
func (s *Service) SftpReadPath(c echo.Context) (err error) {
	body := new(request.SftpReadPathReq)
	if err := utilties.BindingBody(body, c); err != nil {
		return err
	}

	sshClient, err := ssh.NewSshClient(
		s.Config.SFTP.User,
		s.Config.SFTP.Addr,
		s.Config.SFTP.Port,
		"./private_key.pem",
		"pass-ssh",
	)
	if err != nil {
		s.Logger.WithError(err).Errorf("ssh.NewSshClient")
		return err
	}

	conn, _ := sshClient.Connect()

	// open an SFTP session over an existing ssh connection.
	client, err := sftp.NewClient(conn)
	if err != nil {
		fmt.Printf("Failed to create client: %s", err)
		return
	}
	defer client.Close()

	LoopFile(client, body.Dir)

}

func LoopFile(client *sftp.Client, remoteDir string) {
	files, err := client.ReadDir(remoteDir)
	if err != nil {
		fmt.Printf("Unable to list remote dir: %v", err)
	}

	for _, f := range files {
		var name string
		name = f.Name()
		if f.IsDir() {
			name = name + "/"
			LoopFile(client, remoteDir+name)
		} else if strings.Contains(name, ".xlsx") {
			fmt.Println(remoteDir + name)
		}
	}
}
