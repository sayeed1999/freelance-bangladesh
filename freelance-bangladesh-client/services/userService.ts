import { getAccessToken } from "@/utils/sessionTokenAccessor";
import { callApi } from "./api";

export const syncUser = async () => {
  let accessToken = await getAccessToken() || "";

  return await callApi({ 
    url: `${process.env.API_URL}/api/v1/users/sync-user`,
    method: "POST",
    accessToken,
  });
}

export const signupClient = (body: any) => 
  callApi({
    url: `/api/users/client-signup`,
    method: "POST",
    body 
  });

export const signupTalent = (body: any) => 
  callApi({
    url: `/api/users/talent-signup`,
    method: "POST",
    body 
  });
