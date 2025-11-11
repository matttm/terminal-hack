package container

import (
	"math"

	"github.com/gdamore/tcell"
)

// message represents a single text message to be displayed.
type message struct {
	text string
}

// MessageContainer is a specialized container for displaying a scrolling list of messages.
// It maintains a history of messages and automatically handles text wrapping and layout.
type MessageContainer struct {
	gui      *Container // The underlying container for rendering
	messages []message  // List of messages to display
}

// CreateMessageContainer creates a new message container with the specified dimensions and position.
func CreateMessageContainer(s tcell.Screen, x1, y1, dy, dx int) *MessageContainer {
	mc := new(MessageContainer)
	mc.gui = NewContainer(s, x1, y1, dy, dx)
	mc.messages = make([]message, 3)
	return mc

}

// ClearContainer clears all content from the message container.
func (mc *MessageContainer) ClearContainer() {
	mc.gui.ClearContainer()
}

// RenderContainer renders the message container's border to the screen.
func (mc *MessageContainer) RenderContainer() {
	mc.gui.RenderContainer()
}

// AddNewMessage adds a new message to the container and re-renders all messages.
// Messages are displayed in reverse chronological order (newest first) with one blank line between each.
func (mc *MessageContainer) AddNewMessage(s string) {
	mc.messages = append(mc.messages, message{text: s})
	mc.ClearContainer()
	pos := 0
	for i := len(mc.messages) - 1; i >= 0; i-- {
		m := mc.messages[i]
		lines := mc.getLineCountOfMessage(m.text)
		_, pos = mc.gui.WriteLineAtPosition(pos, lines, m.text)
		pos = pos - mc.gui.y1 + 1
	}
}

// getLineCountOfMessage calculates how many lines a message will occupy
// based on the container's width, accounting for text wrapping.
func (mc *MessageContainer) getLineCountOfMessage(s string) int {
	// len() is fine here as all characters should be ascii range
	lines := float64(len(s)) / float64(mc.gui.columns)
	return int(math.Ceil(lines))
}
