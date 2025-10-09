package repository

import (
	"reflect"
	"testing"

	"github.com/SkyFlareInfra/SkyFlare/infra"
	"github.com/SkyFlareInfra/SkyFlare/models"
	"gorm.io/gorm"
)

func TestNewDomainRepository(t *testing.T) {
	type args struct {
		db infra.DatabaseManager
	}
	tests := []struct {
		name string
		args args
		want DomainRepositoryInterface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDomainRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDomainRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDomainRepository_CreateInTransaction(t *testing.T) {
	type fields struct {
		db infra.DatabaseManager
	}
	type args struct {
		tx   *gorm.DB
		site *models.Domain
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := &DomainRepository{
				db: tt.fields.db,
			}
			if err := ss.CreateInTransaction(tt.args.tx, tt.args.site); (err != nil) != tt.wantErr {
				t.Errorf("DomainRepository.CreateInTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDomainRepository_Create(t *testing.T) {
	type fields struct {
		db infra.DatabaseManager
	}
	type args struct {
		site *models.Domain
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := &DomainRepository{
				db: tt.fields.db,
			}
			if err := sr.Create(tt.args.site); (err != nil) != tt.wantErr {
				t.Errorf("DomainRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDomainRepository_GetAll(t *testing.T) {
	type fields struct {
		db infra.DatabaseManager
	}
	type args struct {
		filter models.Domain
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.Domain
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := &DomainRepository{
				db: tt.fields.db,
			}
			got, err := sr.GetAll(tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("DomainRepository.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DomainRepository.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDomainRepository_GetDomainByName(t *testing.T) {
	type fields struct {
		db infra.DatabaseManager
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Domain
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := &DomainRepository{
				db: tt.fields.db,
			}
			got, err := sr.GetDomainByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("DomainRepository.GetDomainByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DomainRepository.GetDomainByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDomainRepository_GetByUID(t *testing.T) {
	type fields struct {
		db infra.DatabaseManager
	}
	type args struct {
		uuid string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Domain
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := &DomainRepository{
				db: tt.fields.db,
			}
			got, err := sr.GetByUID(tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("DomainRepository.GetByUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DomainRepository.GetByUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDomainRepository_GetByID(t *testing.T) {
	type fields struct {
		db infra.DatabaseManager
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Domain
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := &DomainRepository{
				db: tt.fields.db,
			}
			got, err := sr.GetByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DomainRepository.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DomainRepository.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDomainRepository_Get(t *testing.T) {
	type fields struct {
		db infra.DatabaseManager
	}
	type args struct {
		filter *models.Domain
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Domain
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := &DomainRepository{
				db: tt.fields.db,
			}
			got, err := sr.Get(tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("DomainRepository.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DomainRepository.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDomainRepository_Update(t *testing.T) {
	type fields struct {
		db infra.DatabaseManager
	}
	type args struct {
		filter models.Domain
		data   *models.Domain
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := &DomainRepository{
				db: tt.fields.db,
			}
			if err := sr.Update(tt.args.filter, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("DomainRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDomainRepository_Delete(t *testing.T) {
	type fields struct {
		db infra.DatabaseManager
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := &DomainRepository{
				db: tt.fields.db,
			}
			if err := sr.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DomainRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDomainRepository_DeleteByUID(t *testing.T) {
	type fields struct {
		db infra.DatabaseManager
	}
	type args struct {
		uid string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := &DomainRepository{
				db: tt.fields.db,
			}
			if err := sr.DeleteByUID(tt.args.uid); (err != nil) != tt.wantErr {
				t.Errorf("DomainRepository.DeleteByUID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDomainRepository_GetDomainNameByDomainID(t *testing.T) {
	type fields struct {
		db infra.DatabaseManager
	}
	type args struct {
		domainID uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dr := &DomainRepository{
				db: tt.fields.db,
			}
			got, err := dr.GetDomainNameByDomainID(tt.args.domainID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DomainRepository.GetDomainNameByDomainID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DomainRepository.GetDomainNameByDomainID() = %v, want %v", got, tt.want)
			}
		})
	}
}
