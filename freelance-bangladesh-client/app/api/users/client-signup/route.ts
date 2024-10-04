import { makeAuthorizedHttpRequest } from "../../helpers/http-request";

export async function POST(req: any) {
    const url = `${process.env.API_URL}/api/v1/users/client-signup`;
    const body = await req.json();

    return await makeAuthorizedHttpRequest(url, "POST", body);
}