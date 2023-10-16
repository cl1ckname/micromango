package user

type Config struct {
	Addr      string
	DbAddr    string
	Salt      string
	JwtSecret string
}
