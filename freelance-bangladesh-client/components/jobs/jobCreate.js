"use client";
import React from "react";
import Form from "../form";
import { createJob } from "@/services/jobService";

const JobCreate = () => {
    const titleRef = React.useRef();
    const descriptionRef = React.useRef();
    const budgetRef = React.useRef();
    const deadlineRef = React.useRef();

    return (
      <Form
        formTitle="Create Job"
        submitBtnName="Create"
        dispatchAction={() => {
          createJob({
            title: titleRef.current.value,
            description: descriptionRef.current.value,
            budget: budgetRef.current.value,
            deadline: deadlineRef.current.value,
          }).then(() => {
            alert("create success!")
          }).catch((err) => {
            alert(err.message)
          })
        }}
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
