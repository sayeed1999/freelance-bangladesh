import { makeAuthorizedHttpRequest } from "../../helpers/http-request";

export async function PATCH(req: any) {
    const urlParts = req.url.split('/');
    const assignmentid = urlParts[urlParts.length - 1];

    const url = `${process.env.API_URL}/api/v1/assignments/${assignmentid}`;
    const body = await req.json();

    return await makeAuthorizedHttpRequest(url, "PATCH", body);
}
