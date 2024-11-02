import { SetDynamicRoute } from "@/utils/setDynamicRoute";
import { getAllJobs } from "@/services/jobService";
import JobList from "@/components/jobs/jobList";
import { getServerSession } from "next-auth";
import { redirect } from "next/navigation";
import { authOptions } from "../api/auth/[...nextauth]/NextAuthOptions";

export default async function Jobs() {
  const session = await getServerSession(authOptions);

  if (!(session && 
    (session.roles?.includes("talent") || 
    session.roles?.includes("client") || 
    session.roles?.includes("admin")))) {
    redirect("/unauthorized");
  }

  try {
    var jobs = await getAllJobs();

    return (        
      <main>  
        <SetDynamicRoute></SetDynamicRoute>    
        <JobList jobs={jobs.result} total={jobs.total} />
      </main>
    );
  } catch (err) {
    return (
      <main>
        <div className="text-center text-2xl">Failed to load jobs... Please refresh the page!</div>
      </main>
    )
  }
}