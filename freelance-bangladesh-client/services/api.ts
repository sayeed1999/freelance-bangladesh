export const callApi = async ({
  url,
  method = "GET",
  body,
  accessToken
}: {
  url: string,
  method?: string,
  body?: any,
  accessToken?: string
}) => {
    const request: any = {
      method,
      headers: {
        "Content-Type": "application/json",
        "Authorization": accessToken ? `Bearer ${accessToken}` : "",
      },
    };

    if (method !== "GET") {
      request.body = JSON.stringify(body)
    }

    const resp = await fetch(url, request);
  
    const data = await resp.json();
  
    if (resp.ok) {
      return data;
    }
  
    console.error(data)
  
    throw new Error("Failed with error: " + data?.error);
}