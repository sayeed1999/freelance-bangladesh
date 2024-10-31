"use client";
import ReviewWork from "@/components/assignments/review-work";
import SubmitWork from "@/components/assignments/submit-work";
import DynamicList, { Column } from "@/components/dynamic-list";
import { getAssignments, getReviewList } from "@/services/assignmentService";
import { useCanActivatePrivateRoute } from "@/utils/authorizeHelper";
import { useSession } from "next-auth/react";
import { useEffect, useState } from "react";

const SubmitOrReview = ({ selectedAssignment, handleClosePopup }: any) => {
  const { data: session, status } = useSession();
  const [reviewlist, setReviewlist] = useState([]);

  // @ts-expect-error
  const isClient = session?.roles?.includes("client") ?? false;
  // @ts-expect-error
  const isTalent = session?.roles?.includes("talent") ?? false;

  const columns: Column<any>[] = [
    { header: "Review ID", accessor: "review_id" },
    { header: "Comments", accessor: "comments" },
  ];

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
        {isTalent && (
          <div className="flex-1">
            <SubmitWork assignment_id={selectedAssignment.assignment_id} />
          </div>
        )}
        {isClient && (
          <div className="flex-1">
            <ReviewWork assignment_id={selectedAssignment.assignment_id} />
          </div>
        )}

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
  useCanActivatePrivateRoute();

  const [assignmentlist, setAssignmentlist] = useState([]);
  const [selected, setSelected] = useState(null);
  const [isPopupOpen, setIsPopupOpen] = useState(false);

  const columns: Column<any>[] = [
    { header: "Job ID", accessor: "job_id" },
    { header: "Amount", accessor: "amount" },
    { header: "Submission URL", accessor: "submission_url" },
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
        <SubmitOrReview
          selectedAssignment={selected}
          handleClosePopup={handleClosePopup}
        />
      )}
    </>
  );
}
