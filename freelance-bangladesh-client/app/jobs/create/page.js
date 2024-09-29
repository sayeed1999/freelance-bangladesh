"use client";
import React, { useState, useEffect } from "react";
import { useCanActivateClient } from "@/utils/authorizeHelper";
import JobCreate from "@/components/jobs/jobCreate";

export default function CreateJobs() {
  const { status } = useCanActivateClient();

  if (status == "loading") {
    return (
      <main>
        <h1 className="text-4xl text-center">Create Job</h1>
        <div className="text-center text-2xl">Loading...</div>
      </main>
    );
  }

  return (
    <main>
      <JobCreate />
    </main>
  );
}
