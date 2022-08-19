package models

import "gorm.io/gorm"

type MlModel struct {
	gorm.Model
	Name           string
	Ratio_of_train float32
	Look_back      int
	Learning_rate  float32
	Epochs         int
	Batch_size     int
	UserID         uint
}
