package repository

import (
	"fmt"
	entitiesRepository "github.com/wilian746/go-generator/internal/entities/repository"
	enumsRepository "github.com/wilian746/go-generator/internal/enums/repository"
	enumsCommands "github.com/wilian746/go-generator/internal/enums/repository/commands"
)

func setupRepositories() (repositories []entitiesRepository.Repository) {
	for _, repo := range enumsRepository.Values() {
		if repo == enumsRepository.Gorm {
			repositories = append(repositories, setupGormRepository())
		}
	}
	return repositories
}

func setupGormRepository() entitiesRepository.Repository {
	repository := entitiesRepository.Repository{Name: enumsRepository.Gorm}
	repository.Commands = append(repository.Commands, enumsCommands.App)
	return repository
}

func IsValidRepositoryAndCommand(repository, command string) bool {
	repositories := setupRepositories()
	for _, existingRepository := range repositories {
		if enumsRepository.ValueOf(repository) == existingRepository.Name {
			for _, existingCommand := range existingRepository.Commands {
				if enumsCommands.ValueOf(command) == existingCommand {
					return true
				}
			}
		}
	}
	return false
}

func GetCommandsValidByRepository(repository string) (validCommands []string) {
	repositories := setupRepositories()
	for _, existingRepository := range repositories {
		if enumsRepository.ValueOf(repository) == existingRepository.Name {
			validCommands = append(validCommands, existingRepository.GetListCommands()...)
		}
	}
	return validCommands
}

func GetAvailableCommands() (examples string) {
	repositories := setupRepositories()
	for _, existingRepository := range repositories {
		for _, existingCommand := range existingRepository.Commands {
			examples += fmt.Sprintf("go-generator init %s %s"+
				"\n        ",
				existingRepository.Name.String(), existingCommand.String())
		}
	}
	return examples
}
