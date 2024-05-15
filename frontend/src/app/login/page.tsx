"use client";

import React, { useEffect } from "react";
import { useRouter } from "next/navigation";
import { getToken } from "@/utils/getToken";
import Login from "@/components/login";

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

  return <Login />;
}
