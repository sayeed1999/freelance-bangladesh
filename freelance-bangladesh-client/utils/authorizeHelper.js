import { useEffect } from "react";
import { useSession } from "next-auth/react";
import { useRouter } from "next/navigation";
import { authOptions } from "@/app/api/auth/[...nextauth]/route";

const useCanActivateClient = () => {
  const { data: session, status } = useSession(authOptions);
  const router = useRouter();

  useEffect(() => {
    if (
      status == "unauthenticated" ||
      (status == "authenticated" && 
      !(session.roles?.includes("admin") || session.roles?.includes("client")))
    ) {
      router.push("/unauthorized");
      router.refresh();
    }
  }, [session, status, router]);
}

const useCanActivateTalent = () => {
  const { data: session, status } = useSession(authOptions);
  const router = useRouter();

  useEffect(() => {
    if (
      status == "unauthenticated" ||
      (status == "authenticated" && 
      !(session.roles?.includes("admin") || session.roles?.includes("talent")))
    ) {
      router.push("/unauthorized");
      router.refresh();
    }
  }, [session, status, router]);
}

export {
  useCanActivateClient,
  useCanActivateTalent,
};