"use client";

import { data } from "@/lib/data";
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
  data: Course[];
  setData: Dispatch<SetStateAction<Course[]>>;
};

export const CourseContext = createContext<CourseContextType | null>(null);

type CourseContextProviderType = {
  children: Readonly<ReactNode>;
};

export const CourseContextProvider = ({
  children,
}: CourseContextProviderType) => {
  const [courseData, setCourseData] = useState(data);
  const [isLoaded, setIsLoaded] = useState(false);

  useEffect(() => {
    console.log("loading courses");
    const localData = localStorage.getItem("courses");
    if (localData) setCourseData(JSON.parse(localData));
    setIsLoaded(true);
  }, []);

  useEffect(() => {
    if (!isLoaded) return;
    console.log("saving courses");
    localStorage.setItem("courses", JSON.stringify(courseData));
  }, [isLoaded, courseData]);

  return (
    <CourseContext.Provider
      value={{ data: courseData, setData: setCourseData}}
    >
      {children}
    </CourseContext.Provider>
  );
};
