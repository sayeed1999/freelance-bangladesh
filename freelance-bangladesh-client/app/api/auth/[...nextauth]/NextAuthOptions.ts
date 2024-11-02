import KeycloakProvider from "next-auth/providers/keycloak";
import { jwtDecode } from "jwt-decode";
import { encrypt } from "@/utils/encryption";

function requestRefreshOfAccessToken(token: any) {
  return fetch(`${process.env.KEYCLOAK_ISSUER}/protocol/openid-connect/token`, {
    headers: { "Content-Type": "application/x-www-form-urlencoded" },
    body: new URLSearchParams({
      client_id: process.env.KEYCLOAK_CLIENT_ID || "",
      client_secret: process.env.KEYCLOAK_CLIENT_SECRET || "",
      grant_type: "refresh_token",
      refresh_token: token.refresh_token || "",
    }),
    method: "POST",
    cache: "no-store"
  });
}

export const authOptions = {
    providers: [
      KeycloakProvider({
        clientId: `${process.env.KEYCLOAK_CLIENT_ID}`,
        clientSecret: `${process.env.KEYCLOAK_CLIENT_SECRET}`,
        issuer: `${process.env.KEYCLOAK_ISSUER}`,
      }),
    ],
  
    session: {
      maxAge: 60 * 30 // default lifespan of refresh token
    },
  
    callbacks: {
      async jwt({ token, account }: any) {
        if (account) {
          // account is only available the first time this callback is called on a new session (after the user signs in)
          token.decoded = jwtDecode(account.access_token);
          token.access_token = account.access_token;
          token.id_token = account.id_token;
          token.expires_at = account.expires_at;
          token.refresh_token = account.refresh_token;
          return token;
        }
  
        // we take a buffer of one minute(60 * 1000 ms)
        if (Date.now() < (token.expires_at * 1000 - 60 * 1000)) {
          return token
        } else {
          try {
            const response = await requestRefreshOfAccessToken(token)
  
            const tokens = await response.json()
  
            if (!response.ok) throw tokens
  
            const updatedToken = {
              ...token, // Keep the previous token properties
              // @ts-expect-error // TODO: why refresh token doesnt exist?
              decoded: jwtDecode(refreshToken.access_token),
              id_token: tokens.id_token,
              access_token: tokens.access_token,
              expires_at: Math.floor(Date.now() / 1000 + (tokens.expires_in)),
              refresh_token: tokens.refresh_token ?? token.refresh_token,
            }
            return updatedToken
          } catch (error) {
            console.error("Error refreshing access token", error)
            return { ...token, error: "RefreshAccessTokenError" }
          }
        }
      },
      async session({ session, token }: any) {
        // Send properties to the client
        session.access_token = encrypt(token.access_token); // see utils/sessionTokenAccessor.js
        session.id_token = encrypt(token.id_token);  // see utils/sessionTokenAccessor.js
        session.roles = token.decoded.realm_access.roles;
        session.error = token.error;      
        return session;
      },
    },
  };
  