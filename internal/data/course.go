package data

import (
	"curaitor/internal/model"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"sync"
)

type Courses struct {
	Mu      *sync.Mutex
	Courses []model.CourseInfo
}

func LoadCourses() (*Courses, error) {
	courses, err := loadCourses()
	if err != nil {
		return nil, fmt.Errorf("failed to load courses: %w", err)
	}

	slog.Info("courses loaded")

	return &Courses{
		Mu:      &sync.Mutex{},
		Courses: courses,
	}, nil
}

func loadCourses() ([]model.CourseInfo, error) {
	var courses []model.CourseInfo

	data, err := os.ReadFile("courses.json")
	if err != nil {
		// return without error if courses.json doesn't exist
		if os.IsNotExist(err) {
			return courses, nil
		}
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	
	// handle empty json file
	if len(data) == 0 {
		return courses, nil
	}

	if err := json.Unmarshal(data, &courses); err != nil {
		return nil, fmt.Errorf("failed to unmarshal courses: %w", err)
	}

	return courses, nil
}

func (c *Courses) Add(course model.CourseInfo) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Courses = append(c.Courses, course)
	slog.Info("course added", slog.String("course", course.Code))
}

func (c *Courses) Save() error {
	c.Mu.Lock()
	data, err := json.MarshalIndent(c.Courses, "", "  ")
	c.Mu.Unlock()

	if err != nil {
		return fmt.Errorf("failed to marshal courses: %w", err)
	}

	if err := os.WriteFile("courses.json", data, 0644); err != nil {
		return fmt.Errorf("failed to write courses file: %w", err)
	}

	return nil
}

func (c *Courses) Exists(code string) bool {
	c.Mu.Lock() 
	defer c.Mu.Unlock() 
	for _, course := range c.Courses {
		if course.Code == code {
			return true
		}
	}
	return false
}

func (c *Courses) String() string {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	result := ""
	for _, course := range c.Courses {
		result += fmt.Sprintf("%s: %s\n", course.Code, course.Desc)
	}
	return result
}