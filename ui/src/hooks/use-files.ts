"use client";

import { FileContext } from "@/contexts/file-context";
import { useContext } from "react";

const useFiles = () => {
  const files = useContext(FileContext);
  if (!files) {
    throw new Error("FileContext must be used within FileContextProvider");
  }

  return files;
};

export default useFiles;
