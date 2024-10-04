"use client";
import { getTalents } from "@/services/adminService";
import { useCanActivateAdmin } from "@/utils/authorizeHelper";
import React, { useEffect, useState } from "react";

interface User {
  ID: string;
  Email: string;
  Phone?: string;
  IsVerified: boolean;
  Name: string;
}

const usersData: User[] = [
  { ID: "", Name: "John Doe", Email: "example@talent.com", IsVerified: false },
];

const TalentEditModal = ({
  selectedUser,
  handleClosePopup,
}: {
  selectedUser: User;
  handleClosePopup: any;
}) => {
  const handleToggleChange = () => {
    // TODO: call api
  };

  return (
    <div className="fixed inset-0 bg-gray-500 bg-opacity-50 flex justify-center items-center">
      <div className="bg-white p-6 rounded-lg shadow-lg w-80">
        <h2 className="text-xl font-semibold mb-4">{selectedUser.Name}</h2>
        <div className="flex items-center mb-4">
          <label className="mr-2">Is Verified:</label>
          <div
            className={`relative inline-block w-12 h-6 cursor-pointer ${
              selectedUser.IsVerified ? "bg-green-500" : "bg-gray-300"
            } rounded-full transition-colors`}
            onClick={handleToggleChange}
          >
            <span
              className={`absolute left-1 top-1 w-4 h-4 bg-white rounded-full transition-transform transform ${
                selectedUser.IsVerified ? "translate-x-6" : ""
              }`}
            />
          </div>
        </div>
        <button
          onClick={handleClosePopup}
          className="mt-4 bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600"
        >
          Close
        </button>
      </div>
    </div>
  );
};

const TalentList: React.FC = () => {
  useCanActivateAdmin();

  const [selectedUser, setSelectedUser] = useState<User | null>(null);
  const [isPopupOpen, setIsPopupOpen] = useState(false);

  const handleUserClick = (user: User) => {
    setSelectedUser(user);
    setIsPopupOpen(true);
  };

  const handleClosePopup = () => {
    setIsPopupOpen(false);
    setSelectedUser(null);
  };

  useEffect(() => {
    getTalents().then((res) => {
      console.log(res);
    });
  }, []);

  return (
    <div className="p-4">
      <h1 className="text-2xl font-bold mb-4">User List</h1>
      <ul>
        {usersData.map((user) => (
          <li key={user.ID} className="mb-2">
            <button
              onClick={() => handleUserClick(user)}
              className="text-blue-600 hover:underline"
            >
              {user.Name}
            </button>
          </li>
        ))}
      </ul>

      {isPopupOpen && selectedUser && (
        <TalentEditModal
          selectedUser={selectedUser}
          handleClosePopup={handleClosePopup}
        />
      )}
    </div>
  );
};

export default TalentList;
