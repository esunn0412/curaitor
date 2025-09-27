"use client";

import { QuizContext } from "@/contexts/quiz-context";
import { useContext } from "react";

const useQuizzes = () => {
  const quizzes = useContext(QuizContext);
  if (!quizzes) {
    throw new Error("CourseContext must be used within CourseContextProvider");
  }

  return quizzes;
};

export default useQuizzes;
