import { getServerSession } from "next-auth"
import { getIdToken } from "@/utils/sessionTokenAccessor";
import { authOptions } from "../[...nextauth]/NextAuthOptions";

export async function GET() {
  const session = await getServerSession(authOptions);

  if (session) {

    const idToken = await getIdToken();

    // this will log out the user on Keycloak side
    var baseURL = `${process.env.KEYCLOAK_ISSUER}/protocol/openid-connect/logout`;
    var url = `${baseURL}?id_token_hint=${idToken}&post_logout_redirect_uri=${encodeURIComponent(process.env.NEXTAUTH_URL)}`;

    try {
      const resp = await fetch(url, { method: "GET" });
    } catch (err) {
      console.error(err);
      return new Response({ status: 500 });
    }
  }

  return new Response({ status: 200 });
}
