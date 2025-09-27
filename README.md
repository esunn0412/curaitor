# Curaitor

Curaitor is a web-based application that helps you study more effectively. It takes your course materials, organizes them, and automatically generates quizzes to help you test your knowledge.

## Why Curaitor?

In today's fast-paced learning environment, it's easy to get overwhelmed with the amount of information you need to process. Curaitor helps you stay on top of your studies by:

- **Automating quiz generation:** Spend less time creating study materials and more time learning.
- **Organizing your course materials:** Keep all your notes, lectures, and readings in one place.
- **Visualizing your knowledge:** See how different concepts connect with an interactive file graph.

## Features

- **Automatic Quiz Generation:** Curaitor uses AI to automatically generate quizzes from your course materials.
- **Course Management:** Easily upload and organize your course materials.
- **Interactive File Graph:** Visualize the relationships between your files and concepts.
- **Study Guide Generation:** Generate study guides from your course materials.

## Architecture

Curaitor is a monorepo with a Go backend and a Next.js frontend.

### Backend (Go)

The backend is responsible for:  

#### Configuration (`internal/config`)
- Loads all program configuration (e.g., Gemini API key, dump path) via environment variables.  
#### File Operations (`internal/fileops`)
- Directory watcher that monitors new/updated files.  
- File move and organization logic.  
- Directory tree builder (two-level deep). Example:  
```
CS101/ assignments/, homework/, exams/
CS270/ notes/, labs/
CS370/ (no subdirectories)
```
#### Gemini Integration (`internal/gemini`)
- `prepMessage`: Prepares prompts by embedding file data.  
- `sendMessage`: Sends requests to Gemini and returns responses.  
- Handles structured extraction of file metadata (`FileInfo` model).  
#### Application Data (`internal/data`)
- Stores courses, quizzes, study guides.  
- Writes updates to local JSON files (`courses.json`, `quizzes.json`, `cache.json`).  
#### Caching System
- Maintains a `cache.json` file to avoid re-parsing files.  
- Tracks `CachedFile { FilePath, Content }` to ensure integrity in concurrent operations.  
#### Quiz Generation (`gemini/generateQuiz.go`)
- Workers listen for course signals via a channel.  
- Each worker generates quiz questions from files in a course.  
- Quiz data model:  
```go
type Question struct {
  Question string
  Choices  []string
  Answer   int
}
type QuizInfo struct {
  Name     string
  Code     string
  NumFiles int
  QaPairs  []Question
}
```
#### Study Guide Generation
- Generates study guides **per course** instead of globally.  
- Uses mutexes to avoid race conditions when multiple files are processed concurrently.  
- Study guide is regenerated when a new file is added to a course’s directory.  
---

### Frontend

The frontend is a Next.js application that provides a user-friendly interface for interacting with the application. It uses the following technologies:

- **Next.js:** A React framework for building server-rendered applications.
- **Ant Design:** A UI library for React.
- **React Markdown:** A component for rendering Markdown content.
- **Vis Network:** A library for creating interactive network graphs.

### Workflow
1. File Dumping:
   - Place course materials (syllabus, notes, assignments) into the dump_files_here directory.
2. File Watching:
   - A watcher checks the directory every 10 seconds and sends new file paths to a processing channel.
3. File Processing (Concurrent):
   - Workers extract metadata with Gemini:
      - Syllabus: Creates a new course entry in courses.json.
      - Other files: Assigns to the correct course and type; validates against existing course codes.
4. File Organization:
   - Files are moved to their respective directories:
      - destpath = schoolpath + coursecode + filetype
      - destfilename = (is syllabus? same name : fileinfo.Title)
5. Caching:
   - Before parsing, file content is checked against cache.json to avoid duplicate API calls.
6. Quiz & Study Guide Generation:
   - Quizzes are generated per course on demand.
   - Study guides update whenever new course files are added.
