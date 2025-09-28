"use client";

import useCourses from "@/hooks/use-courses";
import { Quiz } from "@/lib/types";
import {
  createContext,
  Dispatch,
  ReactNode,
  SetStateAction,
  useEffect,
  useState,
} from "react";

type QuizContextType = {
  quizzes: Quiz[];
  setQuizzes: Dispatch<SetStateAction<Quiz[]>>;
};

export const QuizContext = createContext<QuizContextType | null>(null);

type CourseContextProviderType = {
  children: Readonly<ReactNode>;
};

export const QuizContextProvider = ({
  children,
}: CourseContextProviderType) => {
  const { courses: coursesData } = useCourses();
  const [quizData, setQuizData] = useState<Quiz[]>([]);
  const [isLoaded, setIsLoaded] = useState(false);

  useEffect(() => {
    const fetchQuiz = async (course: string): Promise<Quiz> => {
      const localData = localStorage.getItem(`quiz-${course}`);
      if (localData) {
        return JSON.parse(localData) as Quiz;
      }

      const res = await fetch(`http://localhost:9000/quiz?course=${course}`);
      return (await res.json()) as Quiz;
    };

    const fetchQuizzes = async () => {
      if (!coursesData) return;
      try {
        const quizzes = await Promise.all(
          coursesData.map((course) => fetchQuiz(course.course_code)),
        );
        setQuizData(quizzes);
        setIsLoaded(true);
      } catch {}
    };

    void fetchQuizzes();
  }, [coursesData]);

  useEffect(() => {
    if (!isLoaded) return;
    quizData.forEach((q) => {
      localStorage.setItem(`quiz-${q.course_code}`, JSON.stringify(q));
    });
  }, [isLoaded, quizData]);

  return (
    <QuizContext.Provider
      value={{ quizzes: quizData, setQuizzes: setQuizData }}
    >
      {children}
    </QuizContext.Provider>
  );
};
