//go:build unix
// +build unix

package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	filename     string
	size         int64
	created      string
	lastModified string
	owner        string
	permission   string
	description  string
	path         string
	error        string
	ready        bool
}

var (
	fileNameStyle       = lipgloss.NewStyle().Bold(true).Padding(0, 1).MarginBottom(1)
	otherItemsStyle     = lipgloss.NewStyle().Padding(0, 1)
	descriptionStyle    = lipgloss.NewStyle().Padding(0, 1)
	permStyle           = lipgloss.NewStyle().Bold(true)
	leftBorderStyle     = lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false, true, false, false)
	rightContainerStyle = lipgloss.NewStyle().MarginLeft(2)
	mainContainerStyle  = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1)
	controlsStyle       = lipgloss.NewStyle().Faint(true).MarginTop(1)
)

func initialModel(path string) model {
	m := model{path: path}
	info, err := os.Stat(path)
	if err != nil {
		m.error = fmt.Sprintf("Error accessing %s: %v", path, err)
		return m
	}
	if !info.Mode().IsRegular() {
		m.error = fmt.Sprintf("Error, %s is not regular file", path)
		return m
	}
	m.populateFromInfo(info)
	m.getDescription(path)
	m.ready = true
	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl-c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	if !m.ready {
		return mainContainerStyle.Render(
			leftBorderStyle.Render(m.error),
		) + "\n" + controlsStyle.Render("Press any key to exit")
	}

	infoLines := []string{
		fileNameStyle.Render(m.filename),
		otherItemsStyle.Render(fmt.Sprintf("Path: %s", m.path)),
		otherItemsStyle.Render(fmt.Sprintf("Size: %d bytes", m.size)),
		otherItemsStyle.Render(fmt.Sprintf("Created: %s", m.created)),
		otherItemsStyle.Render(fmt.Sprintf("Modified: %s", m.lastModified)),
		otherItemsStyle.Render(fmt.Sprintf("Owner: %s", m.owner)),
		otherItemsStyle.Render(fmt.Sprintf("Perms: %s", permStyle.Render(m.permission))),
	}
	left := lipgloss.JoinVertical(lipgloss.Left, infoLines...)

	right := lipgloss.JoinVertical(lipgloss.Left,
		lipgloss.NewStyle().Bold(true).Render("Description:"),
		descriptionStyle.Render(m.description),
	)

	content := lipgloss.JoinHorizontal(lipgloss.Top,
		leftBorderStyle.Render(left),
		rightContainerStyle.Render(right),
	)

	return mainContainerStyle.Render(content) + "\n" + controlsStyle.Render("(Press q to quit)")
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: specula [path]")
		os.Exit(1)
	}
	path := os.Args[1]

	m := initialModel(path)

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
}
