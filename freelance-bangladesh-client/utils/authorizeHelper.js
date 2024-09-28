import { useEffect } from "react";
import { getServerSession } from "next-auth";
import { useSession } from "next-auth/react";
import { useRouter } from "next/navigation";
import { authOptions } from "../app/api/auth/[...nextauth]/route";

const useCanActivateClient = () => {
    const { data: session, status } = useSession();
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

export const canActivateTalent = async () => {
    const session = await getServerSession(authOptions);

    if (session && 
        (session.roles?.includes("talent") 
        || session.roles?.includes("admin"))
    ) {
        return true;
    }

    return false;
}

export default useCanActivateClient;
