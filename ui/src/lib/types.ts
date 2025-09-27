export type Question = {
  question: string;
  choices: string[];
  answer: number;
};

export type Course = {
  name: string;
  code: string;
  numFiles: number;
  quiz: Question[];
};
