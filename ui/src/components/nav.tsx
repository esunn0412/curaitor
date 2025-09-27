"use client";

import config from "@/lib/config";
import { usePathname } from "next/navigation";

const Nav = () => {
  const pathname = usePathname();
  const path = config.BASE_PATH + pathname;
  return (
    <div className="fixed top-0 left-0 w-full h-16 border-b flex items-center px-4">
      <span className="font-bold text-secondary mr-3">CURAITOR</span>
      <span className="font-mono text-tertiary">{path}</span>
    </div>
  );
};

export default Nav;
