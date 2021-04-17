package main

const (
	ScriptPath      = "script"
	SleepScriptName = "sleep.py"
	LoopScriptName  = "loop.py"
)

func main() {
	//TestExecCommandWithContext(filepath.Join(ScriptPath, LoopScriptName))
	TestVirtualEnv()
}
