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

### Backend

The backend is written in Go and is responsible for the following:

- **File Watching:** It watches for changes in your course material directories and automatically processes new or updated files.
- **AI Integration:** It uses the Google Gemini API to generate quizzes and study guides.
- **API:** It provides a RESTful API for the frontend to consume.

### Frontend

The frontend is a Next.js application that provides a user-friendly interface for interacting with the application. It uses the following technologies:

- **Next.js:** A React framework for building server-rendered applications.
- **Ant Design:** A UI library for React.
- **React Markdown:** A component for rendering Markdown content.
- **Vis Network:** A library for creating interactive network graphs.