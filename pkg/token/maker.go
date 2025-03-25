package token

type Maker interface {
	CreateToken(id int64)
	VerifyToken(token string)
}

type JWTMaker struct {
}
