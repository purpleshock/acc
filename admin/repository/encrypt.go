package repository

type Encrypt interface {
	Digest(raw string) string
}
