import { syncUser } from "@/services/userService";
import { getServerSession } from "next-auth";
import { authOptions } from "./api/auth/[...nextauth]/NextAuthOptions";

export default async function Home() {
  const session = await getServerSession(authOptions);

  try {
    if (session?.user?.email) {
      await syncUser();
    }

    return (
      <div className="grid grid-rows-[20px_1fr_20px] items-center justify-items-center min-h-screen p-8 pb-20 gap-16 sm:p-20 font-[family-name:var(--font-geist-sans)]">
        <main className="flex flex-col gap-8 row-start-2 items-center sm:items-start">
          <h1 className="text-4xl">Freelance Bangladesh</h1>
          <h3 className="text-xl">Welcome to freelance bangladesh!</h3>
        </main>
      </div>
    );
  } catch (err) {
    return <h2>Failed to load homepage. Please refresh the page!</h2>;
  }
}
