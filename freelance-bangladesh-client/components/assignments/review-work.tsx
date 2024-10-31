"use client";
import Form from "@/components/form";
import { addReview } from "@/services/assignmentService";
import { useRef } from "react";

const ReviewWork = ({ assignment_id }: any) => {
  const commentsRef = useRef();
  const actionRef = useRef();

  const handleReview = () => {
    addReview(assignment_id, {
      // @ts-expect-error
      comments: commentsRef.current.value,
      // @ts-expect-error
      action: actionRef.current.value,
    })
      .then(() => {
        // @ts-expect-error
        commentsRef.current.value = null;
        // @ts-expect-error
        actionRef.current.value = null;
        alert("review given!");
      })
      .catch((err) => {
        alert(err.message ?? "Some unexpected error has occurred.");
      });
  };

  return (
    <Form
      formTitle="Review Assignment"
      submitBtnName="Review"
      dispatchAction={handleReview}
      formItems={[
        {
          label: "Comments",
          name: "comments",
          ref: commentsRef,
          type: "textarea",
          id: "comments",
          placeholder: "Some comments...",
        },
        {
          label: "Action",
          name: "action",
          ref: actionRef,
          type: "select",
          id: "action",
          placeholder: "Choose an action...",
          options: [
            { label: "Change Request", value: "change-request" },
            { label: "Approve", value: "approve" },
            { label: "Reject", value: "reject" },
          ],
        },
      ]}
    />
  );
};

export default ReviewWork;
