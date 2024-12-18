import { callApi } from "./api";

export const getAssignments = () =>
    callApi({
        url: `/api/assignments`
    });

export const assignTalent = (body: any) => 
    callApi({ 
        url: `/api/assignments`, 
        method: "POST", 
        body 
    });

export const submitWork = (assignmentID: string, body: any) =>
    callApi({ 
        url: `/api/assignments/${assignmentID}`, 
        method: "PATCH", 
        body
    });

export const addReview = (assignmentID: string, body: any) => 
    callApi({ 
        url: `/api/assignments/${assignmentID}/reviews`, 
        method: "POST", 
        body 
    });
  
export const getReviewList = (assignmentID: string) => 
    callApi({ 
        url: `/api/assignments/${assignmentID}/reviews`, 
    });
    