"use client";

import QuizQuestion from "@/components/quiz-question";
import useCourses from "@/hooks/use-courses";
import { use } from "react";

type QuizPageProps = {
  params: Promise<{ course: string }>;
};

const QuizPage = ({ params }: QuizPageProps) => {
  const { course } = use(params);
  const { data } = useCourses();

  return (
    <main >
      <h1 className="text-4xl font-bold uppercase mb-10">
        <span className="text-secondary">{course} / </span>
        Quiz
      </h1>

      <ol className="space-y-10">
        {data
          .find((c) => c.code === course)
          ?.quiz.map((q, i) => (
            <QuizQuestion
              key={i}
              courseCode={course}
              index={i}
              question={q}
            />
          ))}
      </ol>
    </main>
  );
};

export default QuizPage;
