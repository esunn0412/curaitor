import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";
import Nav from "@/components/nav";
import { CourseContextProvider } from "@/contexts/course-context";
import FileGraph from "@/components/file-graph";
import { QuizContextProvider } from "@/contexts/quiz-context";
import { FileContextProvider } from "@/contexts/file-context";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "Curaitor",
  description: "School made easier",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased h-svh`}
      >
        <div className="fixed -bottom-20 -right-20 size-200 rounded-full bg-cyan-100 blur-[200px] -z-40" />
        <CourseContextProvider>
          <QuizContextProvider>
            <FileContextProvider>
              <Nav />
              <div className="flex pt-16 h-full">
                {children}
                <FileGraph />
              </div>
            </FileContextProvider>
          </QuizContextProvider>
        </CourseContextProvider>
      </body>
    </html>
  );
}
