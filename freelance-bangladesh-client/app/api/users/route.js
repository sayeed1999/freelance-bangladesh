import { NextResponse } from "next/server";

export async function POST(req) {
    const url = `${process.env.API_URL}/api/v1/users`;

    const postBody = await req.json();

    const resp = await fetch(url, {
        headers: {
        "Content-Type": "application/json",
        },
        method: "POST",
        body: JSON.stringify(postBody),
    });

    const data = await resp.json();
    return NextResponse.json(data, { status: resp.status });
}