"use client";

import { useCanActivateClient } from "@/utils/authorizeHelper";
import { useSession } from "next-auth/react";
import { useRouter } from "next/navigation";
import React, { useState, useEffect } from "react";

export default function CreateJobs() {
  const { status } = useSession();
  const router = useRouter();

  useCanActivateClient();

  const jobsNameRef = React.useRef();
  const priceRef = React.useRef();

  const [errorMsg, setErrorMsg] = useState("");

  if (status == "loading") {
    return (
      <main>
        <h1 className="text-4xl text-center">Create jobs</h1>
        <div className="text-center text-2xl">Loading...</div>
      </main>
    );
  }

  const handleSubmit = async (event) => {
    event.preventDefault();

    const postBody = {
      Name: jobsNameRef.current.value,
      Price: parseFloat(priceRef.current.value),
    };

    try {
      const resp = await fetch("/api/jobs", {
        method: "POST",
        headers: {
          headers: {
            "Content-Type": "application/json",
          },
        },
        body: JSON.stringify(postBody),
      });

      if (resp.ok) {
        router.push("/jobs");
        router.refresh();
      } else {
        var errMessage = await resp.text();
        setErrorMsg("Unable to create job: " + errMessage);
      }
    } catch (err) {
      setErrorMsg("Unable to create job: " + err);
    }
  };

  return (
    <main>
      <h1 className="text-4xl text-center">Create Job</h1>

      <form onSubmit={handleSubmit} className="mt-6">
        <div className="w-1/2">
          <label htmlFor="jobsName" className="text-2xl">Job name:</label>
          <input autoFocus type="text" id="jobsName" 
              className="w-full p-1 text-black bg-gray-200 text-lg" ref={jobsNameRef} required />
        </div>
        <div className="w-1/2 mt-2">
          <label htmlFor="price" className="text-2xl">
            Price:
          </label>
          <input type="number" step="0.01" id="price" className="w-full p-1 text-black bg-gray-200 text-lg" ref={priceRef} />
        </div>
        <div className="text-center text-2xl text-red-600">{errorMsg}</div>
        <button type="submit" className="mt-3 bg-blue-900 font-bold text-white py-1 px-2 rounded border border-gray-50">
          Create
        </button>
      </form>
    </main>
  );
}
