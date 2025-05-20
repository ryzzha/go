package device

type Switchable interface {
    TurnOn() error
    TurnOff() error
}

type Operatable interface {
    Operate(task string) (string, error)
}

type SmartDevice interface {
    Switchable
    Operatable
    GetBrand() string
}

type Device struct {
	Brand string
	IsOn bool
}