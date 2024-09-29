import { SetDynamicRoute } from "@/utils/setDynamicRoute";
import { getAllJobs } from "@/services/jobService";
import JobList from "@/components/jobs/jobList";
import { getServerSession } from "next-auth";
import { authOptions } from "../api/auth/[...nextauth]/route";
import { redirect } from "next/navigation";

export default async function Jobs() {
  const session = await getServerSession(authOptions);

  if (!(session && (session.roles?.includes("talent") || session.roles?.includes("admin")))) {
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
    {console.error(err)}
    return (
      <main>
        <div className="text-center text-2xl">Failed to load jobs... Please refresh the page!</div>
      </main>
    )
  }
}