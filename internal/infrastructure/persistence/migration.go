package persistence

import "gorm.io/gorm"

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&DelegationDTO{},
		&SettingsDTO{},
	)
}

func Seed(db *gorm.DB) error {
	// Settingsを読み出す
	err := db.First(&SettingsDTO{ID: 1}).Error
	if err == nil {
		// エラーがなければ問題なし（登録済み）
		return nil
	} else if err != nil && err != gorm.ErrRecordNotFound {
		// ErrRecordNotFound 以外のエラーの場合は失敗
		return err
	}
	// Settings未登録の場合はエントリを作成
	return db.Save(&SettingsDTO{ID: 1}).Error
}
