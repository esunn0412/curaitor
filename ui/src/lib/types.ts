export type Question = {
  question: string;
  choices: string[];
  answer: number;
  selected?: number;
  isCorrect?: boolean;
};

export type Course = {
  course_title: string;
  course_code: string;
  desc: string;
  // numFiles: number;
  // quiz: Question[];
};

export type Quiz = {
  course_code: string;
  questions: Question[];
};

export type CourseFile = {
  file_path: string;
  content: string;
};
