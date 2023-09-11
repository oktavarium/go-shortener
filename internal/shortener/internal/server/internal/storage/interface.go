package storage

type Storage interface {
	Save(k, v string)
	Get(k string) (string, bool)
}
