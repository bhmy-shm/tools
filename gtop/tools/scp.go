package tools

import (
	"fmt"
	"github.com/bhmy-shm/tools/gtop/pkg"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"os"
	"path"
)

var (
	VarScpUser          string
	VarScpAddress       string
	VarScpSource        string
	VarScpDirect        string
	VarScpRemoteToLocal bool
)

func scpServiceCommand(_ *cobra.Command, args []string) error {

	session, err := pkg.ThreeSSHConnect(VarScpUser, VarScpAddress)
	if err != nil {
		return err
	}
	defer session.Close()

	//远程传送文件到本地
	if VarScpRemoteToLocal {
		return scpRemoteToLocal(session)
	}

	//本地上传到远程
	in, err := session.StdinPipe()
	if err != nil {
		return err
	}
	defer in.Close()

	//读取本地文件
	localFile, err := os.Open(VarScpSource)
	if err != nil {
		panic(err)
	}
	defer localFile.Close()

	fileStat, err := localFile.Stat()
	if err != nil {
		log.Fatal(err)
	}

	//文件大小校验

	//执行session的远程调用命令
	err = session.Start(fmt.Sprintf("/user/bin/scp %s %s", "-t", VarScpDirect))
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintln(in, fileStat.Mode().Perm(), fileStat.Size(), path.Base(localFile.Name()))
	if err != nil {
		log.Fatal(err)
	}

	//拷贝文件
	n, err := io.Copy(in, localFile)
	if err != nil {
		log.Fatal(err)
	}

	//文件结束符号
	_, err = fmt.Fprintln(in, "\x00")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("传输了 %d 字节", n)
	return session.Wait()
}

func scpRemoteToLocal(session *ssh.Session) error {

	//创建本地文件接收远程文件内容
	localFile, err := os.Create(VarScpDirect)
	if err != nil {
		return err
	}
	defer localFile.Close()

	//执行scp命令
	remoteFile := path.Join(VarScpDirect, VarScpSource)         //direct目录，source具体文件
	scpCommand := fmt.Sprintf("/usr/bin/scp -f %s", remoteFile) // -f 将远程主机文件传到到本地主机。与 -t 相反

	// 运行远程 scp 命令
	output, err := session.Output(scpCommand)
	if err != nil {
		return err
	}

	// 将输出写入本地文件
	if _, err = localFile.Write(output); err != nil {
		return err
	}

	// 如果需要，可以在这里验证文件内容（例如通过校验和）

	// 日志记录传输完成
	log.Printf("文件 %s 已成功从 %s 复制到本地 %s", VarScpSource, VarScpAddress, VarScpDirect)

	return nil
}
