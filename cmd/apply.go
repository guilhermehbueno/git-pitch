/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var forceOverwrite bool

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Create a default configuration file in the .git directory",
	Long: `Creates a default git-pitch configuration file in the .git directory.
	
This command will:
- Check if the current directory is a git repository
- Create a default configuration file in .git/
- Ask for confirmation if the file already exists`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := runApplyCommand(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

const defaultConfigContent = `# git-pitch configuration file
# This file contains default settings for git-pitch

# Default configuration settings
default:
  enabled: true
  auto_commit: false
  template: "default"

# Template settings
templates:
  default:
    title: "Git Pitch Template"
    description: "Default template for git-pitch presentations"
`

type confirmModel struct {
	question string
	answer   bool
	confirm  bool
	styles   confirmStyles
}

type confirmStyles struct {
	question lipgloss.Style
	option   lipgloss.Style
	selected lipgloss.Style
}

func newConfirmModel(question string) confirmModel {
	return confirmModel{
		question: question,
		answer:   false,
		confirm:  false,
		styles: confirmStyles{
			question: lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12")),
			option:   lipgloss.NewStyle().Foreground(lipgloss.Color("8")),
			selected: lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("2")),
		},
	}
}

func (m confirmModel) Init() tea.Cmd {
	return nil
}

func (m confirmModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "left", "h":
			m.answer = false
		case "right", "l":
			m.answer = true
		case "y", "Y":
			m.answer = true
			m.confirm = true
			return m, tea.Quit
		case "n", "N":
			m.answer = false
			m.confirm = true
			return m, tea.Quit
		case "enter":
			m.confirm = true
			return m, tea.Quit
		case "ctrl+c", "esc", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m confirmModel) View() string {
	var yesStyle, noStyle lipgloss.Style
	
	if m.answer {
		yesStyle = m.styles.selected
		noStyle = m.styles.option
	} else {
		yesStyle = m.styles.option
		noStyle = m.styles.selected
	}

	question := m.styles.question.Render(m.question)
	yes := yesStyle.Render("Yes")
	no := noStyle.Render("No")

	return fmt.Sprintf("%s\n\n[%s] / [%s]\n\n%s",
		question,
		yes,
		no,
		lipgloss.NewStyle().Faint(true).Render("Use arrow keys or y/n to choose, Enter to confirm, q to quit"),
	)
}

func askForConfirmation(question string) (bool, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%s [y/N]: ", question)
	
	if scanner.Scan() {
		response := strings.ToLower(strings.TrimSpace(scanner.Text()))
		return response == "y" || response == "yes", nil
	}
	
	if err := scanner.Err(); err != nil {
		return false, err
	}
	
	return false, nil
}

func runApplyCommand() error {
	gitDir := ".git"
	
	if _, err := os.Stat(gitDir); os.IsNotExist(err) {
		return fmt.Errorf("not a git repository. Please run 'git init' first")
	}
	
	configPath := filepath.Join(gitDir, "git-pitch-config.yaml")
	
	if _, err := os.Stat(configPath); err == nil && !forceOverwrite {
		// Try to use interactive confirmation first
		if isTerminal() {
			model := newConfirmModel("Configuration file already exists. Do you want to overwrite it?")
			
			p := tea.NewProgram(model)
			finalModel, err := p.Run()
			if err != nil {
				// Fall back to simple text prompt if bubbletea fails
				confirmed, fallbackErr := askForConfirmation("Configuration file already exists. Do you want to overwrite it?")
				if fallbackErr != nil {
					return fmt.Errorf("failed to get confirmation: %v", fallbackErr)
				}
				if !confirmed {
					fmt.Println("Operation cancelled.")
					return nil
				}
			} else {
				result := finalModel.(confirmModel)
				if !result.confirm || !result.answer {
					fmt.Println("Operation cancelled.")
					return nil
				}
			}
		} else {
			// Non-interactive mode - use simple prompt
			confirmed, err := askForConfirmation("Configuration file already exists. Do you want to overwrite it?")
			if err != nil {
				return fmt.Errorf("failed to get confirmation: %v", err)
			}
			if !confirmed {
				fmt.Println("Operation cancelled.")
				return nil
			}
		}
	}
	
	err := os.WriteFile(configPath, []byte(defaultConfigContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to create configuration file: %v", err)
	}
	
	fmt.Printf("✓ Configuration file created at %s\n", configPath)
	return nil
}

func isTerminal() bool {
	fileInfo, _ := os.Stdin.Stat()
	return (fileInfo.Mode() & os.ModeCharDevice) != 0
}

func init() {
	rootCmd.AddCommand(applyCmd)
	applyCmd.Flags().BoolVarP(&forceOverwrite, "force", "f", false, "Force overwrite existing configuration file")
}
