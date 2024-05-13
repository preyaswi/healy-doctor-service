package domain
type Doctor struct {
    ID                uint   `json:"id" gorm:"unique;not null"`
    FullName          string `json:"full_name" gorm:"not null"`
    Email             string `json:"email" gorm:"unique"`
    PhoneNumber       string `json:"phone_number"`
    Password          string `json:"password"`
    Specialization    string `json:"specialization"`
    YearsOfExperience int32  `json:"years_of_experience"`
    LicenseNumber     string `json:"license_number" gorm:"unique"`
}
type Review struct {
    ID        uint `json:"id" gorm:"primaryKey"`
    DoctorID  uint `json:"doctor_id" gorm:"not null"`
    PatientID uint `json:"patient_id" gorm:"not null"`
    Rating    int  `json:"rating" gorm:"not null"`
}