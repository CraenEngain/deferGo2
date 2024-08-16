package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("my_shell> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		commands := strings.Split(input, "|")

		if len(commands) > 1 {
			handlePipelines(commands)
		} else {
			executeCommand(strings.Fields(input))
		}
	}
}

func executeCommand(args []string) {
	if len(args) == 0 {
		return
	}

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			fmt.Println("cd: argument required")
			return
		}
		if err := os.Chdir(args[1]); err != nil {
			fmt.Println("cd error:", err)
		}

	case "pwd":
		if dir, err := os.Getwd(); err == nil {
			fmt.Println(dir)
		} else {
			fmt.Println("pwd error:", err)
		}

	case "echo":
		fmt.Println(strings.Join(args[1:], " "))

	case "kill":
		if len(args) < 2 {
			fmt.Println("kill: argument required")
			return
		}
		pid, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("kill error:", err)
			return
		}
		terminateProcess(pid)
	case "ps":
		cmd := exec.Command("ps")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("ps error:", err)
		}

	default:
		externalCommand(args)
	}
}

func handlePipelines(commands []string) {
	var cmds []*exec.Cmd

	for i, command := range commands {
		args := strings.Fields(strings.TrimSpace(command))
		if len(args) == 0 {
			continue
		}

		cmd := exec.Command(args[0], args[1:]...)
		if i > 0 {
			// Set the input of the current command to the output of the previous one
			cmd.Stdin, _ = cmds[i-1].StdoutPipe()
		}
		if i < len(commands)-1 {
			// The output of the current command goes to the input of the next one
			cmd.Stdout = nil
		} else {
			// Last command in pipeline outputs to stdout
			cmd.Stdout = os.Stdout
		}
		cmd.Stderr = os.Stderr

		if err := cmd.Start(); err != nil {
			fmt.Println("Error starting command:", err)
			return
		}

		cmds = append(cmds, cmd)
	}

	for _, cmd := range cmds {
		cmd.Wait()
	}
}

func externalCommand(args []string) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Printf("Command execution error: %s\n", err)
	}
}

func terminateProcess(pid int) {
	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Printf("kill error: %v\n", err)
		return
	}
	if err := proc.Kill(); err != nil {
		fmt.Printf("kill error: %v\n", err)
	} else {
		fmt.Printf("Process %d killed\n", pid)
	}
}
