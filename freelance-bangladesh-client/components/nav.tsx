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
  },
  {
    name: "Talent Signup",
    path: "/signup/talent",
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
  },
];

export default function Nav() {
  return (
    <ul className="mt-3">
      {navItems.map((item, index) => (
        <li key={index} className="my-2">
          <NavLink href={item.path}>{item.name}</NavLink>
        </li>
      ))}
    </ul>
  );
}
