package platform

type Store interface {
	Get(id string) (map[string]any, bool)
	Put(id string, value map[string]any)
	Delete(id string)
	HasKey(id string) bool
	GetAll() []map[string]any
}
