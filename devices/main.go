package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"

	"devices/device"
)

func main() {
	tv := device.Tv{
		Device: device.Device{Brand: "Samsung", IsOn: false},
		Size:   "55",
	}

	laptop := device.Laptop{
		Device: device.Device{Brand: "Macbook", IsOn: false},
		Ram:    "15",
	}

	for {
		fmt.Println("\n👉 Enter device name: 'tv', 'laptop' or 'exit'")
		deviceChoice := strings.TrimSpace(readUserInput())

		if deviceChoice == "exit" {
			fmt.Println("Exiting the program. Goodbye!")
			break
		}

		var selected device.SmartDevice
		var rawDevice interface{}

		switch deviceChoice {
		case "tv":
			selected = &tv
			rawDevice = tv
		case "laptop":
			selected = &laptop
			rawDevice = laptop
		default:
			fmt.Println("Unknown device")
			return
		}

		fmt.Println("Choose an action: on / off / operate / info")
		action := strings.TrimSpace(readUserInput())

		switch action {
		case "on":
			turnOnDevice(selected)
		case "off":
			turnOffDevice(selected)
		case "operate":
			fmt.Println("Enter a task to execute:")
			task := strings.TrimSpace(readUserInput())
			operateDevice(selected, task)
		case "info":
			displayInfo(rawDevice)
		default:
			fmt.Println("Unknown action")
		}
	}
}

func turnOnDevice(device device.SmartDevice) {
	err := device.TurnOn()
	if err != nil {
		fmt.Printf("Error turning on device: %v\n", err)
		return
	}
	fmt.Printf("Device %s is now ON\n", device.GetBrand())
}

func turnOffDevice(device device.SmartDevice) {
	err := device.TurnOff()
	if err != nil {
		fmt.Printf("Error turning off device: %v\n", err)
		return
	}
	fmt.Printf("Device %s is now OFF\n", device.GetBrand())
}

func operateDevice(device device.SmartDevice, task string) {
	operatedTask, err := device.Operate(task)
	if err != nil {
		fmt.Printf("Error operating task: %v\n", task)
		return
	}

	fmt.Printf("Task executed: %s\n", operatedTask)
}

func displayInfo(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Println("This is a string:", v)
	case int:
		fmt.Println("This is an int:", v)
	case device.Tv:
		status := "off"
		if v.IsOn {
			status = "on"
		}
		fmt.Println("This is a TV with brand:", v.Brand, ", size:", v.Size, ", status:", status)
	case device.Laptop:
		status := "off"
		if v.IsOn {
			status = "on"
		}
		fmt.Println("This is a Laptop with brand:", v.Brand, ", ram:", v.Ram, ", status:", status)
	default:
		fmt.Println("Unknown type")
	}
}

func readUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter input: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Input read error:", err)
		return ""
	}

	return input
}
