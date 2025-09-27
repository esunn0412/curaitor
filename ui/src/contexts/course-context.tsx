"use client";

import { Course } from "@/lib/types";
import {
  createContext,
  Dispatch,
  ReactNode,
  SetStateAction,
  useEffect,
  useState,
} from "react";

type CourseContextType = {
  courses: Course[];
  setCourses: Dispatch<SetStateAction<Course[]>>;
};

export const CourseContext = createContext<CourseContextType | null>(null);

type CourseContextProviderType = {
  children: Readonly<ReactNode>;
};

export const CourseContextProvider = ({
  children,
}: CourseContextProviderType) => {
  const [courseData, setCourseData] = useState<Course[]>([]);

  useEffect(() => {
    const fetchCourses = async () => {
      const res = await fetch("http://localhost:9000/courses");
      setCourseData((await res.json()) as Course[]);
    };
    void fetchCourses();
  }, []);

  return (
    <CourseContext.Provider
      value={{ courses: courseData, setCourses: setCourseData }}
    >
      {children}
    </CourseContext.Provider>
  );
};
