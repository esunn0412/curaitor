"use client";

import { createContext, ReactNode, useEffect, useState } from "react";

type EdgeContextType = {
  from: string;
  to: string;
}[];

export const EdgeContext = createContext<EdgeContextType | null>(null);

type EdgeContextProviderType = {
  children: Readonly<ReactNode>;
};

export const EdgeContextProvider = ({ children }: EdgeContextProviderType) => {
  const [edgesData, setEdgesData] = useState<EdgeContextType>([]);

  useEffect(() => {
    const fetchEdges = async () => {
      try {
        const res = await fetch("http://localhost:9000/edges");
        setEdgesData((await res.json()) as EdgeContextType);
      } catch {}
    };
    void fetchEdges();
  }, []);

  return (
    <EdgeContext.Provider value={edgesData}>{children}</EdgeContext.Provider>
  );
};
