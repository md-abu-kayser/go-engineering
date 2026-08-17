package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Lesson struct {
	ID        string `json:"id"`
	Level     string `json:"level"`
	Topic     string `json:"topic"`
	Path      string `json:"path"`
	Source    string `json:"source"`
	HasTests  bool   `json:"has_tests"`
	HasReadme bool   `json:"has_readme"`
}

func main() {
	lessons, err := discoverLessons(".")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(lessons); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func discoverLessons(root string) ([]Lesson, error) {
	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, fmt.Errorf("read repository root: %w", err)
	}

	var lessons []Lesson
	for _, levelEntry := range entries {
		if !levelEntry.IsDir() || !strings.HasPrefix(levelEntry.Name(), "level-") {
			continue
		}

		levelPath := filepath.Join(root, levelEntry.Name())
		lessonEntries, err := os.ReadDir(levelPath)
		if err != nil {
			return nil, fmt.Errorf("read %s: %w", levelEntry.Name(), err)
		}

		for _, lessonEntry := range lessonEntries {
			if !lessonEntry.IsDir() || !isLessonDirectory(lessonEntry.Name()) {
				continue
			}

			lessonPath := filepath.Join(levelPath, lessonEntry.Name())
			readmePath := filepath.Join(lessonPath, "README.md")
			mainPath := filepath.Join(lessonPath, "main.go")
			if _, err := os.Stat(mainPath); err != nil {
				continue
			}

			hasTests, err := hasTestFile(lessonPath)
			if err != nil {
				return nil, fmt.Errorf("inspect %s: %w", lessonPath, err)
			}

			_, readmeErr := os.Stat(readmePath)
			lessons = append(lessons, Lesson{
				ID:        lessonEntry.Name(),
				Level:     levelEntry.Name(),
				Topic:     humanizeLessonName(lessonEntry.Name()),
				Path:      filepath.ToSlash(filepath.Join(levelEntry.Name(), lessonEntry.Name())),
				Source:    filepath.ToSlash(filepath.Join(levelEntry.Name(), lessonEntry.Name(), "main.go")),
				HasTests:  hasTests,
				HasReadme: readmeErr == nil,
			})
		}
	}

	sort.Slice(lessons, func(i, j int) bool {
		if lessons[i].Level == lessons[j].Level {
			return lessons[i].Path < lessons[j].Path
		}
		return lessons[i].Level < lessons[j].Level
	})

	return lessons, nil
}

func hasTestFile(dir string) (bool, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false, err
	}

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), "_test.go") {
			return true, nil
		}
	}
	return false, nil
}

func isLessonDirectory(name string) bool {
	if len(name) < 4 {
		return false
	}
	return name[0] >= '0' && name[0] <= '9' && name[1] >= '0' && name[1] <= '9' && name[2] == '-'
}

func humanizeLessonName(name string) string {
	name = strings.TrimSpace(name)
	if len(name) > 3 && name[2] == '-' {
		name = name[3:]
	}
	words := strings.FieldsFunc(name, func(r rune) bool {
		return r == '-' || r == '_'
	})
	for i := range words {
		if words[i] == "" {
			continue
		}
		words[i] = strings.ToUpper(words[i][:1]) + words[i][1:]
	}
	return strings.Join(words, " ")
}
