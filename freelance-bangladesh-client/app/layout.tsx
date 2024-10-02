import { Inter } from "next/font/google";
import AuthStatus from "@/components/authStatus";
import Nav from "@/components/nav";
import SessionProviderWrapper from "@/utils/sessionProviderWrapper";
import SessionGuard from "@/utils/sessionGuard";
import "./globals.css";

const inter = Inter({ subsets: ["latin"] });

export const metadata = {
  title: "Freelance Bangladesh",
  description: "Not written yet...",
};

const mainContentArea = (children: any) => (
  <div className="flex-1 p-6 bg-gray-100 shadow-lg rounded-tr-3xl rounded-br-3xl">
    {children}
  </div>
);

const sideBarArea = (
  <div className="w-1/5 p-6 h-screen bg-teal-400 text-gray-200 shadow-xl flex flex-col justify-between">
    <div>
      <h2 className="text-4xl font-semibold mb-6 text-gray-700">
        Freelance BD
      </h2>

      <AuthStatus />

      <hr className="my-4 border-t-2 border-gray-700 opacity-30" />

      <Nav />
    </div>

    <footer className="mt-6">
      <p className="text-sm text-gray-500">
        © {new Date().getFullYear()} Freelance BD
      </p>
    </footer>
  </div>
);

export default function RootLayout({ children }: { children: any }) {
  return (
    <SessionProviderWrapper>
      <SessionGuard>
        <html lang="en">
          <body className={inter.className}>
            <div className="flex flex-row min-h-screen">
              {mainContentArea(children)}
              {sideBarArea}
            </div>
          </body>
        </html>
      </SessionGuard>
    </SessionProviderWrapper>
  );
}
