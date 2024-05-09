package usecaseint

import "doctor-service/pkg/models"
type DoctorUseCase interface{
	DoctorSignUp(models.DoctorSignUp)(models.DoctorSignUpResponse,error)
	DoctorLogin(models.DoctorLogin)(models.DoctorSignUpResponse,error)
	DoctorsList()([]models.DoctorsDetails,error)
}