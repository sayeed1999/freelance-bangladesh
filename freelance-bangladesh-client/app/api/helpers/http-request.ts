import { getServerSession } from "next-auth";
import { NextResponse } from "next/server";
import { authOptions } from "../auth/[...nextauth]/route";
import { getAccessToken } from "@/utils/sessionTokenAccessor";

export async function makeHttpRequest(
    url: string,
    method: "GET" | "POST" | "PATCH" | "PUT" | "DELETE" = "GET",
    body = {},
    headers = {}
) {
  try {
    const request: any = {
      headers: {
        "Content-Type": "application/json",
        ...headers,
      },
      method,
    };

    if (method !== "GET") {
      request.body = JSON.stringify(body);
    }

    const resp = await fetch(url, request);
    // If the response is not OK, throw an error with the status and error message
    if (!resp.ok) {
      const errorData = await resp.json();
      throw new Error(`Request failed with status ${resp.status}: ${errorData.message || "Unknown error"}`);
    }
    if (resp.status === 204) {
      return NextResponse.json(null, { status: 200 }); // Nextjs giving error with parsing req body for status 204 which should be null or not present
    }

    var data = await resp.json();
    return NextResponse.json(data, { status: resp.status });
  } catch (error: any) {
    console.error(`Error in HTTP request: ${error.message}`);
    return NextResponse.json({ error: error.message }, { status: 500 });
  }
}

export async function makeAuthorizedHttpRequest(
  url: string,
  method: "GET" | "POST" | "PATCH" | "PUT" | "DELETE" = "GET",
  body = {},
  headers = {},
) {
  const session = await getServerSession(authOptions);

  if (!session) {
    return NextResponse.json({ error: "Unauthorized" }, { status: 401 });
  }

  let accessToken = await getAccessToken();

  return await makeHttpRequest(url, method, body, 
  {
    "Content-Type": "application/json",
    Authorization: "Bearer " + accessToken,
    ...headers,
  });
}
