import Link from 'next/link';

export default function Nav() {
  return (
    <ul className="mt-3">
      <li className="my-1"><Link className="hover:bg-gray-500" href="/signup/talent">Talent Signup</Link></li>
      <li className="my-1"><Link className="hover:bg-gray-500" href="/">Home</Link></li>
      <li className="my-1"><Link className="hover:bg-gray-500" href="/jobs">See Jobs</Link></li>
      <li className="my-1"><Link className="hover:bg-gray-500" href="/jobs/create">Create Job</Link></li>
    </ul>
  );
}
