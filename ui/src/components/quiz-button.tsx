import Link from "next/link";
import { Progress, Tooltip } from "antd";
import {
  ChevronRightIcon,
  CircleCheckIcon,
  CircleDashedIcon,
  CircleXIcon,
  SparklesIcon,
} from "lucide-react";

type QuizButtonProps = {
  link: string;
  data: {
    correct: number;
    incorrect: number;
    todo: number;
  };
};

const QuizButton = ({ link, data }: QuizButtonProps) => {
  const total = data.correct + data.incorrect + data.todo;

  return (
    <div className="space-y-4">
      <Link
        href={link}
        className="flex group bg-overlay font-semibold text-2xl p-4 rounded-xl justify-between hover:opacity-80 transition-all"
      >
        <div className="flex flex-col justify-center gap-1">
          <div className="flex items-center gap-2">
            <SparklesIcon className="text-violet-500 size-5" />
            Take Quiz
            <ChevronRightIcon className="text-secondary opacity-0 -translate-x-2 group-hover:opacity-100 group-hover:translate-x-0 transition-all" />
          </div>
          <p className="text-base text-secondary flex gap-3 items-center [&_svg]:size-4 [&_svg]:inline [&_span]:flex [&_span]:items-center [&_span]:gap-1">
            <span className="text-green-500">
              <CircleCheckIcon /> {data.correct}
            </span>
            <span className="text-red-500">
              <CircleXIcon /> {data.incorrect}
            </span>
            <span>
              <CircleDashedIcon /> {data.todo}
            </span>
          </p>
        </div>
        <Tooltip
          title={`${data.correct} correct / ${data.incorrect} incorrect / ${data.todo} to do`}
        >
          <Progress
            percent={((data.correct + data.incorrect) / total) * 100}
            strokeLinecap="butt"
            success={{
              percent: (data.correct / total) * 100,
            }}
            format={(_, successPercent) => <span>{successPercent}%</span>}
            strokeColor="#f87171"
            type="circle"
            size={80}
          />
        </Tooltip>
      </Link>
    </div>
  );
};

export default QuizButton;
