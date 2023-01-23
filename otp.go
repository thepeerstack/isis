package isis

import "time"

type otp struct {
	Id         int64
	Identifier string
	Token      string
	Validity   int
	Valid      bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (otp otp) Empty() bool {
	if otp.Id == 0 {
		return true
	}

	return false
}

func (otp otp) Expired() bool {
	if time.Now().Sub(otp.CreatedAt).Minutes() > float64(otp.Validity) {
		return true
	}

	return false
}
