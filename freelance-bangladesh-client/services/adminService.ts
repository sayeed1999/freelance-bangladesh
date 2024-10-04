import { callApi } from "./api";

export const getClients = () => 
  callApi({
    url: `/api/admin-dashboard/clients`
  });

export const getTalents = () => 
  callApi({
    url: `/api/admin-dashboard/talents`
  });

export const updateClient = (body: any) => 
  callApi({
    url: `/api/admin-dashboard/clients`,
    method: "POST",
    body 
  });

export const updateTalent = (body: any) => 
  callApi({
    url: `/api/admin-dashboard/talents`,
    method: "POST",
    body 
  });
  