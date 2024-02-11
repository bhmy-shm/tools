package pkg

import (
	"fmt"
	"github.com/bhmy-shm/tools/gtop/common"
	"github.com/bhmy-shm/tools/gtop/common/errorx"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"net"
	"syscall"
)

const (
	SSHUser     = "root"
	SSHPassword = "shm19990131."
	SSHIp       = "49.235.156.213"
	SSHPort     = 22
	BuildScript = "sh"
)

func DefaultSSHModels() ssh.TerminalModes {
	return ssh.TerminalModes{
		ssh.ECHO:          0, //ssh 命令回显开关，1开启，0关闭
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
}

func connect(user, password, host string, port int) (*ssh.Session, error) {
	var (
		err    error
		client *ssh.Client
		auth   []ssh.AuthMethod
	)

	//get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))
	hostKeyCallBack := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}

	clientConfig := &ssh.ClientConfig{
		User:            user,
		Auth:            auth,
		HostKeyCallback: hostKeyCallBack,
	}

	addr := fmt.Sprintf("%s:%d", host, port)
	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	return client.NewSession()
}

func ThreeSSHConnect(user, host string) (*ssh.Session, error) {
	var (
		count          uint8
		err            error
		VarSSHPassword []byte
		session        *ssh.Session
	)

	ip, port := common.AddressSplitIpPort(host)

	//支持从终端输入3次密码
	for count < 3 {

		count++

		fmt.Print("password: ")

		VarSSHPassword, err = terminal.ReadPassword(syscall.Stdin)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("")
		}

		session, err = connect(user, string(VarSSHPassword), ip, port)
		if err != nil {
			log.Fatal(err)
		} else {
			break
		}
	}
	if session == nil {
		return nil, errorx.ErrCodeSSHConnectionThreeLimit
	}
	return session, nil
}
