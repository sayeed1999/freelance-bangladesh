import { makeAuthorizedHttpRequest } from "../../../helpers/http-request";

export async function POST(req: any) {
    const urlParts = req.url.split('/');
    const assignmentid = urlParts[urlParts.length - 2];

    const url = `${process.env.API_URL}/api/v1/assignments/${assignmentid}/reviews`;
    const body = await req.json();

    return await makeAuthorizedHttpRequest(url, "POST", body);
}
