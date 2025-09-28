"use client";

import { EdgeContext } from "@/contexts/edge-context";
import { useContext } from "react";

const useEdges = () => {
  const edges = useContext(EdgeContext);
  if (edges === undefined) {
    throw new Error("EdgeContext must be used within EdgeContextProvider");
  }

  return edges;
};

export default useEdges;
