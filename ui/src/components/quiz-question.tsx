"use client";

import useCourses from "@/hooks/use-courses";
import { Question } from "@/lib/types";
import { CheckIcon, XIcon } from "lucide-react";
import { useState } from "react";

type QuestionProps = {
  courseCode: string;
  index: number;
  question: Question;
};

const QuizQuestion = ({ courseCode, index, question }: QuestionProps) => {
  const { data, setData } = useCourses();
  const [selected, setSelected] = useState(-1);

  const handleSelect = (option: number) => {
    if (!!question.selected) return;
    setSelected(option === selected ? -1 : option);
  };

  const handleCheck = () => {
    if (!!question.selected) return;
    const newData = structuredClone(data);
    const q = newData.find((c) => c.code === courseCode)?.quiz[index];
    if (!q) throw new Error("invalid course code");

    q.selected = selected;
    q.isCorrect = selected === q.answer;

    setData(newData);
  };

  return (
    <li className="list-none space-y-4">
      <h3 className="text-2xl font-semibold">
        {index + 1}. {question.question}
      </h3>

      <ol className="">
        {question.choices.map((c, j) => {
          const isGreen = question.isCorrect && question.answer === j;
          const isRed = question.isCorrect === false && question.selected === j;

          return (
            <li
              key={j}
              onClick={() => handleSelect(j)}
              className={`list-none text-lg rounded-xl p-2 flex justify-between items-center cursor-pointer hover:bg-overlay transition-all ${isGreen ? "bg-green-100" : isRed ? "bg-red-100" : !question.selected && selected === j ? "bg-overlay outline-2" : ""}`}
            >
              <span>
                {String.fromCharCode(j + 65)}. {c}
              </span>
              {isGreen ? (
                <span>
                  <CheckIcon className="text-green-500" />
                </span>
              ) : (
                isRed && (
                  <span>
                    <XIcon className="text-red-500" />
                  </span>
                )
              )}
            </li>
          );
        })}
      </ol>

      <button
        onClick={() => handleCheck()}
        className="bg-overlay rounded-xl px-4 py-3 text-lg font-semibold cursor-pointer hover:opacity-70 transition-all"
      >
        Check Answer
      </button>
    </li>
  );
};

export default QuizQuestion;
