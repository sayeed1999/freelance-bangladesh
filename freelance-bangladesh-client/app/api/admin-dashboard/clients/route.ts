import { makeAuthorizedHttpRequest } from "../../helpers/http-request";

export async function GET() {
  const url = `${process.env.API_URL}/api/v1/admin-dashboard/clients`;
  return await makeAuthorizedHttpRequest(url);
}

export async function POST(req: any) {

  const url = `${process.env.API_URL}/api/v1/admin-dashboard/clients`;
  const body = await req.json();

  return await makeAuthorizedHttpRequest(url, "POST", body);
}
