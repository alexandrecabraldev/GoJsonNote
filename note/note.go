package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"time"

	"github.com/google/uuid"
)

type Note struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// exibe a nota criada
func (n Note) Display() {
	fmt.Printf("id: %s\ntitle: %s\ncontent:%s\ncreated_at:%s\n", n.ID, n.Title, n.Content, n.CreatedAt)
}

// salva a nota no arquivo notes.json
func (n Note) Save(fileContent map[string]Note) error {

	fileContent[n.ID] = n

	file, err := os.OpenFile("notes.json", os.O_CREATE|os.O_WRONLY, fs.FileMode(0644))
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if fileInfo.Size() > 0 {
		fileBytes, err := os.ReadFile("notes.json")
		if err != nil {
			fmt.Println(err)
			return err
		}

		err = json.Unmarshal(fileBytes, &fileContent)
		if err != nil {
			fmt.Println(err)
			return err
		}

	}
	fileContent[n.ID] = n
	byteNote, err := json.Marshal(fileContent)
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = file.WriteString(string(byteNote))
	if err != nil {
		fmt.Println("Error on write file ", err)
		return err
	}

	return nil
}

// construtor para cria uma nova nota
func New(title, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("title and Content cannot be empty")
	}

	return Note{
		ID:        uuid.New().String(),
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
