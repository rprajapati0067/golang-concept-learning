package event

import (
	"errors"
	"unicode/utf8"

	"github.com/rprajapati0067/golang-concept-learning/14_encapsulation/calendar"
)

type Event struct {
	title string
	calendar.Date
}

func (e *Event) Title() string {
	return e.title
}
func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("invalid title")
	}
	e.title = title
	return nil
}
