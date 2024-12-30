package container

import (
	"math"

	"github.com/gdamore/tcell"
)

type message struct {
	text string
}

type MessageContainer struct {
	gui      *Container
	messages []message
}

func CreateMessageContainer(s tcell.Screen, x1, y1, dy, dx int) *MessageContainer {
	mc := new(MessageContainer)
	mc.gui = NewContainer(s, x1, y1, dy, dx)
	mc.messages = make([]message, 3)
	return mc

}
func (mc *MessageContainer) clearBoard() {
	mc.gui.clearBoard()
}

func (mc *MessageContainer) RenderContainer() {
	mc.gui.RenderContainer()
}

func (mc *MessageContainer) AddNewMessage(s string) {
	mc.messages = append(mc.messages, message{text: s})
	mc.clearBoard()
	lines := mc.getLineCountOfMessage(s)
	pos := 0
	for i := len(mc.messages) - 1; i >= 0; i-- {
		m := mc.messages[i]
		_, pos = mc.gui.WriteLineAtPosition(pos, lines, m.text)
	}
}
func (mc *MessageContainer) getLineCountOfMessage(s string) int {
	// len() is fine here as all characters should be ascii range
	lines := float64(len(s)) / float64(mc.gui.columns)
	return int(math.Ceil(lines))
}
