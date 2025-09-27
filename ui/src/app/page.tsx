import { FileIcon, NotebookPenIcon } from "lucide-react";
import { data } from "@/lib/data";
import Link from "next/link";
import Tag from "@/component/tag";

export default function Home() {
  return (
    <main className="min-h-svh container mx-auto max-md:p-4 pt-40">
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
        {data.map((q, i) => (
          <Link
            key={i}
            href="#"
            className="rounded-xl border hover:opacity-80 transition-all"
          >
            <div className="bg-overlay rounded-t-xl p-4">
              <p className="font-bold mr-3 text-xl">{q.course.code}</p>
              <p className="text-secondary">{q.course.name}</p>
            </div>
            <div className="flex gap-4 p-2">
              <Tag value={q.course.files.toString()} icon={FileIcon} />
              <Tag
                value={q.questions.length.toString()}
                icon={NotebookPenIcon}
              />
            </div>
          </Link>
        ))}
      </div>
    </main>
  );
}


