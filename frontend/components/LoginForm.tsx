"use client";

import Image from "next/image";

export default function LoginForm() {
  const handleLogin = () => {
    const apiUrl = process.env.NEXT_PUBLIC_API_URL;
    window.location.href = `${apiUrl}/auth/login/google`;
  };

  return (
    <div className="flex items-center justify-center w-1/4 h-1/3 bg-[#006466] text-black">
      <button
        onClick={handleLogin}
        className="flex justify-evenly items-center w-50 h-8 bg-white"
      >
        <Image
          src="https://www.gstatic.com/firebasejs/ui/2.0.0/images/auth/google.svg"
          alt="Google logo"
          width={24}
          height={24}
        />
        Login with Google
      </button>
    </div>
  );
}
