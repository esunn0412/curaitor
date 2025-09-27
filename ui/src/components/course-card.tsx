import { Course } from "@/lib/types";
import Link from "next/link";
import { FileIcon, LucideIcon, NotebookPenIcon } from "lucide-react";

type CourseCardProps = {
  course: Course;
};

const CourseCard = ({ course }: CourseCardProps) => {
  return (
    <Link
      href={course.code}
      className="rounded-xl border h-fit hover:opacity-80 transition-all"
    >
      <div className="bg-overlay rounded-t-xl p-4">
        <p className="font-bold mr-3 text-xl uppercase">{course.code}</p>
        <p className="text-secondary">{course.name}</p>
      </div>
      <div className="flex gap-4 p-2">
        <Tag value={course.numFiles.toString()} icon={FileIcon} />
        <Tag value={course.quiz.length.toString()} icon={NotebookPenIcon} />
      </div>
    </Link>
  );
};

export default CourseCard;

type TagProps = {
  icon: LucideIcon;
  value: string;
};
const Tag = (props: TagProps) => {
  return (
    <div className="flex items-center gap-1">
      <props.icon className="text-secondary size-4" />
      <span className="text-secondary text-sm">{props.value}</span>
    </div>
  );
};
