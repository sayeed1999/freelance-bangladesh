import { callApi } from "./api";

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
