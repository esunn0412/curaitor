"use client";

import CourseCard from "@/components/course-card";
import useCourses from "@/hooks/use-courses";

const Home = () => {
  const { data } = useCourses();

  return (
    <main>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
        {data.map((course, i) => (
          <CourseCard key={i} course={course} />
        ))}
      </div>
    </main>
  );
};

export default Home;
