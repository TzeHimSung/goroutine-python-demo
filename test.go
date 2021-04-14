package main

import (
	"context"
	"log"
	"os/exec"
	"strconv"
	"time"
)

var (
	pid = 0
)

type CancelFuncStruct struct {
	CancelFunc context.CancelFunc
}

func TestChannel() {
	// test channel return method
	quitChan := make(chan int)
	go ForeverLoopUseChannel(quitChan)
	time.Sleep(5 * time.Second)
	quitChan <- 1
}

func TestCancelFunc() {
	// test context cancel func return method
	ctx, cancel := context.WithCancel(context.Background())
	go ForeverLoopUseContext(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func TestCancelFuncStruct() {
	// test save cancel func to struct
	bgrCtx := context.Background()
	cancelFuncArray := make([]CancelFuncStruct, 0)

	// call 5 goroutine
	for i := 1; i <= 5; i++ {
		ctx, cancel := context.WithCancel(bgrCtx)
		go ForeverLoopUseContext2(ctx, i)
		cancelFuncArray = append(cancelFuncArray, CancelFuncStruct{
			CancelFunc: cancel,
		})
	}
	time.Sleep(5 * time.Second)
	for _, item := range cancelFuncArray {
		item.CancelFunc()
	}
	time.Sleep(2 * time.Second)
}

func TestPyScript(scriptName string) {
	// test running python script
	// create first goroutine
	ctx, canf := context.WithCancel(context.Background())
	go RunPyScript(ctx, scriptName, 1)
	// create second goroutine
	ctx2, canf2 := context.WithCancel(context.Background())
	go RunPyScript(ctx2, scriptName, 2)

	time.Sleep(2 * time.Second)
	canf()
	time.Sleep(2 * time.Second)
	canf2()
}

func TestExec(scriptName string) {
	// test running and killing python process with PID
	// create process
	pid = RunPyScriptWithCancel(scriptName)
	time.Sleep(10 * time.Second)
	KillProcessByPID(pid)
}

func KillProcessByPID(pid int) {
	cmd := exec.Command("taskkill", "/PID", strconv.Itoa(pid), "/F")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
