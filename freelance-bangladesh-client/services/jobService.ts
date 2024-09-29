import { getAccessToken } from "@/utils/sessionTokenAccessor";

export const getAllJobs = async () => {
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

export const createJob = async (body: any) => {
  const url = `${process.env.API_URL}/api/v1/jobs`;

  let accessToken = await getAccessToken();

  const resp = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer " + accessToken,
    },
    body: JSON.stringify(body), // body should be inside the fetch options object
  });

  if (resp.ok) {
    const data = await resp.json();
    return data;
  }

  throw new Error("Failed to create data. Status: " + resp.status);
}
