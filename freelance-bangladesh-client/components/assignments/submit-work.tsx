"use client";
import Form from "@/components/form";
import { submitWork } from "@/services/assignmentService";
import { useRef } from "react";

const SubmitWork = ({ assignment_id }: any) => {
  const submissionUrlRef = useRef();
  const commentsRef = useRef();

  const handleSubmit = () => {
    submitWork(assignment_id, {
      // @ts-expect-error
      submission_url: submissionUrlRef.current.value,
      // @ts-expect-error
      comments: commentsRef.current.value,
    })
      .then(() => {
        // @ts-expect-error
        submissionUrlRef.current.value = null;
        // @ts-expect-error
        commentsRef.current.value = null;
        alert("submit success!");
      })
      .catch((err) => {
        alert(err.message ?? "Some unexpected error has occurred.");
      });
  };

  return (
    <Form
      formTitle="Submit Assignment"
      submitBtnName="Submit"
      dispatchAction={handleSubmit}
      formItems={[
        {
          label: "Submission URL",
          name: "submission_url",
          ref: submissionUrlRef,
          type: "text",
          id: "submission_url",
          placeholder: "Google Drive link",
          required: true,
          validationError: "Submission URL must be provided",
        },
        {
          label: "Comments",
          name: "comments",
          ref: commentsRef,
          type: "textarea",
          id: "comments",
          placeholder: "Some comments...",
        },
      ]}
    />
  );
};

export default SubmitWork;
