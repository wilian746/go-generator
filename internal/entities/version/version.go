package version

type Version struct {
	Value string
}

func (v *Version) String() string {
	return v.Value
}
func (v *Version) Set(value string) error {
	v.Value = value
	return nil
}
func (v *Version) Type() string {
	return ""
}
