package dialog

import (
	"fmt"
	"strings"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	uv "github.com/charmbracelet/ultraviolet"
	"github.com/sahilm/fuzzy"

	"github.com/legacy-ai/floyd/internal/ui/common"
	"github.com/legacy-ai/floyd/internal/ui/list"
	"github.com/legacy-ai/floyd/internal/ui/styles"
)

const (
	// PluginsLibraryID is the identifier for the plugins library dialog.
	PluginsLibraryID              = "plugins_library"
	pluginsLibraryDialogMaxWidth  = 80
	pluginsLibraryDialogMaxHeight = 24
)

// PluginStatus represents the status of a plugin.
type PluginStatus int

const (
	PluginStatusConnected PluginStatus = iota
	PluginStatusAvailable
	PluginStatusError
)

// PluginInfo represents information about a plugin/MCP server.
type PluginInfo struct {
	Name        string
	Description string
	Status      PluginStatus
	ToolsCount  int
	Error       string
}

// PluginsLibrary represents a dialog for managing plugins (MCP servers).
type PluginsLibrary struct {
	com     *common.Common
	help    help.Model
	list    *list.FilterableList
	input   textinput.Model
	plugins []PluginInfo

	// Category system (Connected/Available)
	categories       []string
	selectedCategory int

	keyMap struct {
		Select       key.Binding
		Next         key.Binding
		Previous     key.Binding
		UpDown       key.Binding
		Close        key.Binding
		CategoryNext key.Binding
		CategoryPrev key.Binding
		FocusFilter  key.Binding
		Connect      key.Binding
	}
}

// PluginsLibraryItem represents a plugin list item.
type PluginsLibraryItem struct {
	plugin  PluginInfo
	t       *styles.Styles
	m       fuzzy.Match
	cache   map[int]string
	focused bool
}

var (
	_ Dialog   = (*PluginsLibrary)(nil)
	_ ListItem = (*PluginsLibraryItem)(nil)
)

// NewPluginsLibrary creates a new plugins library dialog.
func NewPluginsLibrary(com *common.Common) (*PluginsLibrary, error) {
	p := &PluginsLibrary{com: com}

	help := help.New()
	help.Styles = com.Styles.DialogHelpStyles()
	p.help = help

	p.list = list.NewFilterableList()
	p.list.Focus()

	p.input = textinput.New()
	p.input.SetVirtualCursor(false)
	p.input.Placeholder = "Type to filter plugins..."
	p.input.SetStyles(com.Styles.TextInput)
	p.input.Focus()

	// Categories for plugins
	p.categories = []string{"All", "Connected", "Available"}
	p.selectedCategory = 0

	// Key bindings
	p.keyMap.Select = key.NewBinding(
		key.WithKeys("enter", "ctrl+y"),
		key.WithHelp("enter", "select"),
	)
	p.keyMap.Next = key.NewBinding(
		key.WithKeys("down", "ctrl+n"),
		key.WithHelp("↓", "next"),
	)
	p.keyMap.Previous = key.NewBinding(
		key.WithKeys("up", "ctrl+p"),
		key.WithHelp("↑", "prev"),
	)
	p.keyMap.UpDown = key.NewBinding(
		key.WithKeys("up", "down"),
		key.WithHelp("↑/↓", "navigate"),
	)
	p.keyMap.Close = CloseKey
	p.keyMap.CategoryNext = key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "next category"),
	)
	p.keyMap.CategoryPrev = key.NewBinding(
		key.WithKeys("shift+tab"),
		key.WithHelp("⇧+tab", "prev category"),
	)
	p.keyMap.FocusFilter = key.NewBinding(
		key.WithKeys("/"),
		key.WithHelp("/", "search"),
	)
	p.keyMap.Connect = key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "connect"),
	)

	p.loadPlugins()

	return p, nil
}

// loadPlugins loads plugin information.
func (p *PluginsLibrary) loadPlugins() {
	// TODO: Load actual MCP server status from config
	// For now, placeholder data
	p.plugins = []PluginInfo{
		{Name: "filesystem", Description: "File system operations (read, write, search)", Status: PluginStatusConnected, ToolsCount: 8},
		{Name: "github", Description: "GitHub API integration (repos, issues, PRs)", Status: PluginStatusAvailable},
		{Name: "postgres", Description: "PostgreSQL database operations", Status: PluginStatusConnected, ToolsCount: 5},
	}

	p.updateListItems()
}

// updateListItems updates the list items based on selected category.
func (p *PluginsLibrary) updateListItems() {
	var filtered []PluginInfo

	switch p.selectedCategory {
	case 0: // All
		filtered = p.plugins
	case 1: // Connected
		for _, plugin := range p.plugins {
			if plugin.Status == PluginStatusConnected {
				filtered = append(filtered, plugin)
			}
		}
	case 2: // Available
		for _, plugin := range p.plugins {
			if plugin.Status == PluginStatusAvailable {
				filtered = append(filtered, plugin)
			}
		}
	}

	items := make([]list.FilterableItem, 0, len(filtered))
	for _, plugin := range filtered {
		items = append(items, &PluginsLibraryItem{
			plugin: plugin,
			t:      p.com.Styles,
		})
	}
	p.list.SetItems(items...)
	p.list.SetFilter(p.input.Value())
	p.list.ScrollToTop()
	p.list.SetSelected(0)
}

// categoryTabBar renders the category tab bar.
func (p *PluginsLibrary) categoryTabBar(width int) string {
	t := p.com.Styles
	var tabs []string

	for i, cat := range p.categories {
		var tab string

		if i == p.selectedCategory {
			tab = t.Base.Render("● ") + t.Base.Bold(true).Render(cat)
		} else {
			tab = t.HalfMuted.Render(cat)
		}
		tabs = append(tabs, tab)
	}

	return strings.Join(tabs, t.HalfMuted.Render(" │ "))
}

// ID implements Dialog.
func (p *PluginsLibrary) ID() string {
	return PluginsLibraryID
}

// HandleMsg implements Dialog.
func (p *PluginsLibrary) HandleMsg(msg tea.Msg) Action {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, p.keyMap.Close):
			return ActionClose{}

		case key.Matches(msg, p.keyMap.CategoryNext):
			p.selectedCategory = (p.selectedCategory + 1) % len(p.categories)
			p.updateListItems()

		case key.Matches(msg, p.keyMap.CategoryPrev):
			newIdx := p.selectedCategory - 1
			if newIdx < 0 {
				newIdx = len(p.categories) - 1
			}
			p.selectedCategory = newIdx
			p.updateListItems()

		case key.Matches(msg, p.keyMap.Previous):
			p.list.Focus()
			if p.list.IsSelectedFirst() {
				p.list.SelectLast()
				p.list.ScrollToBottom()
				break
			}
			p.list.SelectPrev()
			p.list.ScrollToSelected()

		case key.Matches(msg, p.keyMap.Next):
			p.list.Focus()
			if p.list.IsSelectedLast() {
				p.list.SelectFirst()
				p.list.ScrollToTop()
				break
			}
			p.list.SelectNext()
			p.list.ScrollToSelected()

		case key.Matches(msg, p.keyMap.Select):
			selectedItem := p.list.SelectedItem()
			if selectedItem == nil {
				break
			}
			pluginItem, ok := selectedItem.(*PluginsLibraryItem)
			if !ok {
				break
			}
			return ActionSelectPlugin{
				PluginName:        pluginItem.plugin.Name,
				PluginDescription: pluginItem.plugin.Description,
				PluginStatus:      pluginItem.plugin.Status,
			}

		case key.Matches(msg, p.keyMap.FocusFilter):
			p.input.Focus()

		default:
			var cmd tea.Cmd
			p.input, cmd = p.input.Update(msg)
			value := p.input.Value()
			p.list.SetFilter(value)
			p.list.ScrollToTop()
			p.list.SetSelected(0)
			return ActionCmd{cmd}
		}
	}
	return nil
}

// Cursor returns cursor position.
func (p *PluginsLibrary) Cursor() *tea.Cursor {
	return InputCursor(p.com.Styles, p.input.Cursor())
}

// Draw implements Dialog.
func (p *PluginsLibrary) Draw(scr uv.Screen, area uv.Rectangle) *tea.Cursor {
	t := p.com.Styles
	width := max(0, min(pluginsLibraryDialogMaxWidth, area.Dx()))
	height := max(0, min(pluginsLibraryDialogMaxHeight, area.Dy()))
	innerWidth := width - t.Dialog.View.GetHorizontalFrameSize()

	categoryBarHeight := 1
	heightOffset := t.Dialog.Title.GetVerticalFrameSize() + titleContentHeight +
		t.Dialog.InputPrompt.GetVerticalFrameSize() + inputContentHeight +
		categoryBarHeight +
		t.Dialog.HelpView.GetVerticalFrameSize() +
		t.Dialog.View.GetVerticalFrameSize()

	p.input.SetWidth(innerWidth - t.Dialog.InputPrompt.GetHorizontalFrameSize() - 1)
	p.list.SetSize(innerWidth, height-heightOffset)
	p.help.SetWidth(innerWidth)

	rc := NewRenderContext(t, width)

	// Title
	totalCount := len(p.plugins)
	rc.Title = "Plugins (MCP Servers)"
	rc.TitleInfo = t.HalfMuted.Render(strings.Join([]string{"", string(rune(countDigits(totalCount))), " plugins"}, ""))

	// Category tab bar
	categoryBar := p.categoryTabBar(innerWidth)
	rc.AddPart(categoryBar)

	// Filter input
	inputView := t.Dialog.InputPrompt.Render(p.input.View())
	rc.AddPart(inputView)

	// List
	visibleCount := len(p.list.FilteredItems())
	if p.list.Height() >= visibleCount {
		p.list.ScrollToTop()
	} else {
		p.list.ScrollToSelected()
	}
	listView := t.Dialog.List.Height(p.list.Height()).Render(p.list.Render())
	rc.AddPart(listView)

	rc.Help = p.help.View(p)

	view := rc.Render()

	cur := p.Cursor()
	DrawCenterCursor(scr, area, view, cur)
	return cur
}

// ShortHelp implements help.KeyMap.
func (p *PluginsLibrary) ShortHelp() []key.Binding {
	return []key.Binding{
		p.keyMap.UpDown,
		p.keyMap.CategoryNext,
		p.keyMap.Select,
		p.keyMap.Close,
	}
}

// FullHelp implements help.KeyMap.
func (p *PluginsLibrary) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{p.keyMap.Select, p.keyMap.Next, p.keyMap.Previous, p.keyMap.FocusFilter},
		{p.keyMap.CategoryNext, p.keyMap.CategoryPrev, p.keyMap.Connect, p.keyMap.Close},
	}
}

// Filter implements ListItem.
func (p *PluginsLibraryItem) Filter() string {
	return p.plugin.Name + " " + p.plugin.Description
}

// ID implements ListItem.
func (p *PluginsLibraryItem) ID() string {
	return p.plugin.Name
}

// SetFocused implements ListItem.
func (p *PluginsLibraryItem) SetFocused(focused bool) {
	if p.focused != focused {
		p.cache = nil
	}
	p.focused = focused
}

// SetMatch implements ListItem.
func (p *PluginsLibraryItem) SetMatch(match fuzzy.Match) {
	p.cache = nil
	p.m = match
}

// Render implements ListItem.
func (p *PluginsLibraryItem) Render(width int) string {
	styles := ListItemStyles{
		ItemBlurred:     p.t.Dialog.NormalItem,
		ItemFocused:     p.t.Dialog.SelectedItem,
		InfoTextBlurred: p.t.Subtle,
		InfoTextFocused: p.t.Base,
	}

	// Status indicator
	var statusIcon string
	switch p.plugin.Status {
	case PluginStatusConnected:
		statusIcon = p.t.Base.Foreground(p.t.GreenLight).Render("● ")
	case PluginStatusAvailable:
		statusIcon = p.t.HalfMuted.Render("○ ")
	case PluginStatusError:
		statusIcon = p.t.Base.Foreground(p.t.Error).Render("✗ ")
	}

	name := statusIcon + p.plugin.Name

	// Show tool count for connected plugins
	if p.focused && p.plugin.Status == PluginStatusConnected && p.plugin.ToolsCount > 0 {
		toolCount := fmt.Sprintf("%d tools", p.plugin.ToolsCount)
		badge := p.t.Subtle.Render(" [" + toolCount + "]")
		name = name + badge
	}

	return renderItem(styles, name, p.plugin.Description, p.focused, width, p.cache, &p.m)
}

// ActionSelectPlugin is a message indicating a plugin has been selected.
type ActionSelectPlugin struct {
	PluginName        string
	PluginDescription string
	PluginStatus      PluginStatus
}
