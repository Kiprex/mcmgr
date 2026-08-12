package project

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func dirExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func Init(projectPath string) error {
	var err error
	if len(projectPath) == 0 || strings.TrimSpace(projectPath) == "" {
		projectPath, err = os.Getwd()
		if err != nil {
			return err
		}
	}

	rootPath := filepath.Join(projectPath, ".mcmgr")

	if dirExists(rootPath) {
		return errors.New("Проект уже инициализирован")
	}

	// создаем корневую папку .mcmgr
	err = os.Mkdir(rootPath, 0755)

	if err != nil {
		return err
	}

	// создаем папку листов lists
	err = os.Mkdir(filepath.Join(rootPath, "lists"), 0755)

	if err != nil {
		os.RemoveAll(rootPath)
		return err
	}

	defaultJson := Config{
		MCVersion:   "",
		Core:        "",
		CoreVersion: "",
	}

	byteValue, err := json.MarshalIndent(defaultJson, "", "\t")

	if err != nil {
		os.RemoveAll(rootPath)
		return err
	}

	// создаем файл config.json
	configPath := filepath.Join(projectPath, ".mcmgr", "config.json")

	err = os.WriteFile(configPath, byteValue, 0644)
	if err != nil {
		os.RemoveAll(rootPath)
		return err
	}

	return nil
}
