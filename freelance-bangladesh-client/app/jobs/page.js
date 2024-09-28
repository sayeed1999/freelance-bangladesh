"use client";

import { redirect } from "next/navigation";
import { getAccessToken } from "@/utils/sessionTokenAccessor";
import { SetDynamicRoute } from "@/utils/setDynamicRoute";
import { canActivateTalent } from "@/utils/authorizeHelper";

async function getAllJobs() {
  const url = `${process.env.API_URL}/api/v1/jobs`;

  let accessToken = await getAccessToken();

  const resp = await fetch(url, {
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer " + accessToken,
    },
  });

  if (resp.ok) {
    const data = await resp.json();
    return data;
  }

  throw new Error("Failed to fetch data. Status: " + resp.status);
}

export default async function Jobs() {
  if (!canActivateTalent()) {
    redirect("/unauthorized");
  }

  try {
    const jobs = await getAllJobs();          
    

    return (        
      <main>  
        <SetDynamicRoute></SetDynamicRoute>    
        <h1 className="text-4xl text-center">Jobs</h1>
        <table className="border border-gray-500 text-lg ml-auto mr-auto mt-6">
          <thead>
            <tr>
              <th className="bg-blue-900 p-2 border border-gray-500">Id</th>
              <th className="bg-blue-900 p-2 border border-gray-500">Name</th>
              <th className="bg-blue-900 p-2 border border-gray-500">
                Price
              </th>
            </tr>
          </thead>
          <tbody>
            {jobs.map((p) => (
              <tr key={p.Id}>
                <td className="p-1 border border-gray-500">{p.Id}</td>
                <td className="p-1 border border-gray-500">{p.Name}</td>
                <td className="p-1 border border-gray-500">{p.Price}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </main>
    );
  } catch (err) {
    console.error(err);

    return (
      <main>
        <h1 className="text-4xl text-center">Jobs</h1>
        <p className="text-red-600 text-center text-lg">
          Sorry, an error happened. Check the server logs.
        </p>
      </main>
    );
  }
}