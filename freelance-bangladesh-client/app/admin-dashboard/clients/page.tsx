"use client";
import { mockClients } from "@/mock_data/mockUsers";
import { Client } from "@/models/user";
import { getClients } from "@/services/adminService";
import { useCanActivateAdmin } from "@/utils/authorizeHelper";
import React, { useEffect, useState } from "react";

const ClientEditModal = ({
  selectedUser,
  handleClosePopup,
}: {
  selectedUser: Client;
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

const ClientList: React.FC = () => {
  useCanActivateAdmin();

  const [users, setUsers] = useState<Client[]>(mockClients);
  const [selectedUser, setSelectedUser] = useState<Client | null>(null);
  const [isPopupOpen, setIsPopupOpen] = useState(false);

  const handleUserClick = (user: Client) => {
    setSelectedUser(user);
    setIsPopupOpen(true);
  };

  const handleClosePopup = () => {
    setIsPopupOpen(false);
    setSelectedUser(null);
  };

  useEffect(() => {
    getClients().then((res) => {
      console.log(res);
    });
  }, []);

  return (
    <div className="p-4">
      <h1 className="text-2xl font-bold mb-4">User List</h1>
      <div className="overflow-x-auto">
        <table className="min-w-full table-auto bg-white border border-gray-200 rounded-md">
          <thead className="bg-gray-100">
            <tr>
              <th className="px-4 py-2 text-left text-sm font-medium text-gray-600">
                #
              </th>
              <th className="px-4 py-2 text-left text-sm font-medium text-gray-600">
                Name
              </th>
              <th className="px-4 py-2 text-left text-sm font-medium text-gray-600">
                Email
              </th>
              <th className="px-4 py-2 text-left text-sm font-medium text-gray-600">
                Phone
              </th>
              <th className="px-4 py-2 text-left text-sm font-medium text-gray-600">
                Actions
              </th>
            </tr>
          </thead>
          <tbody>
            {users.map((user, index) => (
              <tr
                key={user.ID}
                className="border-t hover:bg-gray-100 transition duration-200"
              >
                <td className="px-4 py-2 text-sm text-gray-700">{index + 1}</td>
                <td className="px-4 py-2 text-sm text-gray-700">{user.Name}</td>
                <td className="px-4 py-2 text-sm text-gray-700">
                  {user.Email}
                </td>
                <td className="px-4 py-2 text-sm text-gray-700">
                  {user.Phone}
                </td>
                <td className="px-4 py-2">
                  <button
                    onClick={() => handleUserClick(user)}
                    className="text-blue-600 hover:underline"
                  >
                    Edit
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      {isPopupOpen && selectedUser && (
        <ClientEditModal
          selectedUser={selectedUser}
          handleClosePopup={handleClosePopup}
        />
      )}
    </div>
  );
};

export default ClientList;
