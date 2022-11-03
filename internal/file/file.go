package file

import (
	"fmt"

	"github.com/google/uuid"
)

type File struct {
	ID   uuid.UUID
	Name string
	Data []byte
}

func NewFile(filename string, fileData []byte) (*File, error) {
	FileId, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return &File{
		ID:   FileId,
		Name: filename,
		Data: fileData,
	}, nil
}
func (f *File) String() string {
	return fmt.Sprintf("File(%s, %s)", f.Name, f.ID)
}
