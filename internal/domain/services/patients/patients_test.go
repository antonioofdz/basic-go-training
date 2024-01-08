package patients

import (
	"basic-go-training/internal/config"
	"basic-go-training/internal/database"
	dto "basic-go-training/internal/domain/dtos/patients"
	"basic-go-training/internal/domain/entities/patients"
	"testing"

	"github.com/juju/errors"
	"github.com/stretchr/testify/require"
)

var service *Service

func initTestbed(t *testing.T) {
	config.Settings.Database.Name = "patientdb"
	config.Settings.Database.Host = "localhost"
	config.Settings.Database.Password = "password"
	config.Settings.Database.User = "user"
	config.Settings.Database.Port = "5432"

	require.NoError(t, database.Connect())
	require.NoError(t, database.Repo.AllPatients().Where("id IS NOT NULL").Delete(&patients.Patient{}).Error)

	service = &Service{}

	patient := &patients.Patient{
		ID:    1,
		Name:  "John Doe",
		Email: "foo@bar.com",
	}
	require.NoError(t, database.Repo.AllPatients().Create(patient).Error)
}

func TestGetNotFound(t *testing.T) {
	initTestbed(t)

	req := &dto.GetRequest{
		ID: 111,
	}
	_, err := service.Get(req)
	require.Error(t, err)
	require.ErrorIs(t, err, errors.NotFound)
}

func TestGet(t *testing.T) {
	initTestbed(t)

	req := &dto.GetRequest{
		ID: 1,
	}
	patient, err := service.Get(req)
	require.NoError(t, err)

	require.Equal(t, patient.ID, 1)
	require.Equal(t, patient.Name, "John Doe")
	require.Equal(t, patient.Email, "foo@bar.com")
}
