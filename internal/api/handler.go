package api

import (
	"curaitor/internal/data"
	"curaitor/internal/model"
	"encoding/json"
	"net/http"
	"slices"
)

func GetCoursesHandler(courses *data.Courses) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		courses.Mu.Lock()
		bytes, err := json.Marshal(courses.Courses)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		courses.Mu.Unlock()

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Content-Type", "text/json")
		_, err = w.Write(bytes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func GetQuizHandler(quizzes *data.Quiz, newQuizCodesCh chan<- string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		course := r.URL.Query().Get("course")

		quizzes.Mu.Lock()
		quizExists := slices.ContainsFunc(quizzes.Quizzes, func(quiz model.QuizInfo) bool {
			return quiz.Code == course
		})
		if !quizExists {
			newQuizCodesCh <- course
		}

		var quiz model.QuizInfo
		for _, q := range quizzes.Quizzes {
			if q.Code == course {
				quiz = q
			}
		}

		bytes, err := json.Marshal(quiz)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		quizzes.Mu.Unlock()

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Content-Type", "text/json")
		_, err = w.Write(bytes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func RegenerateQuizHandler(newQuizCodesCh chan string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		course := r.URL.Query().Get("course")

		newQuizCodesCh <- course

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Content-Type", "text/json")
	}
}

func GetFilesHandler(caches *data.CachedFiles) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		caches.Mu.Lock()
		bytes, err := json.Marshal(caches.CachedFiles)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		caches.Mu.Unlock()

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Content-Type", "text/json")

		_, err = w.Write(bytes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
