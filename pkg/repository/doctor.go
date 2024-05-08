package repository

import (
	"doctor-service/pkg/domain"
	"doctor-service/pkg/models"
	interfaces "doctor-service/pkg/repository/interface"
	"errors"

	"gorm.io/gorm"
)

type doctorRepository struct {
	DB *gorm.DB
}

func NewDoctorRepository(DB *gorm.DB) interfaces.DoctorRepository {
	return &doctorRepository{
		DB: DB,
	}
}
func (dr *doctorRepository) CheckDoctorExistsByEmail(email string) (*domain.Doctor, error) {
	var doctor domain.Doctor
	res := dr.DB.Where(&domain.Doctor{Email: email}).First(&doctor)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return &domain.Doctor{}, res.Error
	}
	return &doctor, nil
}
func (dr *doctorRepository) CheckDoctorExistsByPhone(phone string) (*domain.Doctor, error) {
	var doctor domain.Doctor
	res := dr.DB.Where(&domain.Doctor{PhoneNumber: phone}).First(&doctor)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return &domain.Doctor{}, res.Error
	}
	return &doctor, nil
}
func (dr *doctorRepository) CheckDoctorexistsByLicenseNumber(license string) (*domain.Doctor, error) {
	var doctor domain.Doctor
	res := dr.DB.Where(&domain.Doctor{LicenseNumber: license}).First(&doctor)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return &domain.Doctor{}, res.Error
	}
	return &doctor, nil
}
func (dr *doctorRepository) DoctorSignup(doctor models.DoctorSignUp) (models.DoctorDetail, error) {
	var signupDetail models.DoctorDetail
	err := dr.DB.Raw(`
	INSERT INTO doctors(full_name, email, phone_number,password, specialization, years_of_experience, license_number)
	VALUES(?, ?, ?, ?, ?, ?,?)
	RETURNING id, full_name, email, specialization, years_of_experience,license_number
`, doctor.FullName, doctor.Email, doctor.PhoneNumber,doctor.Password, doctor.Specialization, doctor.YearsOfExperience, doctor.LicenseNumber).
	Scan(&signupDetail).Error

	if err != nil {
		return models.DoctorDetail{}, err
	}
	return signupDetail, nil
}
