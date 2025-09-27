"use client";

import { CourseFile } from "@/lib/types";
import {
  createContext,
  Dispatch,
  ReactNode,
  SetStateAction,
  useEffect,
  useState,
} from "react";

type FileContextType = {
  files: CourseFile[];
  setFiles: Dispatch<SetStateAction<CourseFile[]>>;
};

export const FileContext = createContext<FileContextType | null>(null);

type CourseContextProviderType = {
  children: Readonly<ReactNode>;
};

export const FileContextProvider = ({
  children,
}: CourseContextProviderType) => {
  const [filesData, setFilesData] = useState<CourseFile[]>([]);

  useEffect(() => {
    const fetchFiles = async () => {
      const res = await fetch("http://localhost:9000/files");
      setFilesData((await res.json()) as CourseFile[]);
    };
    void fetchFiles();
  }, []);

  return (
    <FileContext.Provider
      value={{ files: filesData, setFiles: setFilesData }}
    >
      {children}
    </FileContext.Provider>
  );
};
