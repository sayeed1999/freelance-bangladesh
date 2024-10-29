import { makeAuthorizedHttpRequest } from "../../../helpers/http-request";

export async function POST(req: any) {
  const urlParts = req.url.split('/');
  const jobid = urlParts[urlParts.length - 2];

  const url = `${process.env.API_URL}/api/v1/jobs/${jobid}/bids`;
  const body = await req.json();

  return await makeAuthorizedHttpRequest(url, "POST", body);
}

export async function GET(req: any) {
  const urlParts = req.url.split('/');
  const jobid = urlParts[urlParts.length - 2];

  const url = `${process.env.API_URL}/api/v1/jobs/${jobid}/bids`;

  return await makeAuthorizedHttpRequest(url);
}