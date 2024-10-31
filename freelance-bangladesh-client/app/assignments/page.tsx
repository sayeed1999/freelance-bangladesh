"use client";
import DynamicList, { Column } from "@/components/dynamic-list";
import { getAssignments } from "@/services/assignmentService";
import { useCanActivateTalent } from "@/utils/authorizeHelper";
import { useEffect, useState } from "react";

const SeeReviews = ({ selected, handleClosePopup }: any) => {
  return (
    <div className="fixed inset-0 bg-gray-500 bg-opacity-50 flex justify-center items-center">
      <div className="bg-white p-8 rounded-lg shadow-lg max-w-sm w-full text-center">
        <h2 className="text-xl font-semibold mb-4">Reviews</h2>
        <div className="flex justify-center space-x-4">
          <button
            onClick={handleClosePopup}
            className="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600 transition-colors duration-300"
          >
            Close
          </button>
        </div>
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
        actionTitle="See Reviews"
      />

      {isPopupOpen && selected && (
        <SeeReviews selected={selected} handleClosePopup={handleClosePopup} />
      )}
    </>
  );
}
