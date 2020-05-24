package prompt

import "github.com/manifoldco/promptui"

type Interface interface {
	Ask(label, defaultValue string) (string, error)
}

type Prompt struct {
}

func NewPrompt() Interface {
	return &Prompt{}
}

func (p *Prompt) Ask(label, defaultValue string) (string, error) {
	cmd := promptui.Prompt{
		Label:   label,
		Default: defaultValue,
	}
	return cmd.Run()
}
