"use client";

import CourseCard from "@/components/course-card";
import useCourses from "@/hooks/use-courses";

const Home = () => {
  const { courses: data } = useCourses();

  return (
    <main className="basis-2/3">
      <div className="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-4">
        {data?.map((course, i) => (
          <CourseCard key={i} course={course} />
        ))}
      </div>
    </main>
  );
};

export default Home;
