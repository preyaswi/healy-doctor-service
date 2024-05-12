package domain
type Doctor struct{
	ID                uint   `json:"id" gorm:"uniquekey; not null"`
    FullName          string `json:"full_name" gorm:"not null;"`
    Email             string `json:"email" gorm:"unique;`
    PhoneNumber       string `json:"phone_number" gorm:"`
    Password          string `json:"password" gorm:"`
    Specialization    string `json:"specialization"`
    YearsOfExperience int32  `json:"years_of_experience"`
    LicenseNumber     string `json:"license_number" gorm:"unique;`
}