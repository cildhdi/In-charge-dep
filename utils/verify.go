package utils

import (
	"errors"
	"strconv"
	"time"

	"github.com/cildhdi/In-charge/models"
)

func CodeVerify(phone string, codeStr string, reuse bool) error {
	code, err := strconv.Atoi(codeStr)
	if err != nil || code < 1000 {
		return errors.New("code is not a number's string format or its value isnt in range")
	}

	var vc models.VerificationCode

	models.IcDb().Where(&models.VerificationCode{
		Phone: phone,
		Code:  uint(code),
	}).First(&vc)
	if vc.ID == 0 || time.Now().Sub(vc.CreatedAt).Minutes() > 5 {
		return errors.New("invalid code")
	}
	if !reuse {
		models.IcDb().Delete(&vc)
	}
	return nil
}
