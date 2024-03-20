// vash (like a bash, but vash)
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)

	for {
		inputInvitation(getWd())

		input, err := scanner.ReadString('\n')
		if err != nil {
			fmt.Printf("vash: %s\n", err.Error())
		}

		cmd := strings.Fields(input)

		if len(cmd) < 1 {
			continue
		}

		switch cmd[0] {
		case "cd":
			changeDirectory(cmd[1:])
		case "echo":
			echo(cmd[1:])
		case "ps":
			ps()
		case "pwd":
			printWorkingDirectory()
		case "kill":
			kill(cmd[1:])
		case "q":
			os.Exit(0)
		default:
			handleUnexpected(cmd)
		}
	}
}

func printWorkingDirectory() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("vash: %s\n", err.Error())
	}
	fmt.Println(wd)
}

func echo(args []string) {
	for _, arg := range args {
		fmt.Printf("%s ", arg)
	}
	fmt.Println()
}

func ps() {
	processes, err := exec.Command("ps", "-a").Output()
	if err != nil {
		fmt.Printf("vash: %s\n", err.Error())
	}
	fmt.Print(string(processes))
}

func kill(args []string) {
	_, err := exec.Command("kill", args...).Output()
	if err != nil {
		fmt.Printf("vash: %s\n", err.Error())
	}

	fmt.Println("killed")
}

func changeDirectory(args []string) {
	if len(args) > 1 {
		fmt.Println("vash: too many arguments for cd")
	}

	if err := os.Chdir(args[0]); err != nil {
		fmt.Printf("vash: %s\n", err.Error())
	}
}

func inputInvitation(currentDirectory string) {
	userName, err := exec.Command("whoami").Output()
	if err != nil {
		fmt.Printf("vash: %s\n", err.Error())
	}
	fmt.Printf("%s#%s : ", string(userName[:len(userName)-1]), currentDirectory)
}

func getWd() string {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("vash: cannot get working directory")
		return ""
	}

	return wd
}

func handleUnexpected(cmd []string) {
	var command string
	for _, arg := range cmd {
		command += arg
		command += " "
	}
	command = command[:len(command)-1]
	//fmt.Println(command)

	com := exec.Command(command)
	com.Stdout = os.Stdout
	com.Stderr = os.Stderr
	if err := com.Run(); err != nil {
		fmt.Printf("vash: %s\n", err.Error())
	}
}
