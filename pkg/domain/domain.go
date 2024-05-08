package domain
type Doctor struct{
	ID                uint   `json:"id" gorm:"uniquekey; not null"`
    FullName          string `json:"full_name" gorm:"not null; validate:required"`
    Email             string `json:"email" gorm:"unique;validate:required"`
    PhoneNumber       string `json:"phone_number" gorm:"validate:required"`
    Password          string `json:"password" gorm:"validate:required"`
    Specialization    string `json:"specialization"`
    YearsOfExperience int32  `json:"years_of_experience"`
    LicenseNumber     string `json:"license_number" gorm:"unique;validate:required"`
}