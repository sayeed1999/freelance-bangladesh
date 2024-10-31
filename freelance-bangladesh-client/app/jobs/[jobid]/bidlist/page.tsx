"use client";
import DynamicList, { Column } from "@/components/dynamic-list";
import { assignTalent } from "@/services/assignmentService";
import { getBidList } from "@/services/jobService";
import { useCanActivateClient } from "@/utils/authorizeHelper";
import { useParams } from "next/navigation";
import { useEffect, useState } from "react";

const AssignJobModal = ({ handleAssign, selected, handleClosePopup }: any) => {
  return (
    <div className="fixed inset-0 bg-gray-500 bg-opacity-50 flex justify-center items-center">
      <div className="bg-white p-8 rounded-lg shadow-lg max-w-sm w-full text-center">
        <h2 className="text-xl font-semibold mb-4">Confirm Assignment</h2>
        <p className="mb-6 text-gray-700">
          Are you sure you want to assign the job to this talent?
        </p>
        <div className="flex justify-center space-x-4">
          <button
            onClick={handleAssign}
            className="bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600 transition-colors duration-300"
          >
            Yes
          </button>
          <button
            onClick={handleClosePopup}
            className="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600 transition-colors duration-300"
          >
            No
          </button>
        </div>
      </div>
    </div>
  );
};

export default function BidListPage() {
  useCanActivateClient();
  const { jobid }: { jobid: string } = useParams();

  const [bidlist, setBidlist] = useState([]);
  const [selected, setSelected] = useState<any>(null);
  const [isPopupOpen, setIsPopupOpen] = useState(false);

  var columns: Column<any>[] = [
    { header: "Bid ID", accessor: "bid_id" },
    { header: "Job ID", accessor: "job_id" },
    { header: "Talent Email", accessor: "talent_email" },
    { header: "Talent Name", accessor: "talent_name" },
    { header: "Amount", accessor: "amount" },
    { header: "Message", accessor: "message" },
  ];

  const handleUserClick = (item: any) => {
    setSelected(item);
    setIsPopupOpen(true);
  };

  const handleClosePopup = () => {
    setIsPopupOpen(false);
    setSelected(null);
  };

  const handleAssign = () => {
    assignTalent({
      job_id: jobid,
      talent_id: selected.talent_id,
      amount: selected.amount,
    })
      .then(() => {
        alert("success!");
        handleClosePopup();
      })
      .catch((err) => {
        alert(err.message ?? "Some unexpected error has occurred.");
      });
  };

  useEffect(() => {
    getBidList(jobid)
      .then((res) => {
        if (res?.result) {
          setBidlist(res.result);
        }
      })
      .catch((err) => {
        alert(err.message ?? "Some unexpected error has occurred.");
      });
  }, [jobid]);

  return (
    <>
      <DynamicList
        items={bidlist}
        columns={columns}
        title="Bid List For Job"
        onActionClick={handleUserClick}
        actionTitle="Assign Talent"
      />

      {isPopupOpen && selected && (
        <AssignJobModal
          handleAssign={handleAssign}
          selected={selected}
          handleClosePopup={handleClosePopup}
        />
      )}
    </>
  );
}
