import { callApi } from "./api";

export const signupUser = (body: any) => 
  callApi({
    url: `/api/users`,
    method: "POST",
    body 
  });
