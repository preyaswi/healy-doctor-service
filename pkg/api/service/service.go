package service

import (
	"context"
	"doctor-service/pkg/models"
	"doctor-service/pkg/pb"
	usecaseint "doctor-service/pkg/usecase/interface"
)
type DoctorServer struct{
	doctorUseCase usecaseint.DoctorUseCase
	pb.UnimplementedDoctorServer

}
func NewDoctorServer(usecase usecaseint.DoctorUseCase) pb.DoctorServer {
	return &DoctorServer{
		doctorUseCase: usecase,
	}
}
func (d *DoctorServer)DoctorSignUp(ctx context.Context,DoctorSignUpRequest *pb.DoctorSignUpRequest) (*pb.DoctorSignUpResponse, error){
	doctorsignup:=models.DoctorSignUp{
		FullName: DoctorSignUpRequest.FullName,
		Email: DoctorSignUpRequest.Email,
		PhoneNumber: DoctorSignUpRequest.PhoneNumber,
		Password: DoctorSignUpRequest.Password,
		ConfirmPassword: DoctorSignUpRequest.ConfirmPassword,
		Specialization: DoctorSignUpRequest.Specialization,
		YearsOfExperience: DoctorSignUpRequest.YearsOfExperience,
		LicenseNumber: DoctorSignUpRequest.LicenseNumber,
	}
	data,err:=d.doctorUseCase.DoctorSignUp(doctorsignup)
	if err!=nil{
		return &pb.DoctorSignUpResponse{},err
	}
	doctordetail:=&pb.DoctorDetail{
		Id: uint64(data.DoctorDetail.Id),
		FullName: data.DoctorDetail.FullName,
		Email: data.DoctorDetail.Email,
		PhoneNumber: data.DoctorDetail.PhoneNumber,
		Specialization: data.DoctorDetail.Specialization,
		YearsOfExperience: data.DoctorDetail.YearsOfExperience,
		LicenseNumber: data.DoctorDetail.LicenseNumber,
	}
	return &pb.DoctorSignUpResponse{
		DoctorDetail: doctordetail,
		AccessToken: data.AccessToken,
		RefreshToken: data.RefreshToken,
	},nil
}
func (d *DoctorServer)DoctorLogin(ctx context.Context,DoctorLoginRequest *pb.DoctorLoginRequest) (*pb.DoctorLoginResponse, error){
	doctorLogin:=models.DoctorLogin{
		Email: DoctorLoginRequest.Email,
		Password: DoctorLoginRequest.Password,
	}
	data,err:=d.doctorUseCase.DoctorLogin(doctorLogin)
	if err!=nil{
		return &pb.DoctorLoginResponse{},err
	}
	doctordetail:=&pb.DoctorDetail{
		Id: uint64(data.DoctorDetail.Id),
		FullName: data.DoctorDetail.FullName,
		Email: data.DoctorDetail.Email,
		PhoneNumber: data.DoctorDetail.PhoneNumber,
		Specialization: data.DoctorDetail.Specialization,
		YearsOfExperience: data.DoctorDetail.YearsOfExperience,
		LicenseNumber: data.DoctorDetail.LicenseNumber,
	}
	return &pb.DoctorLoginResponse{
		DoctorDetail: doctordetail,
		AccessToken: data.AccessToken,
		RefreshToken: data.RefreshToken,
	},nil
}
func (d *DoctorServer)DoctorsDetail(ctx context.Context,req *pb.Doreq) (*pb.DoctorsDetailre, error)  {
	doctors,err:=d.doctorUseCase.DoctorsList()
	if err!=nil{
		return &pb.DoctorsDetailre{},err
	}
	doctorDetails := make([]*pb.DoctorsDetailr, len(doctors))
	for i, doctor := range doctors {
		// Map each DoctorDetail to the corresponding protobuf message
		doctorDetails[i] = &pb.DoctorsDetailr{
			Id:                uint64(doctor.DoctorDetail.Id),
			FullName:          doctor.DoctorDetail.FullName,
			Email:             doctor.DoctorDetail.Email,
			PhoneNumber:       doctor.DoctorDetail.PhoneNumber,
			Specialization:    doctor.DoctorDetail.Specialization,
			YearsOfExperience: doctor.DoctorDetail.YearsOfExperience,
			LicenseNumber:     doctor.DoctorDetail.LicenseNumber,
			Rating: doctor.Rating,
		}
	}

	// Create and return the DoctorsListResponse
	return &pb.DoctorsDetailre{
		DoctorsDetailr: doctorDetails,
	}, nil

}
func ( d *DoctorServer)IndividualDoctor(ctx context.Context,req *pb.Doid) (*pb.DoctorsDetailr, error)  {
	doctor,err:=d.doctorUseCase.IndividualDoctor(req.DoctorId)
	if err!=nil{
		return &pb.DoctorsDetailr{},err
	}
	return &pb.DoctorsDetailr{
		Id: uint64(doctor.DoctorDetail.Id),
		FullName: doctor.DoctorDetail.FullName,
		Email: doctor.DoctorDetail.Email,
		PhoneNumber: doctor.DoctorDetail.PhoneNumber,
		Specialization: doctor.DoctorDetail.Specialization,
		YearsOfExperience: doctor.DoctorDetail.YearsOfExperience,
		LicenseNumber: doctor.DoctorDetail.LicenseNumber,
		Rating: doctor.Rating,
	},nil

}
func (d *DoctorServer) DoctorProfile(ctx context.Context,req *pb.DoId) (*pb.DoctorsDetailr, error) {
	doctor,err:=d.doctorUseCase.DoctorProfile(int(req.Id))
	if err!=nil{
		return &pb.DoctorsDetailr{},err
	}
	return &pb.DoctorsDetailr{
		Id: uint64(doctor.DoctorDetail.Id),
		FullName: doctor.DoctorDetail.FullName,
		Email: doctor.DoctorDetail.Email,
		PhoneNumber: doctor.DoctorDetail.PhoneNumber,
		Specialization: doctor.DoctorDetail.Specialization,
		YearsOfExperience: doctor.DoctorDetail.YearsOfExperience,
		LicenseNumber: doctor.DoctorDetail.LicenseNumber,
		Rating: doctor.Rating,
	},err
}
func (d *DoctorServer)RateDoctor(ctx context.Context,req *pb.RateDoctorReq) (*pb.Rate, error)  {
	patientid:=req.Patientid
	doctorId:=req.DoctorId
	rate:=req.Rate.Rate
	rated,err:=d.doctorUseCase.RateDoctor(int(patientid),doctorId,rate)
	if err!=nil{
		return &pb.Rate{},err
	}
	return &pb.Rate{
		Rate: rated,
	},nil

}
func (d *DoctorServer) UpdateDoctorProifle(ctx context.Context,req *pb.UpdateReq) (*pb.UpdateDoctor, error) {
	doctorid:=req.Id
	body:=models.UpdateDoctor{
		FullName: req.Body.FullName,
		Email: req.Body.Email,
		PhoneNumber: req.Body.PhoneNumber,
		Specialization: req.Body.Specialization,
		YearsOfExperience: req.Body.YearsOfExperience,
	}
	res,err:=d.doctorUseCase.UpdateDoctorProfile(int(doctorid),body)
	if err!=nil{
		return &pb.UpdateDoctor{},err
	}
	return &pb.UpdateDoctor{
		FullName: res.FullName,
		Email: res.Email,
		PhoneNumber: res.PhoneNumber,
		Specialization: res.Specialization,
		YearsOfExperience: res.YearsOfExperience,
	},nil
}
func (d *DoctorServer)DoctorDetailforPayment(ctx context.Context,req  *pb.DoId) (*pb.DoctorPaymentDetail, error)  {
	res,err:=d.doctorUseCase.DoctorDetailforPayment(int(req.Id))
	if err!=nil{
		return &pb.DoctorPaymentDetail{},nil
	}
	return &pb.DoctorPaymentDetail{
		DoctorId: int32(res.Doctor_id),
		DoctorName: res.DoctorName,
		Fees: res.Fees,
	},nil
}