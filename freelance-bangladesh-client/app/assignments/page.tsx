"use client";
import DynamicList, { Column } from "@/components/dynamic-list";
import Form from "@/components/form";
import {
  getAssignments,
  getReviewList,
  submitWork,
} from "@/services/assignmentService";
import { useCanActivateTalent } from "@/utils/authorizeHelper";
import { useEffect, useRef, useState } from "react";

const SubmitOrSeeReviews = ({ selectedAssignment, handleClosePopup }: any) => {
  const submissionUrlRef = useRef();
  const commentsRef = useRef();
  const [reviewlist, setReviewlist] = useState([]);

  const columns: Column<any>[] = [
    { header: "Review ID", accessor: "review_id" },
    { header: "Comments", accessor: "comments" },
  ];

  const handleSubmit = () => {
    submitWork(selectedAssignment.assignment_id, {
      // @ts-expect-error
      title: submissionUrlRef.current.value,
      // @ts-expect-error
      description: commentsRef.current.value,
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

  const submitAssignmentForm = (
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

  useEffect(() => {
    getReviewList(selectedAssignment.assignment_id)
      .then((res) => {
        if (res?.result) {
          setReviewlist(res.result);
        }
      })
      .catch((err) => {
        alert(err.message ?? "Some unexpected error has occurred.");
      });
  }, []);

  return (
    <div className="fixed inset-0 bg-gray-900 bg-opacity-60 flex justify-center items-center">
      <div className="bg-white rounded-lg shadow-lg max-w-4xl w-full mx-4 p-6 flex space-x-8">
        {/* Left side - Form */}
        <div className="flex-1">{submitAssignmentForm}</div>

        {/* Right side - Reviews */}
        <div className="flex-1 bg-gray-100 p-2 rounded-lg shadow-md flex flex-col items-center">
          <DynamicList items={reviewlist} columns={columns} title="Reviews" />
        </div>

        {/* Close button in bottom right */}
        <button
          onClick={handleClosePopup}
          className="absolute bottom-4 right-4 bg-red-500 text-white px-4 py-2 rounded-lg hover:bg-red-600 transition duration-300 shadow-md"
        >
          Close
        </button>
      </div>
    </div>
  );
};

export default function AssignmentListPage() {
  useCanActivateTalent();

  const [assignmentlist, setAssignmentlist] = useState([]);
  const [selected, setSelected] = useState(null);
  const [isPopupOpen, setIsPopupOpen] = useState(false);

  const columns: Column<any>[] = [
    { header: "Job ID", accessor: "job_id" },
    { header: "Amount", accessor: "amount" },
    { header: "Status", accessor: "status" },
  ];

  const handleActionClick = (item: any) => {
    setSelected(item);
    setIsPopupOpen(true);
  };

  const handleClosePopup = () => {
    setIsPopupOpen(false);
    setSelected(null);
  };

  useEffect(() => {
    getAssignments()
      .then((res) => {
        if (res?.result) {
          setAssignmentlist(res.result);
        }
      })
      .catch((err) => {
        alert(err.message ?? "Some unexpected error has occurred.");
      });
  }, []);

  return (
    <>
      <DynamicList
        items={assignmentlist}
        columns={columns}
        title="Assignment List"
        onActionClick={handleActionClick}
        actionTitle="Submit/ See Reviews"
      />

      {isPopupOpen && selected && (
        <SubmitOrSeeReviews
          selectedAssignment={selected}
          handleClosePopup={handleClosePopup}
        />
      )}
    </>
  );
}
