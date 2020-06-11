package repository

type Database string

const (
	Gorm    Database = "gorm"
	Unknown Database = "unknown"
)

func Values() []Database {
	return []Database{
		Gorm,
	}
}

func ValueOf(value string) Database {
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
