"use client";

import { StudyGuideContext } from "@/contexts/study-guide-context";
import { useContext } from "react";

const useStudyGuides = () => {
  const studyGuides = useContext(StudyGuideContext);
  if (!studyGuides) {
    throw new Error("StudyGuideContext must be used within StudyGuideContextProvider");
  }

  return studyGuides;
};

export default useStudyGuides;
