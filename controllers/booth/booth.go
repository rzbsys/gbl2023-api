package booth

import (
	"gbl-api/data"
	"math/rand"

	"gorm.io/gorm"
)

func GetBooths() ([]Booth, error) {
	db := data.GetDatabase()
	var booths []Booth
	err := db.Find(&booths).Error
	return booths, err
}

func GetBooth(bid string) (Booth, error) {
	db := data.GetDatabase()
	var booth Booth
	err := db.Where("b_id = ?", bid).First(&booth).Error
	return booth, err
}

func GetBoothIdByPassword(password string) (string, error) {
	db := data.GetDatabase()
	var boothPw BoothPassword
	err := db.Where("password = ?", password).First(&boothPw).Error
	if err == gorm.ErrRecordNotFound {
		return "", nil
	} else if err != nil {
		return "", err
	} else {
		return boothPw.BID, nil
	}
}

func DeleteBooth(bid string) error {
	db := data.GetDatabase()
	return db.Delete(&Booth{}, "bid = ?", bid).Error
}

func generateRandomString(n int) string {
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = charSet[rand.Intn(len(charSet))]
	}

	return string(b)
}

func generateRandomBID() string {
	return generateRandomString(64)
}

func AddPassword(password string) error {
	db := data.GetDatabase()
	return db.Create(&BoothPassword{
		Password: password,
		BID:      generateRandomBID(),
	}).Error
}

func DeletePasswordByBID(bid string) error {
	db := data.GetDatabase()
	return db.Delete(&BoothPassword{}, "bid = ?", bid).Error
}
