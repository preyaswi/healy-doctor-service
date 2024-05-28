package usecaseint

import "doctor-service/pkg/models"
type DoctorUseCase interface{
	DoctorSignUp(models.DoctorSignUp)(models.DoctorSignUpResponse,error)
	DoctorLogin(models.DoctorLogin)(models.DoctorSignUpResponse,error)
	DoctorsList()([]models.DoctorsDetails,error)
	IndividualDoctor(doctor_id string)(models.DoctorsDetails,error)
	DoctorProfile(id int)(models.DoctorsDetails,error)
	RateDoctor(patientid int,doctorid string,rate uint32)(uint32,error)
	UpdateDoctorProfile(doctorid int,body models.UpdateDoctor)(models.UpdateDoctor,error)
	DoctorDetailforPayment(doctorid int)(models.DoctorPaymentDetail,error)

}