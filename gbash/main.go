package main

import (
	"fmt"
	"gbash/lib"
	"gbash/lib/funcs"
	"gbash/lib/vars"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("File must be specified")
	}

	ScriptPath := os.Args[1]

	go SignalWatch()

	lib.NewBash(ScriptPath, lib.BASHPath).
		Used(
			vars.NewMemInfo().Info(),
			funcs.NewString().Upper(),
			funcs.NewBytesUnitParse().UnitParse()).
		Run()
}

func SignalWatch() {
	sigs := make(chan os.Signal, 1)
	// 创建一个通道用于在处理函数中通知程序可以退出
	done := make(chan bool, 1)

	// 注册想要接收的系统信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		select {
		case <-sigs:
			defer os.Exit(0)
			fmt.Println("Signal received, shutting down...")
			time.Sleep(1 * time.Second)
			done <- true // 发送信号，表示处理完毕
		}
	}()

	<-done // 等待接收处理函数中的信号
	fmt.Println("Program has exited")

	os.Exit(-1)
}
