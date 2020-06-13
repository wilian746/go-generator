package version

type ValueFlag struct {
	Value interface{}
}

func (v *ValueFlag) String() string {
	return v.Value.(string)
}
func (v *ValueFlag) Set(value string) error {
	v.Value = value
	return nil
}
func (v *ValueFlag) Type() string {
	return "string"
}
