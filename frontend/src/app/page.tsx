"use client";

import React, { useEffect } from "react";
import Register from "@/components/register";
import { useRouter } from "next/navigation";
import { getToken } from "@/utils/getToken";

export default function Page() {
  const router = useRouter();

  useEffect(() => {
    async function f() {
      if (localStorage.getItem("token")) {
        const token = await getToken();
        if (token) router.push("/dashboard");
      }
    }

    f();
  }, []);

  return <Register />;
}
