"use client";
import React, { useState } from "react";
import { useRouter } from "next/navigation";
import BidJob from "./bidJob";
import { useSession } from "next-auth/react";
import { authOptions } from "@/app/api/auth/[...nextauth]/NextAuthOptions";

const JobCard = ({ job }: any) => {
  const router = useRouter();
  // @ts-expect-error
  const { data: session, status } = useSession(authOptions);
  const [isBidModalVisible, setIsBidModalVisible] = useState(false);

  return (
    <div className="max-w-sm rounded overflow-hidden shadow-lg bg-white p-6 hover:shadow-2xl transition-shadow duration-300">
      <div className="font-bold text-xl mb-2">{job.title}</div>
      <p className="text-gray-700 text-base">{job.description}</p>
      <div className="text-gray-800 font-semibold mt-4">
        Budget: ${job.budget}
      </div>
      {job.Deadline && (
        <div className="text-gray-600 mt-2">
          Deadline: {new Date(job.deadline).toLocaleDateString()}
        </div>
      )}

      {/* @ts-expect-error */}
      {session?.roles?.includes("talent") && (
        <button
          className="mt-4 bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 transition-colors duration-300"
          onClick={() => setIsBidModalVisible(true)}
        >
          Bid
        </button>
      )}

      {/* @ts-expect-error */}
      {session?.roles?.includes("client") && (
        <button
          className="mt-4 bg-violet-700 text-white px-4 py-2 rounded hover:bg-green-600 transition-colors duration-300"
          onClick={() => router.push(`/jobs/${job.id}/bidlist`)}
        >
          View Bids on this Job
        </button>
      )}

      {/* Modal */}
      {isBidModalVisible && (
        <BidJob jobID={job.id} onCancel={() => setIsBidModalVisible(false)} />
      )}
    </div>
  );
};

export default JobCard;
