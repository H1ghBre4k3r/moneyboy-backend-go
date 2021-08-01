package auth

type JWT interface {
	Sign(map[string]interface{}) (string, error)
}
