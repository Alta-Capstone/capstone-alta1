package repository

import (
	client "capstone-alta1/features/client"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type clientRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) client.RepositoryInterface {
	return &clientRepository{
		db: db,
	}
}

// Create implements user.Repository
func (repo *clientRepository) Create(input client.Core) error {
	clientGorm := fromCore(input)
	tx := repo.db.Create(&clientGorm) // proses insert data
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}

// GetAll implements user.Repository
func (repo *clientRepository) GetAll() (data []client.Core, err error) {
	var client []Client

	tx := repo.db.Preload("User").Find(&client)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(client)
	return dataCore, nil
}

func (repo *clientRepository) GetAllWithSearch(query string) (data []client.Core, err error) {
	var client []Client
	tx := repo.db.Preload("User").Where("name LIKE ?", "%"+query+"%").Find(&client)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	fmt.Println("\n\n 2 getall client = ", client)

	var dataCore = toCoreList(client)
	return dataCore, nil
}

// GetById implements user.RepositoryInterface
func (repo *clientRepository) GetById(id uint) (data client.Core, err error) {
	var client Client

	tx := repo.db.Preload("User").First(&client, id)

	if tx.Error != nil {
		return data, tx.Error
	}

	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	var dataCore = client.toCore()
	return dataCore, nil
}

// Update implements user.Repository
func (repo *clientRepository) Update(input client.Core, clientID uint, userID uint) error {
	clientGorm := fromCore(input)
	var client Client
	var user User
	fmt.Println("\n\nInput update client ", input)
	fmt.Println("\n\nclientID ", clientID, " --- ", "userID", userID)

	tx := repo.db.Model(&user).Where("ID = ?", userID).Updates(&clientGorm)
	yx := repo.db.Model(&client).Where("ID = ?", clientID).Updates(&clientGorm)
	if tx.Error != nil && yx.Error != nil {
		return errors.New("failed update client")
	}
	if tx.RowsAffected == 0 {
		return errors.New("update failed")
	}
	return nil
	// return repo.db.Transaction(func(tx *gorm.DB) (err error) {
	// 	// do some database operations in the transaction (use 'tx' from this point, not 'db')

	// 	fmt.Println("\n\ntx = ", tx)
	// 	fmt.Println("\n\nerr = ", err)

	// 	err = tx.Model(&client).Where("user_id = ?", userID).Updates(&clientGorm).Error
	// 	if err != nil {
	// 		fmt.Println("\n\nFailed Update. Query data user failed. Error : ", err)
	// 		return err
	// 	}

	// 	fmt.Println("\n\ntx2 = ", tx)
	// 	fmt.Println("\n\nerr2 = ", err)

	// 	// if tx.RowsAffected == 0 {
	// 	// 	fmt.Println("\n\nUpdate user failed row affected ", tx.RowsAffected)
	// 	// 	return errors.New("Update failed")
	// 	// }
	// 	fmt.Println("\n\nUpdate user succes  row affected = ", tx.RowsAffected, " hasil = ", clientGorm)

	// 	if err = tx.Model(&client).Where("ID = ?", clientID).Updates(&clientGorm).Error; err != nil {
	// 		// return any error will rollback
	// 		fmt.Println("\n\nFailed Update. Query data client failed. Error : ", err)
	// 		return err
	// 	}

	// 	fmt.Println("\n\nUpdate client succes  row affected = ", tx.RowsAffected, " hasil = ", clientGorm)

	// 	// if tx.RowsAffected == 0 {
	// 	// 	return errors.New("Update failed")
	// 	// }

	// 	// return nil will commit the whole transaction
	// return
}

// })
// }

// Delete implements user.Repository
func (repo *clientRepository) Delete(id uint) error {
	var client Client
	var user User
	repo.db.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.Delete(&client, id).Error; err != nil {
			// return any error will rollback
			fmt.Println("\n\nFailed Delete. Query data client failed. Error : ", err)
			return err
		}

		if err := tx.Delete(&user, id).Error; err != nil {
			fmt.Println("\n\nFailed Update. Query data user failed. Error : ", err)
			return err
		}

		if tx.RowsAffected == 0 {
			fmt.Println("\n\nUpdate failed row affected ", tx.RowsAffected)
			return errors.New("update failed")
		}

		// return nil will commit the whole transaction
		return nil
	})

	return nil
}

func (repo *clientRepository) FindUser(email string) (result client.Core, err error) {
	var clientData Client
	tx := repo.db.Where("email = ?", email).First(&clientData.User)
	if tx.Error != nil {
		return client.Core{}, tx.Error
	}

	result = clientData.toCore()

	return result, nil
}

func (repo *clientRepository) GetOrderById(id uint) (data []client.Order, err error) {
	var clientorder []Order

	tx := repo.db.Find(&clientorder, id)

	if tx.Error != nil {
		return data, tx.Error
	}

	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	var dataCore = toCoreListOrder(clientorder)
	return dataCore, nil
}
