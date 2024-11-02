import NextAuth from "next-auth";
import { authOptions } from "./NextAuthOptions";

const handler = NextAuth(authOptions);

export { handler as GET, handler as POST };