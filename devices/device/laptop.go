package device

import (
	"fmt"
	"strings"
)

type Laptop struct {
	Device
	Ram string
}

func (laptop *Laptop) TurnOn() error {
	laptop.IsOn = true
	return nil
}

func (laptop *Laptop) TurnOff() error {
	laptop.IsOn = false
	return nil
}

func (laptop *Laptop) Operate(task string) (string, error) {
	return fmt.Sprintf("op: %s", strings.ToUpper(task)), nil
}

func (laptop *Laptop) GetBrand() string {
	return laptop.Brand
}