package library

import (
	"encoding/json"
	"fmt"
	"os"
)

func (l *Library) SaveToFile(filename string) error {
	data, err := json.MarshalIndent(l.Books, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal error: %v", err)
	}

	return os.WriteFile(filename, data, 0644)
}

func (l *Library) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No existing library file found, starting fresh.")
			return nil
		}
		return err
	}

	return json.Unmarshal(data, &l.Books)
}