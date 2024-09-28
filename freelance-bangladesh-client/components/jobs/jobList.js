"use client";
import React from "react";
import JobCard from "./jobCard";

const JobList = ({ jobs }) => {
  return (
    <div className="container mx-auto px-4">
        <h1 className="text-3xl font-bold text-center my-8">Available Jobs</h1>
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
            {jobs.map((job) => (
            <JobCard key={job.id} job={job} />
            ))}
        </div>
    </div>
  );
}

export default JobList;
