package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	var clockChoice string
	var clockRate int
	var dir string
	var err error
	var storeDir string
	var changeDir string
	var bufferSize int

	err = godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	dir = os.Getenv("PIPEWIRE_DIR")
	if dir == "" {
		fmt.Println("Enter your pipewire configuration directory, This will likely be in /usr/share/pipewire or /.config/pipewire: ")
		fmt.Scanln(&dir)

		fmt.Println("Would you like to store your pipewire configuration directory? (y/n): ")
		fmt.Scanln(&storeDir)

		if storeDir == "y" {
			storeDirec(dir, &err)
		}
	}

	if dir != "" {
		fmt.Println("Would you like to change your pipewire configuration directory? (y/n): ")
		fmt.Scanln(&changeDir)

		if changeDir == "y" {
			fmt.Println("Enter your pipewire configuration directory, This will likely be in /usr/share/pipewire or /.config/pipewire: ")
			fmt.Scanln(&dir)
			storeDirec(dir, &err)
		}
	}

	fmt.Println("Do you want to set clock rate? (y/n): ")
	fmt.Scanln(&clockChoice)

	if clockChoice == "y" {
		fmt.Println("Enter the clock rate you want to set the system to: ")

		// Define the clock rate the user wants their system to be set to.
		_, err = fmt.Scanln(&clockRate)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Uncomment and update the clock rate
		cmd := exec.Command(
			"sed",
			"-i",
			`/default.clock.rate[[:space:]]*=[[:space:]]*[0-9]*/s/^[#[:space:]]*/    /; s/default.clock.rate[[:space:]]*=[[:space:]]*[0-9]*/default.clock.rate = `+fmt.Sprintf("%d", clockRate)+`/`,
			dir+"/pipewire.conf")

		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error:", err, "If /etc/pipewire it likely will be .conf.d files, try /.config/pipewire")
		}
		fmt.Println(string(out))

		// Uncomment and update the allowed rates
		cmd1 := exec.Command(
			"sed",
			"-i",
			fmt.Sprintf(`/default.clock.allowed-rates[[:space:]]*=[[:space:]]*\[[[:space:]]*[0-9,[:space:]]*\]/s/^[#[:space:]]*/    /; s/default.clock.allowed-rates[[:space:]]*=[[:space:]]*\[[[:space:]]*[0-9,[:space:]]*\]/default.clock.allowed-rates = \[ %d \]/`, clockRate),
			dir+"/pipewire.conf")

		out1, err := cmd1.CombinedOutput()
		if err != nil {
			fmt.Println("Error:", err, "If /etc/pipewire it likely will be .conf.d files, try /.config/pipewire")
		}
		fmt.Println(string(out1))
	}

	fmt.Println("Enter the buffer size you want to set the system to:")

	// Define the buffer size the user wants their system to be set to.
	_, err = fmt.Scanln(&bufferSize)
	if err != nil {
		fmt.Println("Error:", err, "If /etc/pipewire it likely will be .conf.d files, try /.config/pipewire")
		return
	}
	bufferSizeStr := fmt.Sprintf("%d", bufferSize)

	// Buffer size "default.clock.min-quantum" needs to be half the buffer size.
	bufferSizeMinQuantum := bufferSize / 2

	// Uncomment and update clock quantum
	cmd1 := exec.Command(
		"sed",
		"-i",
		`/default.clock.quantum[[:space:]]*=[[:space:]]*[0-9]*/s/^[#[:space:]]*/    /; s/default.clock.quantum[[:space:]]*=[[:space:]]*[0-9]*/default.clock.quantum = `+fmt.Sprintf("%d", bufferSize)+`/`,
		dir+"/pipewire.conf")

	out1, err := cmd1.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err, "If /etc/pipewire it likely will be .conf.d files, try ~/.config/pipewire")
	}
	fmt.Println(string(out1))

	// Uncomment and update min quantum
	cmd2 := exec.Command(
		"sed",
		"-i",
		`/default.clock.min-quantum[[:space:]]*=[[:space:]]*[0-9]*/s/^[#[:space:]]*/    /; s/default.clock.min-quantum[[:space:]]*=[[:space:]]*[0-9]*/default.clock.min-quantum = `+fmt.Sprintf("%d", bufferSizeMinQuantum)+`/`,
		dir+"/pipewire.conf")

	out2, err := cmd2.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err, "If /etc/pipewire it likely will be .conf.d files, try /.config/pipewire")
	}
	fmt.Println(string(out2))

	// Uncomment and update the max quantum
	cmd3 := exec.Command(
		"sed",
		"-i",
		`/default.clock.max-quantum[[:space:]]*=[[:space:]]*[0-9]*/s/^[#[:space:]]*/    /; s/default.clock.max-quantum[[:space:]]*=[[:space:]]*[0-9]*/default.clock.max-quantum = `+fmt.Sprintf("%d", bufferSize)+`/`,
		dir+"/pipewire.conf")

	out3, err := cmd3.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err, "If /etc/pipewire it likely will be .conf.d files, try /.config/pipewire")
	}
	fmt.Println(string(out3))

	// Uncomment and update the quantum limit
	cmd4 := exec.Command(
		"sed",
		"-i",
		`/default.clock.quantum-limit[[:space:]]*=[[:space:]]*[0-9]*/s/^[#[:space:]]*/    /; s/default.clock.quantum-limit[[:space:]]*=[[:space:]]*[0-9]*/default.clock.quantum-limit = `+fmt.Sprintf("%d", bufferSize)+`/`,
		dir+"/pipewire.conf")

	out4, err := cmd4.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err, "If /etc/pipewire it likely will be .conf.d files, try /.config/pipewire")
	}
	fmt.Println(string(out4))

	// Uncomment and update rt prio
	cmd5 := exec.Command(
		"sed",
		"-i",
		`/rt.prio[[:space:]]*=[[:space:]]*[0-9]*/s/^[#[:space:]]*/            /; s/rt.prio[[:space:]]*=[[:space:]]*[0-9]*/rt.prio = 99/`,
		dir+"/pipewire-pulse.conf",
	)

	out5, err := cmd5.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err, "If /etc/pipewire it likely will be .conf.d files, try /.config/pipewire")
	}
	fmt.Println(string(out5))

	// Uncomment and update the nice level
	cmd6 := exec.Command(
		"sed",
		"-i",
		`/nice.level[[:space:]]*=[[:space:]]*-[0-9]*/s/^[#[:space:]]*/            /; s/nice.level[[:space:]]*=[[:space:]]*-[0-9]*/nice.level = -20/`,
		dir+"/pipewire-pulse.conf",
	)

	out6, err := cmd6.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err, "If /etc/pipewire it likely will be .conf.d files, try /.config/pipewire")
	}
	fmt.Println(string(out6))

	// Uncomment and update node latency
	cmd7 := exec.Command(
		"sed",
		"-i",
		`/node.latency[[:space:]]*=[[:space:]]*[0-9]*/s/^[#[:space:]]*/    /; s/node.latency[[:space:]]*=[[:space:]]*[0-9]*/node.latency = `+fmt.Sprintf("%d", bufferSize)+`/`,
		dir+"/jack.conf",
	)

	out7, err := cmd7.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err, "If /etc/pipewire it likely will be .conf.d files, try /.config/pipewire")
	}
	fmt.Println(string(out7))

	// Uncomment and update node quantum
	cmd8 := exec.Command(
		"sed",
		"-i",
		`/node.quantum[[:space:]]*=[[:space:]]*[0-9]*/s/^[#[:space:]]*/    /; s/node.quantum[[:space:]]*=[[:space:]]*[0-9]*/node.quantum = `+fmt.Sprintf("%d", bufferSize)+`/`,
		dir+"/jack.conf",
	)

	out8, err := cmd8.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err, "If /etc/pipewire it likely will be .conf.d files, try /.config/pipewire")
	}
	fmt.Println(string(out8))

	// Load existing environment variables
	envMap, loadErr := godotenv.Read()
	if loadErr != nil {
		fmt.Println("Error loading .env file")
	}

	// Update the PREV_BUFFER variable
	envMap["PREV_BUFFER"] = bufferSizeStr

	// Write the updated environment variables back to the .env file
	err = godotenv.Write(envMap, ".env")
	if err != nil {
		fmt.Println("Error writing to .env file")
	}
}

func storeDirec(dir string, err *error) {
	// Load existing environment variables
	envMap, loadErr := godotenv.Read()
	if loadErr != nil {
		fmt.Println("Error loading .env file")
	}

	// Update the PIPEWIRE_DIR variable
	envMap["PIPEWIRE_DIR"] = dir

	// Write the updated environment variables back to the .env file
	*err = godotenv.Write(envMap, ".env")
	if *err != nil {
		fmt.Println("Error writing to .env file")
	}
}
