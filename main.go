package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	var clockRate int
	var dir string
	var err error
	var bufferSize int
	var funcChoice string
	var funcAppend string

	err = godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	funcMap := make(map[string]func())

	funcMap["bufferSet"] = func() { bufferSet(bufferSize, dir, &err) }
	funcMap["storeDirec"] = func() { storeDirec(dir, &err) }
	funcMap["clockSet"] = func() { clockSet(clockRate, dir, &err) }

	fmt.Println("[1] Set buffer size")
	fmt.Println("[2] Store directory")
	fmt.Println("[3] Change clock rate")
	fmt.Println("[4] Exit")
	fmt.Println("Enter your choice: ")
	fmt.Scanln(&funcChoice)

	switch funcChoice {
	case "1":
		funcAppend = "bufferSet"
	case "2":
		funcAppend = "storeDirec"
	case "3":
		funcAppend = "clockSet"
	case "4":
		os.Exit(0)
	default:
		fmt.Println("Invalid choice")
	}

	funcMap[funcAppend]()
}

func bufferSet(bufferSize int, dir string, err *error) {
	fmt.Println("Enter the buffer size you want to set the system to: ")

	// Define the buffer size the user wants their system to be set to.
	_, errScan := fmt.Scanln(&bufferSize)
	if errScan != nil {
		fmt.Println("Error:", err, "If /etc/pipewire it likely will be .conf.d files, try /.config/pipewire")
		return
	}
	bufferSizeStr := fmt.Sprintf("%d", bufferSize)

	dir = os.Getenv("PIPEWIRE_DIR")
	
	// Buffer size "default.clock.min-quantum" needs to be half the buffer size.
	bufferSizeMinQuantum := bufferSize / 2

	// Uncomment and update clock quantum
	cmd1 := exec.Command(
		"sed",
		"-i",
		`/default.clock.quantum[[:space:]]*=[[:space:]]*[0-9]*/s/^[#[:space:]]*/    /; s/default.clock.quantum[[:space:]]*=[[:space:]]*[0-9]*/default.clock.quantum = `+fmt.Sprintf("%d", bufferSize)+`/`,
		dir+"/pipewire.conf")

	out1, errCmd := cmd1.CombinedOutput()
	if errCmd != nil {
		fmt.Println("Error:", err, "If /etc/pipewire it likely will be .conf.d files, try ~/.config/pipewire")
	}
	fmt.Println(string(out1))

	// Uncomment and update min quantum
	cmd2 := exec.Command(
		"sed",
		"-i",
		`/default.clock.min-quantum[[:space:]]*=[[:space:]]*[0-9]*/s/^[#[:space:]]*/    /; s/default.clock.min-quantum[[:space:]]*=[[:space:]]*[0-9]*/default.clock.min-quantum = `+fmt.Sprintf("%d", bufferSizeMinQuantum)+`/`,
		dir+"/pipewire.conf")

	out2, errCmd := cmd2.CombinedOutput()
	if errCmd != nil {
		fmt.Println("Error:", err, "If /etc/pipewire it likely will be .conf.d files, try /.config/pipewire")
	}
	fmt.Println(string(out2))

	// Uncomment and update the max quantum
	cmd3 := exec.Command(
		"sed",
		"-i",
		`/default.clock.max-quantum[[:space:]]*=[[:space:]]*[0-9]*/s/^[#[:space:]]*/    /; s/default.clock.max-quantum[[:space:]]*=[[:space:]]*[0-9]*/default.clock.max-quantum = `+fmt.Sprintf("%d", bufferSize)+`/`,
		dir+"/pipewire.conf")

	out3, errCmd := cmd3.CombinedOutput()
	if errCmd != nil {
		fmt.Println("Error:", err, "If /etc/pipewire it likely will be .conf.d files, try /.config/pipewire")
	}
	fmt.Println(string(out3))

	// Uncomment and update the quantum limit
	cmd4 := exec.Command(
		"sed",
		"-i",
		`/default.clock.quantum-limit[[:space:]]*=[[:space:]]*[0-9]*/s/^[#[:space:]]*/    /; s/default.clock.quantum-limit[[:space:]]*=[[:space:]]*[0-9]*/default.clock.quantum-limit = `+fmt.Sprintf("%d", bufferSize)+`/`,
		dir+"/pipewire.conf")

	out4, errCmd := cmd4.CombinedOutput()
	if errCmd != nil {
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

	out5, errCmd := cmd5.CombinedOutput()
	if errCmd != nil {
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

	out6, errCmd := cmd6.CombinedOutput()
	if errCmd != nil {
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

	out7, errCmd := cmd7.CombinedOutput()
	if errCmd != nil {
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

	out8, errCmd := cmd8.CombinedOutput()
	if errCmd != nil {
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
	errEnv := godotenv.Write(envMap, ".env")
	if errEnv != nil {
		fmt.Println("Error writing to .env file")
	}
}

func storeDirec(dir string, err *error) {
	fmt.Println("Enter your pipewire configuration directory, This will likely be in /usr/share/pipewire or /.config/pipewire: ")
	fmt.Scanln(&dir)

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

func clockSet(clockRate int, dir string, err *error) {
	fmt.Println("Enter the clock rate you want to set the system to: ")

	// Define the clock rate the user wants their system to be set to.
	_, errScan := fmt.Scanln(&clockRate)
	if errScan != nil {
		fmt.Println("Error:", err)
		return
	}

	// Uncomment and update the clock rate
	cmd := exec.Command(
		"sed",
		"-i",
		`/default.clock.rate[[:space:]]*=[[:space:]]*[0-9]*/s/^[#[:space:]]*/    /; s/default.clock.rate[[:space:]]*=[[:space:]]*[0-9]*/default.clock.rate = `+fmt.Sprintf("%d", clockRate)+`/`,
		dir+"/pipewire.conf")

	out, errCmd := cmd.CombinedOutput()
	if errCmd != nil {
		fmt.Println("Error:", err, "If /etc/pipewire it likely will be .conf.d files, try /.config/pipewire")
	}
	fmt.Println(string(out))

	// Uncomment and update the allowed rates
	cmd1 := exec.Command(
		"sed",
		"-i",
		fmt.Sprintf(`/default.clock.allowed-rates[[:space:]]*=[[:space:]]*\[[[:space:]]*[0-9,[:space:]]*\]/s/^[#[:space:]]*/    /; s/default.clock.allowed-rates[[:space:]]*=[[:space:]]*\[[[:space:]]*[0-9,[:space:]]*\]/default.clock.allowed-rates = \[ %d \]/`, clockRate),
		dir+"/pipewire.conf")

	out1, errCmd := cmd1.CombinedOutput()
	if errCmd != nil {
		fmt.Println("Error:", err, "If /etc/pipewire it likely will be .conf.d files, try /.config/pipewire")
	}
	fmt.Println(string(out1))
}
