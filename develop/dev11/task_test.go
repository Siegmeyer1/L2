package main

import (
	"L2/develop/dev11/internal/calendar"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddDelete(t *testing.T) {
	c := calendar.NewCalendar()
	e := calendar.NewEvent()
	c.AddEventToCalendar(e)
	ok := c.DeleteEvent(e.ID)
	assert.Equal(t, true, ok)
}
