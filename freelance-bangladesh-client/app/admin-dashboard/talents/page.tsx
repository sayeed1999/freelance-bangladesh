"use client";
import React, { useEffect, useState, useRef } from "react";
import { Talent } from "@/models/user";
import { getTalents, updateTalent } from "@/services/adminService";
import { useCanActivateAdmin } from "@/utils/authorizeHelper";
import DynamicList, { Column } from "@/components/dynamic-list";

const TalentEditModal = ({
  selectedUser,
  handleClosePopup,
}: {
  selectedUser: Talent;
  handleClosePopup: any;
}) => {
  const [isVerified, setIsVerified] = useState(selectedUser.is_verified);

  const handleSave = () => {
    updateTalent({
      talent_id: selectedUser.id,
      is_verified: isVerified,
    })
      .then(() => {
        alert("success!");
        handleClosePopup();
      })
      .catch((err) => {
        alert(err.message ?? "Some unexpected error has occurred.");
      });
  };

  return (
    <div className="fixed inset-0 bg-gray-500 bg-opacity-50 flex justify-center items-center">
      <div className="bg-white p-6 rounded-lg shadow-lg w-80">
        <h2 className="text-xl font-semibold mb-4">{selectedUser.name}</h2>
        <div className="flex items-center mb-4">
          <label className="mr-2">Is Verified:</label>
          <div
            className={`relative inline-block w-12 h-6 cursor-pointer ${
              isVerified ? "bg-green-500" : "bg-gray-300"
            } rounded-full transition-colors`}
            onClick={() => setIsVerified(!isVerified)}
          >
            <span
              className={`absolute left-1 top-1 w-4 h-4 bg-white rounded-full transition-transform transform ${
                isVerified ? "translate-x-6" : ""
              }`}
            />
          </div>
        </div>
        <button
          onClick={handleSave}
          className="mt-4 bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600"
        >
          Save
        </button>
        <button
          onClick={handleClosePopup}
          className="ml-2 mt-4 bg-gray-500 text-white px-4 py-2 rounded hover:bg-gray-600"
        >
          Cancel
        </button>
      </div>
    </div>
  );
};

const TalentList: React.FC = () => {
  useCanActivateAdmin();

  const [users, setUsers] = useState<Talent[]>([]);
  const [selectedUser, setSelectedUser] = useState<Talent | null>(null);
  const [isPopupOpen, setIsPopupOpen] = useState(false);

  const columns: Column<any>[] = [
    { header: "Name", accessor: "name" },
    { header: "Email", accessor: "email" },
    { header: "Phone", accessor: "phone" },
    { header: "Is Verfied", accessor: "is_verified" },
  ];

  const handleUserClick = (user: Talent) => {
    setSelectedUser(user);
    setIsPopupOpen(true);
  };

  const handleClosePopup = () => {
    setIsPopupOpen(false);
    setSelectedUser(null);
  };

  useEffect(() => {
    getTalents()
      .then((res) => {
        if (res?.result) {
          setUsers(res.result);
        }
      })
      .catch((err) => {
        alert(err.message ?? "Some unexpected error has occurred.");
      });
  }, []);

  return (
    <>
      <DynamicList
        items={users}
        columns={columns}
        title="Talent List"
        onActionClick={handleUserClick}
        actionTitle="Verify"
      />

      {isPopupOpen && selectedUser && (
        <TalentEditModal
          selectedUser={selectedUser}
          handleClosePopup={handleClosePopup}
        />
      )}
    </>
  );
};

export default TalentList;
