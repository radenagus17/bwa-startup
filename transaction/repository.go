package transaction

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetByCampaignID(campaignID int) ([]Transaction, error)
	GetByUserID(userID int) ([]Transaction, error)
	GetByID(ID int) (Transaction, error)
	Save(input Transaction) (Transaction,error)
	Update(input Transaction) (Transaction,error)
}

func NewRepository(db *gorm.DB) *repository{
	return &repository{db}
}

func (r *repository) GetByCampaignID(campaignID int) ([]Transaction, error){
	var transactions []Transaction

	err := r.db.Preload("User").Order("id desc").Where("campaign_id = ?", campaignID).Find(&transactions).Error

	if err != nil{
		return transactions, err
	}

	return transactions, nil
}

func (r *repository) GetByUserID(userID int) ([]Transaction, error){
	var transactions []Transaction

	err := r.db.Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Where("user_id = ?", userID).Order("id DESC").Find(&transactions).Error

	if err != nil{
		return transactions, err
	}

	return transactions, nil
}

func (r *repository) GetByID(ID int) (Transaction, error) {
	var transaction Transaction

	err := r.db.Where("id = ?", ID).Find(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) Save(input Transaction) (Transaction,error){
	err := r.db.Create(&input).Error

	if err != nil{
		return input, err
	}

	return input, nil
}

func (r *repository) Update(input Transaction) (Transaction,error){
	err := r.db.Save(&input).Error

	if err != nil{
		return input, err
	}

	return input, nil
}
