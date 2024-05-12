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
`, doctor.FullName, doctor.Email, doctor.PhoneNumber, doctor.Password, doctor.Specialization, doctor.YearsOfExperience, doctor.LicenseNumber).
		Scan(&signupDetail).Error

	if err != nil {
		return models.DoctorDetail{}, err
	}
	return signupDetail, nil
}
func (dr *doctorRepository) GetDoctorsDetail() ([]models.DoctorsDetails, error) {
	// Query to select doctors along with their average ratings from the review table
	rows, err := dr.DB.Raw(`
		SELECT 
			d.id, d.full_name, d.email, d.phone_number, d.specialization, d.years_of_experience, d.license_number, COALESCE(AVG(r.rating), 0) AS average_rating
		FROM 
			doctors AS d
		LEFT JOIN 
			review AS r ON d.id = r.doctor_id
		GROUP BY 
			d.id
		ORDER BY 
			average_rating DESC
	`).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var doctorsDetails []models.DoctorsDetails
	for rows.Next() {
		var doctorDetail models.DoctorDetail
		var rating float64 // Use float64 for average rating

		// Scan the row into variables
		if err := rows.Scan(&doctorDetail.Id, &doctorDetail.FullName, &doctorDetail.Email, &doctorDetail.PhoneNumber, &doctorDetail.Specialization, &doctorDetail.YearsOfExperience, &doctorDetail.LicenseNumber, &rating); err != nil {
			return nil, err
		}

		// Append doctor details to the result
		doctorsDetails = append(doctorsDetails, models.DoctorsDetails{
			DoctorDetail: doctorDetail,
			Rating:       int32(rating), // Convert float64 to int32
		})
	}

	return doctorsDetails, nil
}
func (dr *doctorRepository)ShowIndividualDoctor(doctor_id string)(models.DoctorsDetails,error)  {
	rows, err := dr.DB.Raw(`
		SELECT 
			d.id, d.full_name, d.email, d.phone_number, d.specialization, d.years_of_experience, d.license_number, COALESCE(AVG(r.rating), 0) AS average_rating
		FROM 
			doctors AS d
		LEFT JOIN 
			review AS r ON d.id = r.doctor_id
		WHERE 
			d.id = ?
		GROUP BY 
			d.id
		ORDER BY 
			average_rating DESC
	`, doctor_id).Rows()
	if err != nil {
		return models.DoctorsDetails{}, err
	}
	defer rows.Close()

	var doctorDetails models.DoctorsDetails
	if rows.Next() {
		var doctorDetail models.DoctorDetail
		var rating float64 // Use float64 for average rating

		// Scan the row into variables
		if err := rows.Scan(&doctorDetail.Id, &doctorDetail.FullName, &doctorDetail.Email, &doctorDetail.PhoneNumber, &doctorDetail.Specialization, &doctorDetail.YearsOfExperience, &doctorDetail.LicenseNumber, &rating); err != nil {
			return models.DoctorsDetails{}, err
		}

		// Populate the result with doctor details
		doctorDetails = models.DoctorsDetails{
			DoctorDetail: doctorDetail,
			Rating:       int32(rating), // Convert float64 to int32
		}
	}

	return doctorDetails, nil
}