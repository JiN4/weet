package model

// formatTableにInsertする
func CreateFormat(format Format) (Format, error) {
	err := db.Create(&format).Error
	if err != nil {
		return Format{}, err
	}
	return format, nil
}
