"use client";

import { ChevronLeftIcon } from "lucide-react";
import { useRouter } from "next/navigation";

const BackButton = () => {
  const router = useRouter();
  return (
    <button
      onClick={router.back}
      className="cursor-pointer p-5 font-semibold text-secondary flex items-center gap-1"
    >
      <ChevronLeftIcon size={20} />
      Back
    </button>
  );
};

export default BackButton;
