'use client';
import { assignTalent } from '@/services/assignmentService';
import { getBidList } from '@/services/jobService';
import { useCanActivateClient } from '@/utils/authorizeHelper';
import { useParams } from 'next/navigation';
import { useEffect, useState } from 'react';

const AssignJobModal = ({
  handleAssign,
  selected,
  handleClosePopup,
}) => {

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
    const {jobid} = useParams();

    const [bidlist, setBidlist] = useState([]);
    const [selected, setSelected] = useState(null);
    const [isPopupOpen, setIsPopupOpen] = useState(false);
  
    const handleUserClick = (item) => {
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
          console.log(res);
          if (res?.result) {
            setBidlist(res.result);
          }
        })
        .catch((err) => {
          alert(err.message ?? "Some unexpected error has occurred.");
        });
    }, [jobid]);

    return (
        <div className="p-4">
        <h1 className="text-2xl font-bold mb-4">Bid List For Job - {jobid}</h1>
        <div className="overflow-x-auto">
          <table className="min-w-full table-auto bg-white border border-gray-200 rounded-md">
            <thead className="bg-gray-100">
              <tr>
                <th className="px-4 py-2 text-left text-sm font-medium text-gray-600">
                  #
                </th>
                <th className="px-4 py-2 text-left text-sm font-medium text-gray-600">
                  Talent Name
                </th>
                <th className="px-4 py-2 text-left text-sm font-medium text-gray-600">
                  Talent Email
                </th>
                <th className="px-4 py-2 text-left text-sm font-medium text-gray-600">
                  Amount
                </th>
                <th className="px-4 py-2 text-left text-sm font-medium text-gray-600">
                  Message
                </th>
                <th className="px-4 py-2 text-left text-sm font-medium text-gray-600">
                  Actions
                </th>
              </tr>
            </thead>
            <tbody>
              {bidlist.map((item, index) => (
                <tr
                  key={item.ID}
                  className="border-t hover:bg-gray-100 transition duration-200"
                >
                  <td className="px-4 py-2 text-sm text-gray-700">{index + 1}</td>
                  <td className="px-4 py-2 text-sm text-gray-700">{item.talent_name}</td>
                  <td className="px-4 py-2 text-sm text-gray-700">
                    {item.talent_email}
                  </td>
                  <td className="px-4 py-2 text-sm text-gray-700">
                    {item.amount ?? "-"}
                  </td>
                  <td className="px-4 py-2 text-sm text-gray-700">
                    {item.message ?? "-"}
                  </td>
                  <td className="px-4 py-2">
                    <button
                      onClick={() => handleUserClick(item)}
                      className="text-blue-600 hover:underline"
                    >
                      Assign
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
  
        {isPopupOpen && selected && (
          <AssignJobModal
            handleAssign={handleAssign}
            selected={selected}
            handleClosePopup={handleClosePopup}
          />
        )}
      </div>
    );
}

/// fields in bid item entity: -
// amount
// bid_id
// job_id
// message
// talent_email
// talent_id
// talent_name
