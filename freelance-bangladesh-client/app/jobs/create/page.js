"use client";

import Form from "@/components/form";
import { createJob } from "@/services/jobService";
import { useCanActivateClient } from "@/utils/authorizeHelper";
import { useSession } from "next-auth/react";
import { useRouter } from "next/navigation";
import React, { useState } from "react";

export default function CreateJobs() {
  const { status } = useSession();
  const router = useRouter();

  useCanActivateClient();

  const titleRef = React.useRef();
  const descriptionRef = React.useRef();
  const budgetRef = React.useRef();
  const deadlineRef = React.useRef();

  const [errorMsg, setErrorMsg] = useState("");

  if (status == "loading") {
    return (
      <main>
        <h1 className="text-4xl text-center">Create jobs</h1>
        <div className="text-center text-2xl">Loading...</div>
      </main>
    );
  }

  return (
    <main>
      <Form
        formTitle="Create Job"
        submitBtnName="Create"
        dispatchAction={createJob}
        formItems={[
          {
            label: "Job Title",
            name: "title",
            ref: titleRef,
            type: "text",
            id: "text",
            placeholder: "Design Custom Wordpress Theme",
            required: true,
          },
          {
            label: "Job Description",
            name: "description",
            ref: descriptionRef,
            type: "textarea",
            id: "description",
            placeholder: "Design Custom Wordpress Theme which shall inlude (100 lines)...",
            required: true,
          },
          {
            label: "Budget",
            name: "budget",
            ref: budgetRef,
            type: "number",
            id: "budget",
            // min: 0,
            placeholder: "(in BDT)",
            required: true,
          },
          {
            label: "Deadline",
            name: "budget",
            ref: deadlineRef,
            type: "date",
            id: "deadline",
            required: false,
          },
        ]}
      />
    </main>
  );
}
