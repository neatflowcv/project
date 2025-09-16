package filesystem

type Filesystem interface {
	CreateDirectory(path string) error
	CreateFile(path string, content []byte) error
}
