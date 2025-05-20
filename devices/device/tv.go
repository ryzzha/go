package device

import (
	"fmt"
	"strings"
)

type Tv struct {
	Device
	Size string
}

func (tv *Tv) TurnOn() error {
	tv.IsOn = true
	return nil
}

func (tv *Tv) TurnOff() error {
	tv.IsOn = false
	return nil
}

func (tv *Tv) Operate(task string) (string, error) {
	return fmt.Sprintf("op: %s", strings.ToUpper(task)), nil
}

func (tv *Tv) GetBrand() string {
	return tv.Brand
}