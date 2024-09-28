"use client";
import React from "react";

const JobCard = ({ job }) => {
  return (
    <div className="max-w-sm rounded overflow-hidden shadow-lg bg-white p-6 hover:shadow-2xl transition-shadow duration-300">
      <div className="font-bold text-xl mb-2">{job.title}</div>
      <p className="text-gray-700 text-base">{job.description}</p>
      <div className="text-gray-800 font-semibold mt-4">Budget: ${job.budget}</div>
      {job.deadline && (
        <div className="text-gray-600 mt-2">Deadline: {new Date(job.deadline).toLocaleDateString()}</div>
      )}
    </div>
  );
};

export default JobCard;
