package gemini

import (
	"context"
	"curaitor/internal/config"
	"curaitor/internal/data"
	"curaitor/internal/fileops"
	"curaitor/internal/model"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"google.golang.org/genai"
)

func ParseFileWorker(cfg *config.Config, ctx context.Context, wg *sync.WaitGroup, courses *data.Courses, newFilesCh <-chan string, errCh chan<- error) {
	defer wg.Done()
	var genai_config = &genai.GenerateContentConfig{
		ResponseMIMEType: "application/json",
		ResponseSchema: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"course_code":  {Type: genai.TypeString},
				"course_title": {Type: genai.TypeString},
				"desc":         {Type: genai.TypeString},
				"file_type":    {Type: genai.TypeString},
				"title":        {Type: genai.TypeString},
			},
			PropertyOrdering: []string{"course_code", "course_title", "desc", "file_type", "title"},
		},
	}

	for {
		select {
		case file := <-newFilesCh:
			var dirTreeStr string
			if _, err := os.Stat(cfg.SchoolPath); err == nil {
				dirTree, err := GetDirTree(cfg.SchoolPath)
				if err != nil {
					errCh <- fmt.Errorf("failed to get directory tree: %w", err)
					continue
				}

				dirTreeStr = formatDirTree(dirTree)
				slog.Info("directory tree", slog.String("tree", dirTreeStr))
			} else {
				dirTreeStr = "/school (no subdirectories)"
			}

			coursesStr := courses.String()

			msg, err := prepMessage(parseFilePrompt(dirTreeStr, coursesStr), file)
			if err != nil {
				errCh <- fmt.Errorf("failed to prepare message for Gemini: %w", err)
				continue
			}

			res, err := sendMessage(cfg, ctx, msg, genai_config)
			if err != nil {
				errCh <- fmt.Errorf("failed to send message to Gemini: %w", err)
				continue
			}

			fileInfo := model.FileInfo{}
			var destPath string
			destFileName := filepath.Base(file)

			if err := json.Unmarshal([]byte(res), &fileInfo); err != nil {
				errCh <- fmt.Errorf("failed to parse Gemini response: %w", err)
				// TODO: move to unknown folder or error folder idk
				continue
			}

			// if file is a syllabus
			if fileInfo.Code != "" && fileInfo.Desc != "" {
				if !courses.Exists(fileInfo.Code) {
					courses.Add(model.CourseInfo{Code: fileInfo.Code, Desc: fileInfo.Desc, CourseTitle: fileInfo.CourseTitle})
					if err := courses.Save(); err != nil {
						errCh <- fmt.Errorf("failed to save courses: %w", err)
					}
					destPath = filepath.Join(cfg.SchoolPath, fileInfo.Code)
					destFileName = "syllabus" + filepath.Ext(file)
				}
			}

			// if file is not a syllabus
			if fileInfo.FileType != "" && fileInfo.Title != "" {
				if !courses.Exists(fileInfo.Code) {
					errCh <- fmt.Errorf("course code from file does not exist in courses.json: %s", fileInfo.Code)
					// TODO: move to unknown or error folder idk
					continue
				}
				// create course directory if it doesn't exist
				destPath = filepath.Join(cfg.SchoolPath, fileInfo.Code, fileInfo.FileType)
				destFileName = fileInfo.Title + filepath.Ext(file)
			}

			if err := os.MkdirAll(destPath, os.ModePerm); err != nil {
				errCh <- fmt.Errorf("failed to create directory %s: %w", destPath, err)
				continue
			}

			if err := fileops.MoveFile(file, destPath, destFileName); err != nil {
				errCh <- fmt.Errorf("failed to move file to %s: %w", destPath, err)
				continue
			}
			slog.Info("file moved", slog.String("file", file), slog.String("destination", destPath))

		case <-ctx.Done():
			slog.Info("worker done")
			return
		}
	}
}

func parseFilePrompt(dirTreeStr string, coursesStr string) string {

	return fmt.Sprintf(`
	You are given an academic file (PDF, document, etc.).
	Decide if this file is a syllabus or another academic file.

	Return strictly as a valid JSON object of form below:
	{
		"course_code": "CS370",   
		"course_title": "Computer Science Practicum"
		"desc": "Covers AI foundations including search, logic, and learning.",
		"file_type": "assignment",
		"title": "Homework3-Search-Algorithms"
	}

	If the file IS a syllabus, leave "file_type" and "title" empty and return:

	{
		"course_code": "CS370",
		"course_title": "Computer Science Practicum"
		"desc": "Covers AI foundations including search, logic, and learning.",
		"file_type": "",
		"title": ""
	}

	Rules for syllabus files:
	- Extract the main course code from the syllabus. The code should not exist on directory structure as you are creating a new course.
	- Must be a single alphanumeric identifier (A–Z, 0–9).
	- If multiple course codes are listed, choose the first/main one only.
	- Do not include dashes, slashes, or section numbers.
	- Extract a brief 2-3 sentence description of the course from the syllabus content for the "desc" field.

	If the file is NOT a syllabus, leave "desc" empty and return:
	{
		"course_code": "CS370",
		"course_title": ""
		"desc": "",
		"file_type": "assignment",
		"title": "Homework3-Search-Algorithms"
	}

	Rules for non-syllabus files:
	- Check the provided directory structure for existing categories under the course.
	- If a matching folder already exists, use its category.
	- If no folder exists for this file_type, propose a new directory under the course code like "assignments", "exams", "lectures", "notes", etc.
	- Always generate a meaningful, descriptive title for the file.
	- Title will be used as new filename in the directory, so return a valid filename

	Current directory structure for context:
	%s

	Registered courses and descriptions: 
	%s
	`, dirTreeStr, coursesStr)
}

// Convert the directory tree map to a readable string format
func formatDirTree(dirTree map[string][]string) string {
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
