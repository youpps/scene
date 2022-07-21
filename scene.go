package scenes

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UpdateData struct {
	Update tg.Update
	Ctx    *Context
}

type Handler = func(updateData *UpdateData)

type Scene struct {
	Name           string
	usersStages    map[int64]int
	leftUsers      map[int64]bool
	handlers       []Handler
	defaultHandler Handler
	joinHandler    NotJoinedHandler
}

func NewScene(name string) *Scene {
	return &Scene{
		Name:           name,
		usersStages:    map[int64]int{},
		leftUsers:      map[int64]bool{},
		handlers:       []Handler{},
		defaultHandler: nil,
		joinHandler:    nil,
	}
}

func (s *Scene) makeOneStep(updateData *UpdateData) {
	userId, err := getUserId(updateData.Update)
	if err != nil {
		return
	}

	if _, ok := s.usersStages[userId]; !ok {
		s.usersStages[userId] = 1
		s.leftUsers[userId] = false
	}

	if s.usersStages[userId] > len(s.handlers) {
		s.usersStages[userId] = 0
	}

	if s.usersStages[userId] == 0 {
		if s.defaultHandler != nil {
			s.defaultHandler(updateData)

			ctx := updateData.Ctx

			if ctx.isNext {
				s.usersStages[userId]++
			}

			if ctx.isLeft {
				s.leftUsers[userId] = true
				s.usersStages[userId] = 1
			}
		}
		return
	}

	s.handlers[s.usersStages[userId]-1](updateData)

	ctx := updateData.Ctx

	if ctx.isNext {
		s.usersStages[userId]++
	}

	if ctx.isLeft {
		s.leftUsers[userId] = true
		s.usersStages[userId] = 1
	}

	ctx.Cancel()
}

func (s *Scene) AddDefaultHandler(handler Handler) {
	s.defaultHandler = handler
}

func (s *Scene) AddNextHandler(handler Handler) {
	s.handlers = append(s.handlers, handler)
}

func (s *Scene) AddJoinHandler(handler NotJoinedHandler) {
	s.joinHandler = handler
}
