package interfaceDesignMonarchy

type Monarchy interface {
	Birth(name, parent string) error
	Death(name string) error
	GetOrderOfSuccession() []string
}
