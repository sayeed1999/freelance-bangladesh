import { NextResponse } from "next/server";
import { getServerSession } from "next-auth";
import { getAccessToken } from "@/utils/sessionTokenAccessor";
import { authOptions } from "../auth/[...nextauth]/route";

export async function POST(req) {
  const session = await getServerSession(authOptions);

  if (session) {
    const url = `${process.env.API_URL}/api/v1/jobs`;

    const postBody = await req.json();
    let accessToken = await getAccessToken();

    const resp = await fetch(url, {
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + accessToken,
      },
      method: "POST",
      body: JSON.stringify(postBody),
    });

    const data = await resp.json();
    return NextResponse.json(data, { status: resp.status });
  }

  return NextResponse.json({ error: "Unauthorized" }, { status: res.status });
}