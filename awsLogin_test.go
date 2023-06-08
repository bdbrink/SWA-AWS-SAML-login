package main

import (
	"errors"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/urfave/cli"
)

// Mock implementation of the exec.Cmd struct for testing purposes.
type mockCmd struct {
	command string
	args    []string
}

func (m *mockCmd) SetStdin(stdin *os.File) {
	// Implement as needed for your tests.
}

func (m *mockCmd) SetStdout(stdout *os.File) {
	// Implement as needed for your tests.
}

func (m *mockCmd) SetStderr(stderr *os.File) {
	// Implement as needed for your tests.
}

func (m *mockCmd) Run() error {
	// Mock the Run method based on the test case.
	if m.command == "awssaml get-credentials" {
		// Simulate a successful execution.
		return nil
	}

	// Simulate an error.
	return errors.New("mock command error")
}

func TestAwsLogin(t *testing.T) {
	// Create a test context with the required flag value.
	context := cli.NewContext(nil, nil, nil)
	context.Set("name", "lab")

	// Create an instance of the AWSLogin struct with the test context.
	awsLogin := &AWSLogin{}

	// Replace the execution of the command with a mockCmd implementation.
	oldCmd := exec.Command
	defer func() { exec.Command = oldCmd }()
	exec.Command = func(command string, args ...string) *exec.Cmd {
		return &mockCmd{
			command: command,
			args:    args,
		}
	}

	// Call the AwsLogin method.
	err := awsLogin.AwsLogin(context)

	// Assert the result based on the expected behavior.
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Assert the command execution.
	expectedCommand := "awssaml get-credentials --account-id 217906394988 --name lab --role swa/SWACSDeveloper --user-name e143608 --duration 14400"
	cmd := exec.Command("bash", "-c", expectedCommand)
	mockCmd, ok := cmd.(*mockCmd)
	if !ok {
		t.Errorf("Failed to assert command type")
	}
	if strings.Join(mockCmd.args, " ") != strings.Join(cmd.Args, " ") {
		t.Errorf("Unexpected command arguments: %v", cmd.Args)
	}
}

func TestListRoles(t *testing.T) {
	// Create a test context without any flags.
	context := cli.NewContext(nil, nil, nil)

	// Create an instance of the AWSLogin struct.
	awsLogin := &AWSLogin{}

	// Replace the execution of the command with a mockCmd implementation.
	oldCmd := exec.Command
	defer func() { exec.Command = oldCmd }()
	exec.Command = func(command string, args ...string) *exec.Cmd {
		return &mockCmd{
			command: command,
			args:    args,
		}
	}

	// Call the ListRoles method.
	err := awsLogin.ListRoles(context)

	// Assert the result based on the expected behavior.
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Assert the command execution.
	expectedCommand := "awssaml list"
	cmd := exec.Command("bash", "-c", expectedCommand)
	mockCmd, ok := cmd.(*mockCmd)
	if !ok {
		t.Errorf("Failed to assert command type")
	}
	if strings.Join(mockCmd.args, " ") != strings.Join(cmd.Args, " ") {
		t.Errorf("Unexpected command arguments: %v", cmd.Args)
	}
}
