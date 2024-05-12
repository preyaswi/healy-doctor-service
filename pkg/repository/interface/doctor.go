package interfaces

import (
	"doctor-service/pkg/domain"
	"doctor-service/pkg/models"
)
type DoctorRepository interface{
	CheckDoctorExistsByEmail(email string)(*domain.Doctor,error)
	CheckDoctorExistsByPhone(phone string)(*domain.Doctor,error)
	CheckDoctorexistsByLicenseNumber(licenceNumber string)(*domain.Doctor,error)
	DoctorSignup(models.DoctorSignUp)(models.DoctorDetail,error)
	GetDoctorsDetail()([]models.DoctorsDetails,error)
	ShowIndividualDoctor(doctor_id string)(models.DoctorsDetails,error)
}