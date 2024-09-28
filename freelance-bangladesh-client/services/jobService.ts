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
