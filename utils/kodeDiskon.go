package utils

import (
	"projek/toko-retail/model"
	repository "projek/toko-retail/repository/config"
	"projek/toko-retail/repository/modelfunc"
	"time"

	"github.com/siruspen/logrus"
)

func CreateKodeDiskon(data model.Diskon) (model.Diskon, error) {
	repoDiskon := modelfunc.Diskon{
		Diskon: data,
	}

	repoDiskon.CreatedAt = time.Now()
	repoDiskon.UpdatedAt = time.Now()

	err := repoDiskon.CreateDiskon(repository.Mysql.DB)
	if err != nil {
		return model.Diskon{}, err
	}

	return repoDiskon.Diskon, nil
}

// function untuk mendapatkan semua diskon
func GetDiskon() ([]model.Diskon, error) {
	var diskon modelfunc.Diskon
	repoDiskons, err := diskon.GetAll(repository.Mysql.DB)
	if err!= nil {
        return nil, err
    }

	var result []model.Diskon
	for _, repoDiskon := range repoDiskons {
		result = append(result, repoDiskon.Diskon)
	}

	return result, nil
}

// function untuk mendapatkan diskon berdasarkan kode
func GetDiskonByCode(s string) (model.Diskon, error) {
    logrus.Println("Searching for discount code:", s) // Log the search parameter
    diskon := modelfunc.Diskon{
        Diskon: model.Diskon{
            KodeDiskon: s,
        },
    }

    repoDiskon, err := diskon.GetByCode(repository.Mysql.DB)
    if err != nil {
        logrus.Println("Error in GetByCode:", err) // Log the error
        return model.Diskon{}, err
    }

    logrus.Println("Discount code found:", repoDiskon.Diskon) // Log the found record
    return repoDiskon.Diskon, nil
}


// function untuk mendapatkan diskon berdasarkan ID
func GetDiskonByID(id uint) (model.Diskon, error) {
	diskon := modelfunc.Diskon{
		Diskon: model.Diskon{
			ID: id,
		},
	}

	repoDiskon, err := diskon.GetByID(repository.Mysql.DB)
	if err != nil {
		return model.Diskon{}, err
	}

	return repoDiskon.Diskon, nil
}

// function untuk memperbarui diskon
func UpdateDiskon(id uint, updatedDiskon model.Diskon) (model.Diskon, error) {
	existingDiskon := modelfunc.Diskon{
		Diskon: model.Diskon{
			ID: id,
		},
	}

	if err := repository.Mysql.DB.First(&existingDiskon.Diskon).Error; err != nil {
		return model.Diskon{}, err
	}

	existingDiskon.Amount = updatedDiskon.Amount
	existingDiskon.Type = updatedDiskon.Type
	existingDiskon.UpdatedAt = time.Now()

	if err := repository.Mysql.DB.Save(&existingDiskon.Diskon).Error; err != nil {
		return model.Diskon{}, err
	}

	return existingDiskon.Diskon, nil
}

// function untuk menghapus diskon berdasarkan ID
func DeleteKode(id uint64) error {
	diskon := modelfunc.Diskon{
		Diskon: model.Diskon{
			ID: uint(id),
		},
	}

	return diskon.Delete(repository.Mysql.DB)
}