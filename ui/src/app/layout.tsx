import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";

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
        className={`${geistSans.variable} ${geistMono.variable} antialiased bg-gradient-to-tl from-cyan-100 to-transparent`}
      >
        <div className="fixed top-0 left-0 w-full h-16 border-b flex items-center px-4">
          <span className="font-bold text-secondary mr-3">CURAITOR</span>
          <span className="font-mono text-tertiary">/Users/ethantlee/school</span>
        </div>
        {children}
      </body>
    </html>
  );
}
