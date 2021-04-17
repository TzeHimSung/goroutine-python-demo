package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	TestProjectName = "testProject"
)

func seeEnv() {
	envList := os.Environ()
	for _, item := range envList {
		fmt.Println(item)
	}
}

func createVirtualEnv() error {
	// create venv dir name
	venvDirName := filepath.Join("testProject", "venv")
	// check whether virtual environment exists or not
	_, err := os.Stat(venvDirName)
	// if exists
	if err == nil {
		fmt.Println("Virtual environment is already exists.")
		return nil
	}
	// create virtual environment
	fmt.Println("Creating virtual environment, please wait.")
	cmd := exec.Command("python", "-m", "venv", venvDirName)
	err = cmd.Run()
	if err != nil {
		return err
	}
	fmt.Println("Virtual environment created successfully.")
	return nil
}

func launch() {
	// create command
	//cmd := exec.Command("python", filepath.Join("testProject", "ds.py"))

	// another command
	pythonFullPath := "C:\\Users\\JHSeng\\go\\src\\goroutine-python-demo\\testProject\\venv\\Scripts\\python.exe"
	cmd := exec.Command(pythonFullPath, filepath.Join("testProject", "ds.py"))

	// create path env
	oldPath := os.Getenv("Path")
	oldPathEnv := "_OLD_VIRTUAL_PATH=" + oldPath
	// remove source python path
	oldPathList := strings.Split(oldPath, ";")
	newPathList := make([]string, 0)
	for _, item := range oldPathList {
		if strings.HasPrefix(item, "C:\\Python38\\") {
			continue
		}
		newPathList = append(newPathList, item)
	}
	fixedOldPath := strings.Join(newPathList, ";")
	cwd, _ := os.Getwd()
	newPath := filepath.Join(cwd, TestProjectName, "venv", "Scripts")
	newPathEnv := "Path=" + newPath + ";" + fixedOldPath

	// create virtual environment env
	venvEnv := "VIRTUAL_ENV=" + filepath.Join(cwd, TestProjectName, "venv")

	// create prompt env
	oldPrompt := "$P$G"
	oldPromptEnv := "_OLD_VIRTUAL_PROMPT=" + oldPrompt
	newPromptEnv := "PROMPT=(venv) " + oldPrompt

	// get source env
	newEnv := os.Environ()
	newEnv = append(newEnv, venvEnv, oldPathEnv, newPromptEnv, oldPromptEnv)

	// fix path env
	pathIdx := 0
	for idx, item := range newEnv {
		// if path or prompt
		if strings.HasPrefix(item, "Path=") {
			pathIdx = idx
			break
		}
	}
	newEnv[pathIdx] = newPathEnv

	// set env
	//cmd.Env = newEnv

	// print env
	//for _, item := range cmd.Env {
	//	fmt.Println(item)
	//}

	// launch script
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

func TestVirtualEnv() {
	// create python virtual environment
	err := createVirtualEnv()
	if err != nil {
		log.Fatal(err)
		return
	}

	// launch ds.py script
	launch()
}
