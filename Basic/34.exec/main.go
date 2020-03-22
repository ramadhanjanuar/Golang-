package main

import "os/exec"
import "fmt"

func main() {
	var output1, _ = exec.Command("ls").Output()

	fmt.Println("LS: ", string(output1))

	var output2, _ = exec.Command("pwd").Output()
	fmt.Println("PWD: ", string(output2))

	var output3, _ = exec.Command("git", "config", "user.name").Output()
	fmt.Println("Git Name:", string(output3))

	if runtime.GOOS == "windows" {
		output, err = exec.Command("cmd", "/C", "git config user.name").Output()
	} else {
		output, err = exec.Command("bash", "-c", "git config user.name")
	}
}
