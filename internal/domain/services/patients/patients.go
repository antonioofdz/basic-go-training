package patients

import (
	"basic-go-training/internal/database"
	dto "basic-go-training/internal/domain/dtos/patients"
	entities "basic-go-training/internal/domain/entities/patients"

	"github.com/juju/errors"
	"gorm.io/gorm"
)

type PatientsService interface {
	Search(req *dto.SearchRequest) (*dto.SearchReply, error)
	Get(req *dto.GetRequest) (*dto.Patient, error)
	Create(req *dto.CreateRequest) (*dto.Patient, error)
	Update(req *dto.UpdateRequest) (*dto.Patient, error)
	Delete(req *dto.DeleteRequest) (*dto.Empty, error)
}

func serializePatient(patient *entities.Patient) *dto.Patient {
	return &dto.Patient{
		ID:    patient.ID,
		Name:  patient.Name,
		Email: patient.Email,
	}
}

type Service struct {
}

func NewService() PatientsService {
	return new(Service)
}

func (s *Service) Search(req *dto.SearchRequest) (*dto.SearchReply, error) {
	q := database.Repo.AllPatients()
	if req.Name != "" {
		q = q.Where("name = ?", req.Name)
	}
	if req.Email != "" {
		q = q.Where("email = ?", req.Email)
	}

	patients := []*entities.Patient{}
	if err := q.Find(&patients).Error; err != nil {
		return nil, errors.Trace(err)
	}

	reply := &dto.SearchReply{}
	for _, patient := range patients {
		reply.Patients = append(reply.Patients, serializePatient(patient))
	}

	return reply, nil
}

func (s *Service) Get(req *dto.GetRequest) (*dto.Patient, error) {
	patient := &entities.Patient{}
	if err := database.Repo.Patients(req.ID).First(patient).Error; err != nil {
		// ¿Falta algo aquí?
		return nil, errors.Trace(err)
	}
	return serializePatient(patient), nil
}

func (s *Service) Create(req *dto.CreateRequest) (*dto.Patient, error) {
	if req.Email == "" {
		return nil, errors.Trace(errors.BadRequestf("email required"))
	}
	if req.Name == "" {
		return nil, errors.Trace(errors.BadRequestf("name required"))
	}

	// TODO: ¿Falta alguna/s comprobación aquí?

	patient := &entities.Patient{
		Name:  req.Name,
		Email: req.Email,
	}
	if err := database.Repo.AllPatients().Create(patient).Error; err != nil {
		return nil, errors.Trace(err)
	}

	return serializePatient(patient), nil
}

func (s *Service) Update(req *dto.UpdateRequest) (*dto.Patient, error) {
	if req.Email == "" {
		return nil, errors.Trace(errors.BadRequestf("email required"))
	}
	if req.Name == "" {
		return nil, errors.Trace(errors.BadRequestf("name required"))
	}

	// TODO: ¿Falta alguna/s comprobación aquí?

	patient := &entities.Patient{}
	if err := database.Repo.Patients(req.ID).First(patient).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Trace(errors.NotFoundf("patient: %v", req.ID))
		}

		return nil, errors.Trace(err)
	}

	patient.Name = req.Name
	patient.Email = req.Email

	if err := database.Repo.Patients(req.ID).Save(patient).Error; err != nil {
		return nil, errors.Trace(err)
	}

	return serializePatient(patient), nil
}

func (s *Service) Delete(req *dto.DeleteRequest) (*dto.Empty, error) {
	return &dto.Empty{}, nil
}

/*
if err == gorm.ErrRecordNotFound {
	return nil, errors.Trace(errors.NotFoundf("patient: %v", req.ID))
}
*/

/*
var count int64
if err := database.Repo.AllPatients().Debug().Where("email = ?", req.Email).Count(&count).Error; err != nil {
	return nil, errors.Trace(err)
} else if count > 0 {
	return nil, errors.Trace(errors.BadRequestf("email already exists"))
}
*/
