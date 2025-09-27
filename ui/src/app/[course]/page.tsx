type CoursePageProps = {
  params: Promise<{ course: string }>;
};

const CoursePage = async ({ params }: CoursePageProps) => {
  const { course } = await params;
  return <main>{course}</main>;
};

export default CoursePage;
