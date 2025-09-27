"use client";

import { CourseContext } from "@/contexts/course-context";
import { useContext } from "react";

const useCourses = () => {
  const courses = useContext(CourseContext);
  if (!courses) {
    throw new Error("CourseContext must be used within CourseContextProvider");
  }

  return courses;
};

export default useCourses;
