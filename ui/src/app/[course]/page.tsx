type CoursePageProps = {
  params: Promise<{ course: string }>;
};

const CoursePage = async ({ params }: CoursePageProps) => {
  const { course } = await params;
  return (
    <main>
      <h1 className="text-4xl font-bold uppercase">{course}</h1>
      <div>
        
        </div>
    </main>
  );
};

export default CoursePage;
