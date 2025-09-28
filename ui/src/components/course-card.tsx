"use client";

import { Course } from "@/lib/types";
import Link from "next/link";
import { FileIcon, LucideIcon, NotebookPenIcon } from "lucide-react";
import useQuizzes from "@/hooks/use-quizzes";
import useFiles from "@/hooks/use-files";

type CourseCardProps = {
  course: Course;
};

const CourseCard = ({ course }: CourseCardProps) => {
  const { quizzes } = useQuizzes();
  const { files } = useFiles();

  return (
    <Link
      href={course.course_code}
      className="rounded-xl border h-fit hover:opacity-80 hover:shadow-md transition-all"
    >
      <div className="bg-overlay rounded-t-xl p-4">
        <p className="font-bold mr-3 text-2xl uppercase">
          {course.course_code}
        </p>
        <p className="text-secondary text-lg truncate text-nowrap">
          {course.course_title}
        </p>
      </div>
      <div className="flex gap-4 p-2">
        <Tag
          value={
            files?.filter((f) => f.file_path.includes(course.course_code))
              .length
          }
          icon={FileIcon}
        />
        <Tag
          value={
            quizzes.find((q) => q.course_code === course.course_code)?.questions
              .length
          }
          icon={NotebookPenIcon}
        />
      </div>
    </Link>
  );
};

export default CourseCard;

type TagProps = {
  icon: LucideIcon;
  value?: string | number;
};
const Tag = (props: TagProps) => {
  return (
    <div className="flex items-center gap-1">
      <props.icon className="text-secondary size-4" />
      <span className="text-secondary">{props.value}</span>
    </div>
  );
};
