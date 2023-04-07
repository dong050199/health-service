package repository

import (
	"health_service/pkg/infra"
	"health_service/service/model/entity"
	"health_service/service/model/request"

	"gorm.io/gorm"
)

type IproductRepo interface {
}

type productRepo struct {
	ormDB *gorm.DB
}

func NewProductRepo() IproductRepo {
	return &productRepo{
		ormDB: infra.GetDB(),
	}
}

func (r *productRepo) Insert(body entity.Products, tx *gorm.DB) error {
	return tx.Create(&body).Error
}

func (r *productRepo) Update(body entity.Products, tx *gorm.DB) error {
	return tx.Save(&body).Error
}

func (r *productRepo) Delete(body entity.Products, tx *gorm.DB) error {
	return tx.Save(&entity.Products{
		ID:        body.ID,
		DeletedAt: &timeNow,
		DeletedBy: body.DeletedBy,
	}).Error
}

func (m *productRepo) GetByID(
	id uint,
) (meals entity.Products, err error) {
	dbQuery := m.ormDB.Find(&entity.Products{ID: id})
	if dbQuery.Error != nil {
		err = dbQuery.Error
		if dbQuery.Error == gorm.ErrEmptySlice {
			err = nil
			return
		}
		return
	}
	return
}

func (m *productRepo) GetList(
	req request.PagingRequest,
) (meals []entity.Products, err error) {
	dbQuery := m.ormDB.Offset(req.Offset).Limit(req.Limit).Find(&entity.Products{ID: id})
	if dbQuery.Error != nil {
		err = dbQuery.Error
		if dbQuery.Error == gorm.ErrEmptySlice {
			err = nil
			return
		}
		return
	}
	return
}

func (m *productRepo) GetPaging(
	req request.PagingRequest,
) (totalPages int, err error) {
	var NumberRecode int64
	dbQuery := m.ormDB.Model(&entity.Products{}).Count(&NumberRecode)
	if dbQuery.Error != nil {
		err = dbQuery.Error
		if dbQuery.Error == gorm.ErrEmptySlice {
			err = nil
			return
		}
		return
	}

	totalPages = int(NumberRecode) / int(req.Limit)
	if totalPages == 0 {
		return 
	}

	return
}
