"use client";

import QuizQuestion from "@/components/quiz-question";
import useQuizzes from "@/hooks/use-quizzes";
import { use } from "react";

type QuizPageProps = {
  params: Promise<{ course: string }>;
};

const QuizPage = ({ params }: QuizPageProps) => {
  const { course } = use(params);
  const { quizzes } = useQuizzes();

  return (
    <main>
      <h1 className="text-4xl font-bold uppercase mb-10">
        <span className="text-secondary">{course} / </span>
        Quiz
      </h1>

      <ol className="space-y-10">
        {quizzes
          .find((c) => c.course_code === course)
          ?.questions.map((q, i) => (
            <QuizQuestion key={i} courseCode={course} index={i} question={q} />
          ))}
      </ol>
    </main>
  );
};

export default QuizPage;
