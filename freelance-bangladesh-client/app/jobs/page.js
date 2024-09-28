"use client";

import { redirect, useRouter } from "next/navigation";
import { SetDynamicRoute } from "@/utils/setDynamicRoute";
import { useCanActivateTalent } from "@/utils/authorizeHelper";
import { getAllJobs } from "@/services/jobService";
import { useEffect, useState } from "react";
import { mockJobs } from "@/mock_data/mockJobs";
import JobCard from "@/components/jobs/jobCard";
import { useSession } from "next-auth/react";

export default async function Jobs() {
  const { data: session, status } = useSession();
  const [jobs, setJobs] = useState(mockJobs);
  const router = useRouter();

  useCanActivateTalent();

  // useEffect(() => {
  //   getAllJobs()
  //     .then((res) => setJobs(res))
  //     .catch((err) => alert(err));
  // }, []);

  if (status == "loading") {
    return (
      <main>
        <h1 className="text-4xl text-center">See jobs</h1>
        <div className="text-center text-2xl">Loading...</div>
      </main>
    );
  }

  return (        
    <main>  
      <SetDynamicRoute></SetDynamicRoute>    
      <div className="container mx-auto px-4">
        <h1 className="text-3xl font-bold text-center my-8">Available Jobs</h1>
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
          {jobs.map((job) => (
            <JobCard key={job.id} job={job} />
          ))}
        </div>
      </div>
    </main>
  );
}