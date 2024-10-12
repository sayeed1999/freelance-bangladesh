import { bidJob } from "@/services/jobService";
import React from "react";

interface Props {
  jobID: string;
  onCancel: () => void;
}

const BidJob = ({ jobID, onCancel }: Props) => {
  const amountRef = React.useRef();
  const messageRef = React.useRef();

  const onSubmit = () => {
    bidJob(jobID, {
      // @ts-expect-error
      amount: parseFloat(amountRef.current.value),
      // @ts-expect-error
      message: messageRef.current.value,
    })
      .then(() => {
        // @ts-expect-error
        amountRef.current.value = null;
        // @ts-expect-error
        messageRef.current.value = null;
        alert("bid success!");
      })
      .catch((err: Error) => {
        alert(err.message ?? "Some unexpected error has occurred.");
      });
  };

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center z-50">
      <div className="bg-white p-8 rounded shadow-lg max-w-sm w-full">
        <h2 className="text-xl font-bold mb-4">Place your bid</h2>
        <form>
          <label className="block text-gray-700 text-sm font-bold mb-2">
            Amount:
          </label>
          <input
            // @ts-expect-error
            ref={amountRef}
            type="number"
            className="w-full px-4 py-2 border rounded mb-4"
            placeholder="Enter your bid amount"
          />
          <label className="block text-gray-700 text-sm font-bold mb-2">
            Message (optional):
          </label>
          <textarea
            // @ts-expect-error
            ref={messageRef}
            className="w-full px-4 py-2 border rounded mb-4"
            placeholder="Enter your message"
          ></textarea>
          <div className="flex justify-end">
            <button
              type="button"
              className="bg-red-500 text-white px-4 py-2 rounded mr-2 hover:bg-red-600"
              onClick={onCancel}
            >
              Cancel
            </button>
            <button
              type="submit"
              className="bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600"
              onClick={onSubmit}
            >
              Submit Bid
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default BidJob;
