package values_objects

type ValueObject interface {
	GetValue() interface{}
	Equals(interface{}) bool
}
