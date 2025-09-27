"use client";

import config from "@/lib/config";
import { usePathname } from "next/navigation";
import Image from "next/image";

const Nav = () => {
  const pathname = usePathname();
  const path = config.BASE_PATH + pathname;
  return (
    <div className="fixed top-0 left-0 w-full h-16 border-b flex items-center gap-3 px-4">
      <div className="size-9 relative">
        <Image src="/images/icon.svg" fill alt="logo" />
      </div>
      <span className="font-mono text-tertiary">{path}</span>
    </div>
  );
};

export default Nav;
