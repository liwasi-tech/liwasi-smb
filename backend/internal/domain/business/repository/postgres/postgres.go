package postgres

import (
	"log"

	"github.com/liwasi-tech/liwasi-sbm/backend/internal/domain/business/entity"
	"github.com/liwasi-tech/liwasi-sbm/backend/internal/domain/business/repository"
	"gorm.io/gorm"
)

type postgresBusinessRepository struct {
	gormDB *gorm.DB
}

func NewPostgresBusinessRepository(gormDB *gorm.DB) repository.BusinessRepository {
	return &postgresBusinessRepository{
		gormDB: gormDB,
	}
}

func (r *postgresBusinessRepository) GetByID(ID string) (business *entity.Business, err error) {
	log.Println("getting the business ID: ", ID)
	business = &entity.Business{
		ID: ID,
	}
	err = r.gormDB.First(business).Error
	return
}

func (r *postgresBusinessRepository) Create(business *entity.Business) (err error) {
	err = r.gormDB.Create(business).Error
	return
}

func (r *postgresBusinessRepository) Update(business *entity.Business) (err error) {
	find := &entity.Business{
		ID: business.ID,
	}
	result := r.gormDB.Unscoped().First(find)
	if result.Error != nil {
		err = result.Error
		return
	}
	err = r.gormDB.Save(business).Error
	return
}

func (r *postgresBusinessRepository) DeleteByID(ID string) (err error) {
	business := &entity.Business{
		ID: ID,
	}
	result := r.gormDB.Unscoped().First(business)
	if result.Error != nil {
		err = result.Error
		return
	}
	err = r.gormDB.Unscoped().Delete(business).Error
	return
}

func (r *postgresBusinessRepository) GetAll() (businesses []entity.Business, err error) {
	err = r.gormDB.Unscoped().Find(&businesses).Error
	return
}
