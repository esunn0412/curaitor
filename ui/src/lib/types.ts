export type Question = {
  question: string;
  choices: string[];
  answer: number;
  selected?: number;
  isCorrect?: boolean;
};

export type Course = {
  name: string;
  code: string;
  numFiles: number;
  quiz: Question[];
};
