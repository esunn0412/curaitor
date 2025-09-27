import BackButton from "@/components/back-button";
import { ReactNode } from "react";

type CourseLayoutProps = {
  children: Readonly<ReactNode>;
};

const CourseLayout = ({ children }: CourseLayoutProps) => {
  return (
    <div className="flex-1 overflow-y-scroll h-full">
      <BackButton />
      {children}
    </div>
  );
};

export default CourseLayout;
