package fileops

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Convert the directory tree map to a readable string format
func FormatDirTree(dirTree map[string][]string) string {
	if len(dirTree) == 0 {
		return "No courses found under school directory."
	}

	var result strings.Builder

	// Sort course codes for consistent output
	var courses []string
	for course := range dirTree {
		courses = append(courses, course)
	}
	sort.Strings(courses)

	for _, course := range courses {
		categories := dirTree[course]
		result.WriteString(fmt.Sprintf("%s/", course))

		if len(categories) == 0 {
			result.WriteString(" (no subdirectories)")
		} else {
			for i, category := range categories {
				if i == 0 {
					result.WriteString(fmt.Sprintf(" %s/", category))
				} else {
					result.WriteString(fmt.Sprintf(", %s/", category))
				}
			}
		}
		result.WriteString("\n")
	}

	return result.String()
}

func GetDirTree(root string) (map[string][]string, error) {
	tree := make(map[string][]string)

	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		coursePath := filepath.Join(root, entry.Name())
		subentries, err := os.ReadDir(coursePath)
		if err != nil {
			return nil, err
		}

		var subdirs []string
		for _, subentry := range subentries {
			if subentry.IsDir() {
				subdirs = append(subdirs, subentry.Name())
			}
		}
		tree[entry.Name()] = subdirs
	}
	return tree, nil
}
