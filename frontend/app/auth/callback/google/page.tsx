"use client";

import { useEffect } from "react";
import { useRouter, useSearchParams } from "next/navigation";

export default function GoogleCallbackPage() {
  const searchParams = useSearchParams();

  useEffect(() => {
    const queryString = searchParams.toString();
    window.location.href = `http://localhost:8080/auth/callback/google?${queryString}`;
  }, [searchParams]);

  return (
    <div className="flex items-center justify-center min-h-screen bg-[#312244] text-white">
      <p>Finalizing login...</p>
    </div>
  );
}
