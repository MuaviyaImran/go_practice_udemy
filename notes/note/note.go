package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct{
	Title string `json:"title"`
	Content string `json:"content"`
	CreateAt time.Time `json:"created_at"`
}

func New(title,content string) (Note,error){
	if title == "" || content == "" {
		return Note{}, errors.New("title and content are required")
	}
	return Note{
		Title: title,
		Content: content,
		CreateAt: time.Now(),
	},nil
}

func(note Note)Display (){
	fmt.Printf("Your note titled %v has the following content: \n\n%v\n\n",note.Title,note.Content)
}

func (note Note) Save() error{
	fileName := strings.ReplaceAll(note.Title," ","_")
	fileName = strings.ToLower(fileName)
	json,err:=json.Marshal(note)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return  os.WriteFile(fileName+".json",json,0666)
	
}