package scenes

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type NotJoinedHandler = func(update tg.Update)

type user struct {
	scene *Scene
	ctx   *Context
}

type ScenesManager struct {
	users map[int64]*user
}

func NewScenesManager(scenes ...*Scene) *ScenesManager {
	return &ScenesManager{
		users: map[int64]*user{},
	}
}

func (s *ScenesManager) Join(update tg.Update, scene *Scene) {
	userId, err := getUserId(update)
	if err != nil {
		return
	}

	ctx := NewContext()
	user := &user{scene, ctx}

	s.users[userId] = user

	if scene.joinHandler != nil {
		scene.joinHandler(update)
	}
}

func (s *ScenesManager) Handle(update tg.Update, onNotJoinedCase NotJoinedHandler) {
	userId, err := getUserId(update)
	if err != nil {
		return
	}

	if _, ok := s.users[userId]; !ok {
		onNotJoinedCase(update)
		return
	}

	user := s.users[userId]
	scene := user.scene
	ctx := user.ctx

	updateData := &UpdateData{Ctx: ctx, Update: update}

	scene.makeOneStep(updateData)

	if scene.leftUsers[userId] {
		delete(s.users, userId)
	}

	scene.leftUsers[userId] = false
}
