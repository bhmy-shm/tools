package tools

import (
	"github.com/bhmy-shm/tools/gtop/pkg"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
)

var (
	VarSSHAddress string
	VarSSHUser    string
)

func sshServiceCommand(_ *cobra.Command, args []string) error {

	var (
		err     error
		session *ssh.Session
	)

	session, err = pkg.ThreeSSHConnect(VarSSHUser, VarSSHAddress)
	if err != nil {
		return err
	}
	defer session.Close()

	session.Stdin = os.Stdin
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	err = session.RequestPty("xterm", 40, 80, pkg.DefaultSSHModels())
	if err != nil {
		log.Fatal(err)
	}

	err = session.Run("sh")
	if err != nil {
		log.Println(err)
	}

	return nil
}
