"use client";
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

  const handleUserClick = (item: any) => {
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
        console.log(res);
        if (res?.result) {
          setAssignmentlist(res.result);
        }
      })
      .catch((err) => {
        alert(err.message ?? "Some unexpected error has occurred.");
      });
  }, []);

  return (
    <div className="p-4">
      <h1 className="text-2xl font-bold mb-4">Assignment List For Talent</h1>
      <div className="overflow-x-auto">
        <table className="min-w-full table-auto bg-white border border-gray-200 rounded-md">
          <thead className="bg-gray-100">
            <tr>
              <th className="px-4 py-2 text-left text-sm font-medium text-gray-600">
                #
              </th>
              <th className="px-4 py-2 text-left text-sm font-medium text-gray-600">
                Job ID
              </th>
              <th className="px-4 py-2 text-left text-sm font-medium text-gray-600">
                Amount
              </th>
              <th className="px-4 py-2 text-left text-sm font-medium text-gray-600">
                Message
              </th>
              <th className="px-4 py-2 text-left text-sm font-medium text-gray-600">
                Status
              </th>
              <th className="px-4 py-2 text-left text-sm font-medium text-gray-600">
                Actions
              </th>
            </tr>
          </thead>
          <tbody>
            {assignmentlist.map((item: any, index) => (
              <tr
                key={item.assignment_id}
                className="border-t hover:bg-gray-100 transition duration-200"
              >
                <td className="px-4 py-2 text-sm text-gray-700">{index + 1}</td>
                <td className="px-4 py-2 text-sm text-gray-700">
                  {item.job_id}
                </td>
                <td className="px-4 py-2 text-sm text-gray-700">
                  {item.amount ?? "-"}
                </td>
                <td className="px-4 py-2 text-sm text-gray-700">
                  {item.message ?? "-"}
                </td>
                <td className="px-4 py-2 text-sm text-gray-700">
                  {item.status}
                </td>
                <td className="px-4 py-2">
                  <button
                    onClick={() => handleUserClick(item)}
                    className="text-blue-600 hover:underline"
                  >
                    See Reviews
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      {isPopupOpen && selected && (
        <SeeReviews selected={selected} handleClosePopup={handleClosePopup} />
      )}
    </div>
  );
}
