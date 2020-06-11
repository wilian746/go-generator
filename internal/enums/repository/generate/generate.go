package generate

type Command string

const (
	App     Command = "app"
	Unknown Command = "unknown"
)

func Values() []Command {
	return []Command{
		App,
	}
}

func ValueOf(value string) Command {
	for _, db := range Values() {
		if string(db) == value {
			return db
		}
	}
	return Unknown
}

func Valid(value string) bool {
	return ValueOf(value) != Unknown
}
