import Link from 'next/link';

export default function Unauthorized() {
  return (
    <main className="h-screen flex flex-col items-center justify-center bg-gradient-to-r from-blue-400 to-purple-600 text-white">
      <div className="text-center">
        <h1 className="text-6xl font-bold mb-4">403 - Forbidden!</h1>
        <p className="text-xl mb-6">Oops!🫢 You don't have permission to access this page.</p>

        <div className="flex justify-center space-x-4">
          <Link href="/" className="px-6 py-3 bg-white text-blue-600 font-semibold rounded-lg shadow-lg hover:bg-gray-100 transition duration-300">
              Go to Homepage
          </Link>

        </div>
      </div>
      
      <div className="absolute bottom-0 w-full h-64 bg-bottom bg-no-repeat bg-cover" style={{ backgroundImage: "url('/images/unauthorized-bg.svg')" }}>
      </div>
    </main>
  );
}
