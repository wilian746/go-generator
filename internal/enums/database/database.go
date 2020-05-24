package database

type Database string

const (
	Gorm = "gorm"
	Unknown = "unknown"
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