package main

import "path/filepath"

const (
	ScriptPath      = "script"
	SleepScriptName = "sleep.py"
	LoopScriptName  = "loop.py"
)

func main() {
	TestExec(filepath.Join(ScriptPath, LoopScriptName))
}
