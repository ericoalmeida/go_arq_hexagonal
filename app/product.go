package app

type IProduct interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID     string
	Name   string
	Status string
	Price  float64
}

func (product *Product) IsValid() (bool error) {
	//
}

func (product *Product) Enable() error {
	//
}

func (product *Product) Disable() error {
	//
}

func (product *Product) GetID() string {
	//
}

func (product *Product) GetName() string {
	//
}

func (product *Product) GetStatus() string {
	//
}

func (product *Product) GetPrice() float64 {
	//
}
