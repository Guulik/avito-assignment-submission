package validator

const (
	Invalid = 0
	User    = 1
	Admin   = 2
)

func Validate(token string) int {
	if !valid(token) {
		return 0
	}
	//TODO check if token contains required data (e.x. is_admin = false)
	//if not admin
	return 0
	//if admin
	return 1
}

func valid(token string) bool {
	//TODO implement me
	return true
	return false
}
