package storages

type Storage interface {
	Save(string) (string,error)
	Load(string) (string,error)
}
