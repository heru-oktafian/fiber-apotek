package models

// Initialize custom type for ENUM JournalMethod
type JournalMethod string

const (
	Manual    JournalMethod = "manual"
	Automatic JournalMethod = "automatic"
)

// Initialize custom type for ENUM SubcriptionType
type SubcriptionType string

const (
	Quota    SubcriptionType = "quota"
	Month    SubcriptionType = "month"
	Semester SubcriptionType = "semester"
	Year     SubcriptionType = "year"
)
