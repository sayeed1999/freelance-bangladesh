import { Inter } from "next/font/google";
import "./globals.css";

const inter = Inter({subsets: ["latin"]});

export const metadata = {
  title: "Freelance Bangladesh",
  description: "Not written yet...",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <div className="flex flex-row">
          <div className="w-4/5 p-3 h-screen bg-white">{children}</div>
          <div className="w-1/5 p-3 h-screen bg-blue-300">
            <h2 className="text-3xl">Demo - frontend</h2>
              Auth
            <hr />
              Nav
          </div>
        </div>
      </body>
    </html>
  );
}
