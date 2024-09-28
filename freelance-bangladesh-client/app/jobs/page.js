import { SetDynamicRoute } from "@/utils/setDynamicRoute";
import { getAllJobs } from "@/services/jobService";
import JobList from "@/components/jobs/jobList";
import { redirect } from "next/dist/server/api-utils";
import { getServerSession } from "next-auth";
import { authOptions } from "../api/auth/[...nextauth]/route";

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
        <JobList jobs={jobs} />
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