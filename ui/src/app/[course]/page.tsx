import Link from "next/link";
import "./study-guide.css";
import Markdown from "react-markdown";
import remarkGfm from "remark-gfm";

type CoursePageProps = {
  params: Promise<{ course: string }>;
};

const CoursePage = async ({ params }: CoursePageProps) => {
  const { course } = await params;
  const quizLink = course + "/" + "quiz";
  return (
    <main>
      <h1 className="text-4xl font-bold uppercase mb-10">{course}</h1>
      <Link
        href={quizLink}
        className="h-30 block w-full border rounded-xl bg-overlay mb-10 p-4"
      >
        Quiz
      </Link>
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
