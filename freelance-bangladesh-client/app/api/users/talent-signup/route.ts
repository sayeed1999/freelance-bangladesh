import { makeHttpRequest } from "../../helpers/http-request";

export async function POST(req: any) {
    const url = `${process.env.API_URL}/api/v1/users/talent-signup`;
    const body = await req.json();

    return await makeHttpRequest(url, "POST", body);
}