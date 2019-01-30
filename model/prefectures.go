package model

// parametersTableにInsertする
func CreatePrefectures(prefectures Prefectures) (Prefectures, error) {
	err := db.Create(&prefectures).Error
	if err != nil {
		return Prefectures{}, err
	}
	return prefectures, nil
}

func GetPrefecturesNameById(id uint) string {
	prefectures := Prefectures{}
	db.Where("id = ?", id).First(&prefectures)

	return prefectures.Name
}
