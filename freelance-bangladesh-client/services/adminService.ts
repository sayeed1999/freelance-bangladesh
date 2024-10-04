import { callApi } from "./api";

export const getClients = () => 
  callApi({
    url: `/api/admin-dashboard/clients`
  });

export const getTalents = () => 
  callApi({
    url: `/api/admin-dashboard/talents`
  });

export const verifyClient = (body: any) => 
  callApi({
    url: `/api/admin-dashboard/clients/verify`,
    method: "POST",
    body 
  });

export const verifyTalent = (body: any) => 
  callApi({
    url: `/api/admin-dashboard/talents/verify`,
    method: "POST",
    body 
  });
  