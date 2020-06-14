package prompt

import (
	"github.com/manifoldco/promptui"
)

type Interface interface {
	Ask(label, defaultValue string) (string, error)
	Select(label string, items []string) (string, error)
}

type Prompt struct {
	prompt    *promptui.Prompt
	selection *promptui.Select
}

func NewPrompt() Interface {
	return &Prompt{
		prompt: &promptui.Prompt{
			AllowEdit: true,
			Pointer:   promptui.PipeCursor,
		},
		selection: &promptui.Select{},
	}
}

func (p *Prompt) Ask(label, defaultValue string) (string, error) {
	p.prompt.Label = label
	p.prompt.Default = defaultValue
	return p.prompt.Run()
}

func (p *Prompt) Select(label string, items []string) (string, error) {
	p.selection.Label = label
	p.selection.Items = items
	_, response, err := p.selection.Run()
	return response, err
}
