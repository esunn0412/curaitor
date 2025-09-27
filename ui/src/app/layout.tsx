import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";
import Nav from "@/components/nav";
import { CourseContextProvider } from "@/contexts/course-context";
import FileGraph from "@/components/file-graph";

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
        className={`${geistSans.variable} ${geistMono.variable} antialiased bg-cyan-50`}
      >
        <CourseContextProvider>
          <Nav />
          <div className="flex">
            {children}
            <FileGraph />
          </div>
        </CourseContextProvider>
      </body>
    </html>
  );
}
