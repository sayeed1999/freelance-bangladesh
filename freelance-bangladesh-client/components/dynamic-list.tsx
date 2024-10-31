"use client";
import React, { useState } from "react";

// Define the column structure type to include a header and an accessor key for each column
export interface Column<T> {
  header: string;
  accessor: keyof T;
}

interface DynamicListProps<T> {
  items: T[];
  columns: Column<T>[];
  onActionClick?: (item: T) => void;
  actionTitle?: string;
  title?: string;
}

const DynamicList = <T extends { ID: string | number }>({
  items,
  columns,
  onActionClick,
  title = "List",
  actionTitle = "Edit",
}: DynamicListProps<T>) => {
  return (
    <div className="p-4">
      <h1 className="text-2xl font-bold mb-4">{title}</h1>
      <div className="overflow-x-auto">
        <table className="min-w-full table-auto bg-white border border-gray-200 rounded-md">
          <thead className="bg-gray-100">
            <tr>
              <th className="px-4 py-2 text-left text-sm font-medium text-gray-600">
                #
              </th>
              {columns.map((column, index) => (
                <th
                  key={index}
                  className="px-4 py-2 text-left text-sm font-medium text-gray-600"
                >
                  {column.header}
                </th>
              ))}
              {onActionClick && (
                <th className="px-4 py-2 text-left text-sm font-medium text-gray-600">
                  Actions
                </th>
              )}
            </tr>
          </thead>
          <tbody>
            {items.map((item, index) => (
              <tr
                key={item.ID}
                className="border-t hover:bg-gray-100 transition duration-200"
              >
                <td className="px-4 py-2 text-sm text-gray-700">{index + 1}</td>
                {columns.map((column, colIndex) => (
                  <td
                    key={colIndex}
                    className="px-4 py-2 text-sm text-gray-700"
                  >
                    {String(item[column.accessor])}
                  </td>
                ))}
                {onActionClick && (
                  <td className="px-4 py-2">
                    <button
                      onClick={() => onActionClick(item)}
                      className="text-blue-600 hover:underline"
                    >
                      {actionTitle}
                    </button>
                  </td>
                )}
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default DynamicList;
