package usecase

import (
	"doctor-service/pkg/helper"
	"doctor-service/pkg/models"
	interfaces "doctor-service/pkg/repository/interface"
	usecaseint "doctor-service/pkg/usecase/interface"
	"errors"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type doctorUseCase struct {
	doctorRepository interfaces.DoctorRepository
}

func NewDoctorUseCase(repository interfaces.DoctorRepository) usecaseint.DoctorUseCase {
	return &doctorUseCase{
		doctorRepository: repository,
	}
}
func (du *doctorUseCase) DoctorSignUp(doctor models.DoctorSignUp) (models.DoctorSignUpResponse, error) {
	email, err := du.doctorRepository.CheckDoctorExistsByEmail(doctor.Email)
	if err != nil {
		return models.DoctorSignUpResponse{}, errors.New("error with server")
	}
	if email != nil {
		return models.DoctorSignUpResponse{}, errors.New("user with this email is already exists")
	}

	phone, err := du.doctorRepository.CheckDoctorExistsByPhone(doctor.PhoneNumber)
	if err != nil {
		return models.DoctorSignUpResponse{}, errors.New("error with server")
	}
	if phone != nil {
		return models.DoctorSignUpResponse{}, errors.New("user with this phone is already exists")
	}
	if doctor.Password != doctor.ConfirmPassword {
		return models.DoctorSignUpResponse{}, errors.New("confirm password is not matching")
	}

	hashPassword, err := helper.PasswordHash(doctor.Password)
	if err != nil {
		return models.DoctorSignUpResponse{}, errors.New("error in hashing password")
	}
	doctor.Password = hashPassword
	data, err := du.doctorRepository.CheckDoctorexistsByLicenseNumber(doctor.LicenseNumber)
	if err != nil {
		return models.DoctorSignUpResponse{}, errors.New("error with the server")
	}
	if data != nil {
		return models.DoctorSignUpResponse{}, errors.New("doctor with this license number is already exists,please double check")
	}
	DoctorData, err := du.doctorRepository.DoctorSignup(doctor)
	if err != nil {
		return models.DoctorSignUpResponse{}, errors.New("could not add the doctor")
	}
	accessToken, err := helper.GenerateAccessToken(DoctorData)
	if err != nil {
		return models.DoctorSignUpResponse{}, errors.New("couldn't create access token due to error")
	}
	refreshToken, err := helper.GenerateRefreshToken(DoctorData)
	if err != nil {
		return models.DoctorSignUpResponse{}, errors.New("couldn't create refresh token due to error")
	}
	return models.DoctorSignUpResponse{
		DoctorDetail: DoctorData,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
func (du *doctorUseCase) DoctorLogin( doctor models.DoctorLogin) (models.DoctorSignUpResponse, error) {
	data, err := du.doctorRepository.CheckDoctorExistsByEmail(doctor.Email)
	if err != nil {
		return models.DoctorSignUpResponse{}, errors.New("error with server")
	}
	if data == nil {
		return models.DoctorSignUpResponse{}, errors.New("email doesn't exist")
	}
	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(doctor.Password))

	if err != nil {
		return models.DoctorSignUpResponse{}, errors.New("password not matching")
	}
	var doctorDetail models.DoctorDetail
	err = copier.Copy(&doctorDetail, &data)
	if err != nil {
		return models.DoctorSignUpResponse{}, err
	}
	accessToken, err := helper.GenerateAccessToken(doctorDetail)
	if err != nil {
		return models.DoctorSignUpResponse{}, errors.New("couldn't create accesstoken due to internal error")
	}
	refreshToken, err := helper.GenerateRefreshToken(doctorDetail)
	if err != nil {
		return models.DoctorSignUpResponse{}, errors.New("counldn't create refreshtoken due to internal error")
	}
	return models.DoctorSignUpResponse{
		DoctorDetail:      doctorDetail,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
func (du *doctorUseCase)DoctorsList() ([]models.DoctorsDetails, error) {
	doctors,err:=du.doctorRepository.GetDoctorsDetail()
	if err!=nil{
		return nil,err
	}
	return doctors,nil
}