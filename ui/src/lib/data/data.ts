export const data: Quiz[] = [
  {
    course: { name: "Analysis of Algorithms", code: "CS 326", files: 5 },
    questions: [
      {
        question: "Lorem ipsum",
        choices: ["Lorem ipsum", "Dolor sit", "amet", "asdf", "qwer"],
        answer: 1,
      },
      {
        question: "Lorem ipsum",
        choices: ["Lorem ipsum", "Dolor sit", "amet", "asdf", "qwer"],
        answer: 1,
      },
      {
        question: "Lorem ipsum",
        choices: ["Lorem ipsum", "Dolor sit", "amet", "asdf", "qwer"],
        answer: 1,
      },
      {
        question: "Lorem ipsum",
        choices: ["Lorem ipsum", "Dolor sit", "amet", "asdf", "qwer"],
        answer: 1,
      },
      {
        question: "Lorem ipsum",
        choices: ["Lorem ipsum", "Dolor sit", "amet", "asdf", "qwer"],
        answer: 1,
      },
      {
        question: "Lorem ipsum",
        choices: ["Lorem ipsum", "Dolor sit", "amet", "asdf", "qwer"],
        answer: 1,
      },
    ],
  },
  {
    course: { name: "Human Computer Interaction", code: "CS 485-2", files: 6 },
    questions: [
      {
        question: "Lorem ipsum",
        choices: ["Lorem ipsum", "Dolor sit", "amet", "asdf", "qwer"],
        answer: 1,
      },
      {
        question: "Lorem ipsum",
        choices: ["Lorem ipsum", "Dolor sit", "amet", "asdf", "qwer"],
        answer: 1,
      },
      {
        question: "Lorem ipsum",
        choices: ["Lorem ipsum", "Dolor sit", "amet", "asdf", "qwer"],
        answer: 1,
      },
      {
        question: "Lorem ipsum",
        choices: ["Lorem ipsum", "Dolor sit", "amet", "asdf", "qwer"],
        answer: 1,
      },
    ],
  },
  {
    course: { name: "Information Security", code: "CS 485-5", files: 13 },
    questions: [
      {
        question: "Lorem ipsum",
        choices: ["Lorem ipsum", "Dolor sit", "amet", "asdf", "qwer"],
        answer: 1,
      },
      {
        question: "Lorem ipsum",
        choices: ["Lorem ipsum", "Dolor sit", "amet", "asdf", "qwer"],
        answer: 1,
      },
      {
        question: "Lorem ipsum",
        choices: ["Lorem ipsum", "Dolor sit", "amet", "asdf", "qwer"],
        answer: 1,
      },
    ],
  },
  {
    course: { name: "Computer Science Practicum", code: "CS 370", files: 12 },
    questions: [
      {
        question: "Lorem ipsum",
        choices: ["Lorem ipsum", "Dolor sit", "amet", "asdf", "qwer"],
        answer: 1,
      },
      {
        question: "Lorem ipsum",
        choices: ["Lorem ipsum", "Dolor sit", "amet", "asdf", "qwer"],
        answer: 1,
      },
      {
        question: "Lorem ipsum",
        choices: ["Lorem ipsum", "Dolor sit", "amet", "asdf", "qwer"],
        answer: 1,
      },
      {
        question: "Lorem ipsum",
        choices: ["Lorem ipsum", "Dolor sit", "amet", "asdf", "qwer"],
        answer: 1,
      },
    ],
  },
];

type Question = {
  question: string;
  choices: string[];
  answer: number;
};

type Course = {
  name: string;
  code: string;
  files: number;
};

type Quiz = {
  course: Course;
  questions: Question[];
};


