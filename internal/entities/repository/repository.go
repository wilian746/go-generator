package repository

import (
	enumsRepository "github.com/wilian746/go-generator/internal/enums/repository"
	enumsCommands "github.com/wilian746/go-generator/internal/enums/repository/commands"
)

type Repository struct {
	Name     enumsRepository.Repository
	Commands []enumsCommands.Command
}

func (r Repository) GetListCommands() (existingCommands []string) {
	for _, value := range r.Commands {
		existingCommands = append(existingCommands, value.String())
	}
	return existingCommands
}
