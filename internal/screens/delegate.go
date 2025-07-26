package screens

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type MenuItem string

func (i MenuItem) FilterValue() string { return "" }

type MenuDelegate struct {
	Styles *MenuStyles
}

func (d MenuDelegate) Height() int {
	return 1
}

func (d MenuDelegate) Spacing() int {
	return 0
}

func (d MenuDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd {
	return nil

}

func (d MenuDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(MenuItem)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)
	fn := d.Styles.ItemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return d.Styles.SelectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}
	fmt.Fprint(w, fn(str))
}
