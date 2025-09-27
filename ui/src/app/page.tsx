import { data } from "@/lib/data";
import CourseCard from "@/component/course-card";

const Home = () => {
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
