package dialog

import (
	"github.com/legacy-ai/floyd/internal/ui/styles"
	"github.com/sahilm/fuzzy"
)

// CommandItem wraps a uicmd.Command to implement the ListItem interface.
type CommandItem struct {
	id          string
	title       string
	description string
	shortcut    string
	action      Action
	t           *styles.Styles
	m           fuzzy.Match
	cache       map[int]string
	focused     bool
	expanded    bool
}

var _ ListItem = &CommandItem{}

// NewCommandItem creates a new CommandItem.
func NewCommandItem(t *styles.Styles, id, title, shortcut string, action Action) *CommandItem {
	return &CommandItem{
		id:       id,
		t:        t,
		title:    title,
		shortcut: shortcut,
		action:   action,
	}
}

// NewCommandItemWithDescription creates a new CommandItem with a description.
func NewCommandItemWithDescription(t *styles.Styles, id, title, description, shortcut string, action Action) *CommandItem {
	return &CommandItem{
		id:          id,
		t:           t,
		title:       title,
		description: description,
		shortcut:    shortcut,
		action:      action,
	}
}

// Filter implements ListItem.
func (c *CommandItem) Filter() string {
	return c.title + " " + c.description
}

// ID implements ListItem.
func (c *CommandItem) ID() string {
	return c.id
}

// SetFocused implements ListItem.
func (c *CommandItem) SetFocused(focused bool) {
	if c.focused != focused {
		c.cache = nil
	}
	c.focused = focused
}

// SetMatch implements ListItem.
func (c *CommandItem) SetMatch(m fuzzy.Match) {
	c.cache = nil
	c.m = m
}

// Action returns the action associated with the command item.
func (c *CommandItem) Action() Action {
	return c.action
}

// SetTitle updates the display title and clears the render cache.
func (c *CommandItem) SetTitle(title string) {
	c.title = title
	c.cache = nil
}

// Shortcut returns the shortcut associated with the command item.
func (c *CommandItem) Shortcut() string {
	return c.shortcut
}

// Render implements ListItem.
func (c *CommandItem) Render(width int) string {
	styles := ListItemStyles{
		ItemBlurred:     c.t.Dialog.NormalItem,
		ItemFocused:     c.t.Dialog.SelectedItem,
		InfoTextBlurred: c.t.Base,
		InfoTextFocused: c.t.Base,
	}

	displayTitle := c.title
	displayDescription := c.shortcut

	if c.description != "" {
		if c.expanded {
			displayDescription = c.description
		} else if c.shortcut != "" {
			displayTitle = c.title + " (" + c.shortcut + ")"
			displayDescription = c.description
		} else {
			displayDescription = c.description
		}
	}

	return renderItem(styles, displayTitle, displayDescription, c.focused, width, c.cache, &c.m)
}
