"use client";
import React from "react";
import Form from "../form";
import { createJob } from "@/services/jobService";

const JobCreate = () => {
    const titleRef = React.useRef();
    const descriptionRef = React.useRef();
    const budgetRef = React.useRef();
    const deadlineRef = React.useRef();

    const handleCreate = () => {
      createJob({
        title: titleRef.current.value,
        description: descriptionRef.current.value,
        budget: parseFloat(budgetRef.current.value),
        deadline: deadlineRef.current.value ? new Date(deadlineRef.current.value).toISOString(): null,
      }).then(() => {
        titleRef.current.value = null;
        descriptionRef.current.value = null;
        budgetRef.current.value = null;
        deadlineRef.current.value = null;
        alert("create success!")
      }).catch((err) => {
        alert(err.message ?? "Some unexpected error has occurred.")
      })
    }

    return (
      <Form
        formTitle="Create Job"
        submitBtnName="Create"
        dispatchAction={handleCreate}
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
      />);
}

export default JobCreate;
