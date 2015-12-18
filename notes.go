package safenotes_core

import (
	"github.com/pborman/uuid"
	"time"
	"encoding/json"
)

type Note struct {
	Id string
	Title string
	Content string
	Timestamp int64
}

var notes []Note = make([]Note, 0)

func GetNotes() string {
	json := serializeNotes()
	return json
}

func serializeNotes() string {
	json, _ := json.Marshal(notes)
	return string(json)
}

func AddNote(title string, content string) {
	note := makeNote(title, content)
	notes = append(notes, note)
}

func makeNote(title string, content string) Note {
	return Note{
		Id: uuid.NewRandom().String(),
		Title: title,
		Content: content,
		Timestamp: time.Now().Unix(),
	}
}
