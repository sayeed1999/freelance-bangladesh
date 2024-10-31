import { authOptions } from "@/app/api/auth/[...nextauth]/route";
import { getServerSession } from "next-auth";
import Link from "next/link";
import { ReactNode } from "react";

interface NavLinkProps {
  href: string;
  children: ReactNode;
}

function NavLink({ href, children }: NavLinkProps) {
  return (
    <Link
      href={href}
      className="block px-4 py-2 rounded-lg text-white bg-blue-600 hover:bg-blue-700 transition-colors duration-300 shadow-md"
    >
      {children}
    </Link>
  );
}

const navItems = [
  {
    name: "Client Signup",
    path: "/signup/client",
    requireRole: "admin",
  },
  {
    name: "Talent Signup",
    path: "/signup/talent",
    publicRoute: true,
  },
  {
    name: "Client List",
    path: "/admin-dashboard/clients",
    requireRole: "admin",
  },
  {
    name: "Talent List",
    path: "/admin-dashboard/talents",
    requireRole: "admin",
  },
  {
    name: "Home",
    path: "/",
  },
  {
    name: "See Jobs",
    path: "/jobs",
  },
  {
    name: "Create Job",
    path: "/jobs/create",
    requireRole: "client",
  },
  {
    name: "Assignment List",
    path: "/assignments",
  },
];

export default async function Nav() {
  const session = await getServerSession(authOptions);

  return (
    <ul className="mt-3">
      {navItems
        .filter((item) => {
          if (item.publicRoute) return !session; // skip public routes for logged in state
          if (session && !item.requireRole) return true;
          return session?.roles?.includes(item.requireRole);
        })
        .map((item, index) => (
          <li key={index} className="my-2">
            <NavLink href={item.path}>{item.name}</NavLink>
          </li>
        ))}
    </ul>
  );
}
