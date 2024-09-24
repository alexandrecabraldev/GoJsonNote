package main

import (
	"JsonNote/GoJsonNote/note"
	"bufio"
	"fmt"
	"os"
)

var FileContent = make(map[string]note.Note)

func main() {

	valueTitle, valueContent := getNoteData()

	//cria uma nova nota passando os inputs para ela
	note, err := note.New(valueTitle, valueContent)
	if err != nil {
		fmt.Println(err)
		return
	}

	//salva no arquivo notes.json
	err = note.Save(FileContent)
	if err != nil {
		return
	}

}

// executa a função getUserInput capturando o input completo
func getNoteData() (string, string) {
	title := getUserInput("Enter with note title: ")

	content := getUserInput("Enter with note content: ")

	return title, content
}

// captura os inputs do usuário
func getUserInput(prompt string) string {
	fmt.Println(prompt)
	var inputValue string

	reader := bufio.NewReader(os.Stdin)
	inputValue, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return inputValue
}
