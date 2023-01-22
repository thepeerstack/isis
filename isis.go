package isis

import (
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

func Generate(conn *gorm.DB, identifier string, digits int, validity int) (token string, err error) {
	token = generatePin(digits)

	err = conn.Create(&otp{
		Identifier: identifier,
		Token:      token,
		Validity:   validity,
		Valid:      true,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}).Error

	return
}

func Validate(conn *gorm.DB, identifier string, token string) (validated bool, err error) {
	var foundOtp otp

	err = conn.Model(otp{}).Where("identifier = ?", identifier).Where("token = ?", token).Find(&foundOtp).Error

	if err != nil {
		return false, err
	}

	if foundOtp.Empty() || !foundOtp.Valid {
		return false, nil
	}

	conn.First(&foundOtp)

	if foundOtp.Exipired() {
		foundOtp.Valid = false
		foundOtp.UpdatedAt = time.Now()
		conn.Save(&foundOtp)

		return false, nil
	}

	foundOtp.Valid = false
	foundOtp.UpdatedAt = time.Now()
	conn.Save(&foundOtp)

	return true, nil
}

func generatePin(digits int) string {
	pin := ""

	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 9

	for i := 0; i < digits; i++ {
		digit := fmt.Sprintf("%d", rand.Intn(max-min+1)+min)
		pin = pin + digit
	}

	return pin
}
