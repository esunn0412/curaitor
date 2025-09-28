"use client";

import QuizQuestion from "@/components/quiz-question";
import useQuizzes from "@/hooks/use-quizzes";
import { Loader2Icon, RefreshCwIcon } from "lucide-react";
import { use, useEffect, useState } from "react";

type QuizPageProps = {
  params: Promise<{ course: string }>;
};

const QuizPage = ({ params }: QuizPageProps) => {
  const { course } = use(params);
  const { quizzes } = useQuizzes();
  const [regenerateTriggered, setRegenerateTriggered] = useState(false);
  const localQuizKey = `quiz-${course}`;

  const handleRegenerate = async () => {
    setRegenerateTriggered(true);
    localStorage.removeItem(localQuizKey);
    await fetch(`http://localhost:9000/quiz/regenerate?course=${course}`);
  };

  useEffect(() => {
    const localData = localStorage.getItem(localQuizKey);
    if (localData) {
      setRegenerateTriggered(false);
    }
  }, [localQuizKey]);

  return (
    <main className="h-full">
      <div className="flex justify-between">
        <h1 className="text-4xl font-bold uppercase mb-10">
          <span className="text-secondary">{course} / </span>
          Quiz
        </h1>

        <button
          onClick={handleRegenerate}
          disabled={regenerateTriggered}
          className="px-4 py-3 bg-overlay rounded-xl flex items-center gap-2 h-fit hover:opacity-80 transition-all cursor-pointer"
        >
          <RefreshCwIcon className="size-5 text-secondary" />
          Regenerate
        </button>
      </div>

      {regenerateTriggered ? (
        <div className="flex flex-col gap-2 items-center justify-center">
          Generating ...
          <Loader2Icon className="animate-spin" />
        </div>
      ) : (
        <ol className="space-y-10 pb-10">
          {quizzes
            .find((c) => c.course_code === course)
            ?.questions.map((q, i) => (
              <QuizQuestion
                key={i}
                courseCode={course}
                index={i}
                question={q}
              />
            ))}
        </ol>
      )}
    </main>
  );
};

export default QuizPage;
