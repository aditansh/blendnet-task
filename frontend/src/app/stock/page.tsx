"use client";

import { APIResponse, StockData, User } from "@/types/common";
import { getToken } from "@/utils/getToken";
import { useRouter, useSearchParams } from "next/navigation";
import React, { Suspense, useEffect, useState } from "react";
import StockDataTable from "@/components/stock-data";
import { Box, Button, Stack } from "@mui/material";

function Temp() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const symbol = searchParams.get("symbol");
  const [stock, setStock] = useState<StockData | null>(null);

  const addStock = async (key: string) => {
    let url = `${process.env.NEXT_PUBLIC_API_URL}/user/watchlist`;
    try {
      let token: string | null = "";
      if (localStorage.getItem("token")) {
        token = await getToken();
        if (!token) router.push("/login");
      }
      let response = await fetch(url, {
        method: "PATCH",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
          watchlist: [key],
          method: "add",
        }),
      });
      let data: APIResponse = await response.json();
      if (data.status) {
        router.push("/dashboard");
      } else {
        setStock(stock);
      }
    } catch (err) {
      console.log(err);
    }
  };

  const viewStock = async (key: string) => {
    let url = `${
      process.env.NEXT_PUBLIC_API_URL
    }/user/stock?symbol=${key.toUpperCase()}&interval=5min`;
    try {
      let token: string | null = "";
      if (localStorage.getItem("token")) {
        token = await getToken();
        if (!token) router.push("/login");
      }
      let response = await fetch(url, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
      });
      let data: APIResponse = await response.json();
      if (data.status) {
        const stock = data.data as StockData;
        setStock(stock);
      } else {
        setStock(null);
      }
    } catch (err) {
      console.log(err);
    }
  };

  useEffect(() => {
    if (!symbol) {
      router.push("/dashboard");
      return;
    }
    viewStock(symbol);
  }, []);

  return (
    <Box>
      <Stack spacing={2} direction="row" sx={{ m: 2 }}>
        <Button onClick={() => router.push("/dashboard")} variant="outlined">
          Back
        </Button>

        <Button onClick={() => addStock(symbol ?? "")} variant="outlined">
          Add
        </Button>
      </Stack>
      {stock && <StockDataTable stock={stock} />}
    </Box>
  );
}

export default function Page() {
  return (
    <Suspense>
      <Temp />;
    </Suspense>
  );
}
