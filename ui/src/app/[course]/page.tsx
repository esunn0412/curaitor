"use client";

import "./study-guide.css";
import Markdown from "react-markdown";
import remarkGfm from "remark-gfm";
import QuizButton from "@/components/quiz-button";
import useQuizzes from "@/hooks/use-quizzes";
import { use } from "react";
import useCourses from "@/hooks/use-courses";

type CoursePageProps = {
  params: Promise<{ course: string }>;
};

const CoursePage = ({ params }: CoursePageProps) => {
  const { course } = use(params);
  const { quizzes } = useQuizzes();
  const { courses } = useCourses();
  const quizLink = course + "/" + "quiz";
  const data = {
    correct: 0,
    incorrect: 0,
    todo: 0,
  };

  const quiz = quizzes.find((q) => q.course_code === course);

  quiz?.questions.forEach((question) => {
    if (question.selected === undefined) data.todo++;
    else if (question.isCorrect) data.correct++;
    else data.incorrect++;
  });

  return (
    <main>
      <h1 className="text-4xl font-bold uppercase mb-1">{course}</h1>
      <h2 className="text-xl text-secondary mb-8">
        {courses.find((c) => c.course_code === course)?.course_title}
      </h2>
      <QuizButton link={quizLink} data={data} />

      <StudyGuide />
    </main>
  );
};

export default CoursePage;

const markdown = `
## Heading

This is a sample markdown text.

### List

Here is a list of items:

-   Item 1
-   Item 2
-   Item 3

### Table

| Column 1 | Column 2 | Column 3 |
| --- | --- | --- |
| Cell 1 | Cell 2 | Cell 3 |
| Cell 4 | Cell 5 | Cell 6 |

 `;

const StudyGuide = () => {
  return (
    <article className="study-guide">
      <Markdown remarkPlugins={[remarkGfm]}>{markdown}</Markdown>
    </article>
  );
};
