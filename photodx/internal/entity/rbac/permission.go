package rbac

type Permission struct {
	Resource Resource
	Action   Action
}
