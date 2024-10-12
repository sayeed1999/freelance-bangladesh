import { getAccessToken } from "@/utils/sessionTokenAccessor";
import { callApi } from "./api";

export const getAllJobs = async () => {
  let accessToken = await getAccessToken() || "";

  return await callApi({ 
    url: `${process.env.API_URL}/api/v1/jobs`, 
    accessToken,
  });
}

// Note: This api lives inside next.js server
export const createJob = (body: any) => 
  callApi({ 
    url: `/api/jobs`, 
    method: "POST", 
    body 
  });

export const bidJob = (jobID: string, body: any) => 
  callApi({ 
    url: `/api/jobs/${jobID}/bids`, 
    method: "POST", 
    body 
  });
