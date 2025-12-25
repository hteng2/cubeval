package main

func main() {
	cube := InitCube()
	
	cmdHandler := MakeCmdHandler(&cube)
	
	t := InitTerm(&cmdHandler)
	
	defer t.Restore()
	
	t.Loop()
}
