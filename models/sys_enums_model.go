package models

// Inisialisasi tipe kustom untuk ENUM JournalMethod
type JournalMethod string

const (
	Manual    JournalMethod = "manual"
	Automatic JournalMethod = "automatic"
)

// Inisialisasi tipe kustom untuk ENUM SubcriptionType
type SubcriptionType string

const (
	Quota    SubcriptionType = "quota"
	Month    SubcriptionType = "month"
	Semester SubcriptionType = "semester"
	Year     SubcriptionType = "year"
)
