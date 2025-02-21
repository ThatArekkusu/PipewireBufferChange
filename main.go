package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var clockChoice string
	var clockRate int

	fmt.Println("Do you want to set clock rate? (y/n)")
	fmt.Scanln(&clockChoice)

	if clockChoice == "y" {
		fmt.Println("Enter the clock rate you want to set the system to:")

		// Define the clock rate the user wants their system to be set to.
		var err error

		_, err = fmt.Scanln(&clockRate)
		if err != nil {
			fmt.Println("Error:", err)

			// Fifth command
			cmd := exec.Command(
				"sed",
				"-i",
				"s/default.clock.rate[[:space:]]*=[[:space:]]*[0-9]*/default.clock.rate = "+fmt.Sprintf("%d", clockRate)+"/",
				".config/pipewire/pipewire.conf")

			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Println(string(out))

			cmd1 := exec.Command(
				"sed",
				"-i",
				"s/default.clock.allowed-rates[[:space:]]*=[[:space:]]*\\[[[:space:]]*[0-9]+[[:space:]]*\\]/default.clock.allowed-rates = \\["+fmt.Sprintf("%d", clockRate)+"\\]/",
				".config/pipewire/pipewire.conf")

			out1, err := cmd1.CombinedOutput()
			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Println(string(out1))

			return
		}
	}

	fmt.Println("Enter the buffer size you want to set the system to:")

	// Define the buffer size the user wants their system to be set to.
	var bufferSize int
	_, err := fmt.Scanln(&bufferSize)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Buffer size "default.clock.min-quantum" needs to be half the buffer size.
	bufferSizeMinQuantum := bufferSize / 2

	// First command
	cmd1 := exec.Command(
		"sed",
		"-i",
		"s/default.clock.quantum[[:space:]]*=[[:space:]]*[0-9]*/default.clock.quantum = "+fmt.Sprintf("%d", bufferSize)+"/",
		".config/pipewire/pipewire.conf")

	out1, err := cmd1.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(out1))

	// Second command
	cmd2 := exec.Command(
		"sed",
		"-i",
		"s/default.clock.min-quantum[[:space:]]*=[[:space:]]*[0-9]*/default.clock.min-quantum = "+fmt.Sprintf("%d", bufferSizeMinQuantum)+"/",
		".config/pipewire/pipewire.conf")

	out2, err := cmd2.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(out2))

	// Third command
	cmd3 := exec.Command(
		"sed",
		"-i",
		"s/default.clock.max-quantum[[:space:]]*=[[:space:]]*[0-9]*/default.clock.max-quantum = "+fmt.Sprintf("%d", bufferSize)+"/",
		".config/pipewire/pipewire.conf")

	out3, err := cmd3.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(out3))

	// Fourth command
	cmd4 := exec.Command(
		"sed",
		"-i",
		"s/default.clock.quantum-limit[[:space:]]*=[[:space:]]*[0-9]*/default.clock.quantum-limit = "+fmt.Sprintf("%d", bufferSize)+"/",
		".config/pipewire/pipewire.conf")

	out4, err := cmd4.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(out4))

}
