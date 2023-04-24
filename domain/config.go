package domain

type ConfigRepository interface {
	IsLimited(string) error
}
