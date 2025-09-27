"use client";

import useCourses from "@/hooks/use-courses";
import { createContext, ReactNode, useEffect, useState } from "react";

type StudyGuideContextType = {
  course: string;
  content: string;
}[];

export const StudyGuideContext = createContext<StudyGuideContextType | null>(
  null,
);

type StudyGuideContextProviderType = {
  children: Readonly<ReactNode>;
};

export const StudyGuideContextProvider = ({
  children,
}: StudyGuideContextProviderType) => {
  const [studyGuideData, setStudyGuideData] = useState<StudyGuideContextType>(
    [],
  );
  const { courses } = useCourses();

  useEffect(() => {
    const fetchStudyGuide = async (course: string) => {
      const res = await fetch(
        `http://localhost:9000/study-guide?course=${course}`,
      );
      return {
        course: course,
        content: await res.text(),
      };
    };

    const fetchAllStudyGuide = async () => {
      setStudyGuideData(
        await Promise.all(courses.map((c) => fetchStudyGuide(c.course_code))),
      );
    };

    void fetchAllStudyGuide();
  }, [courses]);

  return (
    <StudyGuideContext.Provider value={studyGuideData}>
      {children}
    </StudyGuideContext.Provider>
  );
};
