package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct{
	Text string `json:"text"`
}

func New(content string) (Todo,error){
	if content == ""  {
		return Todo{}, errors.New("text is required")
	}
	return Todo{
		Text: content,
	},nil
}

func(todo Todo)Display (){
	fmt.Println((todo.Text))
}

func (todo Todo) Save() error{
	fileName := "todo.json"
	json,err:=json.Marshal(todo)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return  os.WriteFile(fileName,json,0666)
	
}