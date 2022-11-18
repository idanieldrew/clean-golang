package user

type Request struct {
	Email string `json:"email"`
	//Phone string `json:"phone"`
}

func (u *Request) Validation(r int) bool {
	if r == 1 {
		return false
	}
	return true
}
