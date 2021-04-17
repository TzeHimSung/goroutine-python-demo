package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os/exec"
	"syscall"
	"time"
)

func ForeverLoopUseChannel(quitChan <-chan int) {
	for {
		select {
		case <-quitChan:
			fmt.Println("Quit channel received single, function returned.")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("ForeverLoopUseChannel is running")
		}
	}
}

func ForeverLoopUseContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context is done, function returned.")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("ForeverLoopUseContext is running")
		}
	}
}

func ForeverLoopUseContext2(ctx context.Context, n int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context is done, function ", n, " is returned.")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("ForeverLoopUseContext is running with number ", n)
		}
	}
}

func RunPyScript(ctx context.Context, scriptName string, n int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context is done, func ", n, " is returned.")
			return
		default:
			cmd := exec.Command("python", scriptName)
			err := cmd.Run()
			if err != nil {
				panic(err)
			}
			return
		}
	}
}

func RunPyScriptRetPid(scriptName string) (pid int) {
	cmd := exec.Command("python", scriptName)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	pid = cmd.Process.Pid
	fmt.Println("Python process running, pid:", pid)
	return pid
}

func RunPyScriptWithContext(ctx context.Context, scriptName string) {
	cmd := exec.CommandContext(ctx, "python", scriptName)
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	var b bytes.Buffer
	cmd.Stdout = &b
	cmd.Stderr = &b
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
	}
	if err := cmd.Wait(); err != nil {

	}

	pid = cmd.Process.Pid
}
