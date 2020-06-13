package repository

type Repository string

const (
	Gorm    Repository = "gorm"
	Unknown Repository = "unknown"
)

func (r Repository) String() string {
	return string(r)
}

func Values() []Repository {
	return []Repository{
		Gorm,
	}
}

func ValueOf(value string) Repository {
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
