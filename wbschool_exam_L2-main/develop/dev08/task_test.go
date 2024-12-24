package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
)

// Вспомогательная функция для выполнения команд
func runShellCommand(command string) (string, error) {
	cmd := exec.Command("./goshell")
	cmd.Stdin = strings.NewReader(command + "\n\\quit\n")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	return out.String(), err
}

// Тест команды pwd
func TestPwd(t *testing.T) {
	output, err := runShellCommand("pwd")
	if err != nil {
		t.Fatalf("Failed to run pwd command: %v", err)
	}

	if !strings.Contains(output, string(os.PathSeparator)) {
		t.Errorf("Expected pwd output to contain path separator, got: %s", output)
	}
}

// Тест команды echo
func TestEcho(t *testing.T) {
	output, err := runShellCommand("echo Hello, World!")
	if err != nil {
		t.Fatalf("Failed to run echo command: %v", err)
	}

	expected := "Hello, World!"
	if !strings.Contains(output, expected) {
		t.Errorf("Expected echo output to be '%s', got: %s", expected, output)
	}
}

// Тест команды cd
func TestCd(t *testing.T) {
	tempDir := os.TempDir()
	command := "cd " + tempDir + "\npwd"
	output, err := runShellCommand(command)
	if err != nil {
		t.Fatalf("Failed to run cd command: %v", err)
	}

	if !strings.Contains(output, tempDir) {
		t.Errorf("Expected pwd to return temp directory '%s', got: %s", tempDir, output)
	}
}

// Тест команды ps
func TestPs(t *testing.T) {
	output, err := runShellCommand("ps")
	if err != nil {
		t.Fatalf("Failed to run ps command: %v", err)
	}

	expected := "PID"
	if !strings.Contains(output, expected) {
		t.Errorf("Expected ps output to contain '%s', got: %s", expected, output)
	}
}

// Тест команды fork
func TestFork(t *testing.T) {
	output, err := runShellCommand("fork echo Forked Process")
	if err != nil {
		t.Fatalf("Failed to run fork command: %v", err)
	}

	expected := "Forked Process"
	if !strings.Contains(output, expected) {
		t.Errorf("Expected fork output to contain '%s', got: %s", expected, output)
	}
}

// Тест пайпов (pipeline)
func TestPipeline(t *testing.T) {
	output, err := runShellCommand("echo Hello | findstr Hello")
	if err != nil {
		t.Fatalf("Failed to run pipeline command: %v", err)
	}

	expected := "Hello"
	if !strings.Contains(output, expected) {
		t.Errorf("Expected pipeline output to contain '%s', got: %s", expected, output)
	}
}

// Тест команды kill (только проверка синтаксиса)
func TestKill(t *testing.T) {
	output, err := runShellCommand("kill 1")
	if err != nil {
		t.Fatalf("Failed to run kill command: %v", err)
	}

	expected := "Failed to find process"
	if !strings.Contains(output, expected) {
		t.Errorf("Expected kill to fail with message '%s', got: %s", expected, output)
	}
}
