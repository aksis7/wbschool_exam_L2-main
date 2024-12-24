package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

// Поддержка fork/exec
func runForkCommand(args []string) {
	if len(args) == 0 {
		fmt.Println("fork: missing command")
		return
	}

	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		// Для Windows используем cmd.exe
		cmd = exec.Command("cmd", "/C", strings.Join(args, " "))
	} else {
		// Для Unix используем sh
		cmd = exec.Command("sh", "-c", strings.Join(args, " "))
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Start()
	if err != nil {
		fmt.Println("fork error:", err)
		return
	}

	fmt.Printf("Started process with PID: %d\n", cmd.Process.Pid)
	err = cmd.Wait()
	if err != nil {
		fmt.Println("fork process error:", err)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	fmt.Println("Welcome to GoShell. Type \\quit to exit.")

	for {
		// Отображение приглашения
		fmt.Print("GoShell> ")

		// Чтение ввода пользователя
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		// Удаление пробелов и новой строки
		input = strings.TrimSpace(input)

		// Пропускаем пустой ввод
		if input == "" {
			continue
		}

		// Выход из оболочки
		if input == "\\quit" {
			fmt.Println("Exiting GoShell.")
			break
		}

		// Проверяем на пайпы
		if strings.Contains(input, "|") {
			executePipeline(strings.Split(input, "|"))
		} else {
			executeCommand(input)
		}
	}
}

// Выполнение одиночной команды
// Выполнение одиночной команды
func executeCommand(input string) {
	args := strings.Fields(input)
	if len(args) == 0 {
		return
	}

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			fmt.Println("cd: missing argument")
			return
		}
		err := os.Chdir(args[1])
		if err != nil {
			fmt.Println("cd error:", err)
		}
	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println("pwd error:", err)
			return
		}
		fmt.Println(dir)
	case "echo":
		fmt.Println(strings.Join(args[1:], " "))
	case "kill":
		if len(args) < 2 {
			fmt.Println("kill: missing PID")
			return
		}
		pid, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid PID:", args[1])
			return
		}
		killProcess(pid)
	case "ps":
		if runtime.GOOS == "windows" {
			cmd := exec.Command("tasklist")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
		} else {
			cmd := exec.Command("ps", "aux")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
		}
	case "fork":
		runForkCommand(args[1:])
	default:
		runExternalCommand(args)
	}
}

// Завершение процесса
func killProcess(pid int) {
	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println("Failed to find process:", err)
		return
	}

	if runtime.GOOS == "windows" {
		err = process.Kill()
		if err != nil {
			fmt.Println("Failed to kill process on Windows:", err)
		} else {
			fmt.Println("Process", pid, "killed successfully on Windows")
		}
	} else {
		err = process.Signal(syscall.SIGKILL)
		if err != nil {
			fmt.Println("Failed to kill process on UNIX:", err)
		} else {
			fmt.Println("Process", pid, "killed successfully on UNIX")
		}
	}
}

// Выполнение внешней команды
func runExternalCommand(args []string) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing command:", err)
	}
}

// Выполнение пайплайна (конвейер команд)
func executePipeline(commands []string) {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		// Windows: используем cmd.exe /C для выполнения команд
		cmd = exec.Command("cmd", "/C", strings.Join(commands, " | "))
	} else {
		// Linux/MacOS: используем sh -c
		cmd = exec.Command("sh", "-c", strings.Join(commands, " | "))
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		fmt.Println("Pipeline error:", err)
	}
}

