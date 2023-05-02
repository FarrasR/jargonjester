package domain

type PermissionRepository interface {
	AddChannelPermission(string) error
	RemoveChannelPermission(string) error
	CheckChannelPermission(string) bool

	AdduserPermission(string) error
	RemoveUserPermission(string) error
	CheckUserPermission(string) bool
}
