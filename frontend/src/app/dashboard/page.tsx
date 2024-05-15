"use client";

import BasicTable from "@/components/basic-table";
import { APIResponse, StockData, User } from "@/types/common";
import { getToken } from "@/utils/getToken";
import { useRouter } from "next/navigation";
import React, { useEffect } from "react";
import StockDataTable from "@/components/stock-data";
import SearchBar from "@/components/search";

export default function Page() {
  const router = useRouter();
  const [stocks, setStocks] = React.useState<string[]>([]);
  const [stock, setStock] = React.useState<StockData | null>(null);

  const getProfile = async () => {
    let url = `${process.env.NEXT_PUBLIC_API_URL}/user/me`;
    try {
      let token: string | null = "";
      if (localStorage.getItem("token")) {
        token = await getToken();
        if (!token) router.push("/login");
      }
      let response = await fetch(url, {
        method: "GET",
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      let data: APIResponse = await response.json();
      if (data.status) {
        const user = data.data as User;
        setStocks(user.watchlist);
      } else {
        router.push("/login");
      }
    } catch (err) {
      router.push("/login");
    }
  };

  const deleteStock = async (key: string) => {
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
          method: "delete",
        }),
      });
      let data: APIResponse = await response.json();
      if (data.status) {
        const search = data.data as string[];
        setStocks(search);
      } else {
        setStocks([]);
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
    getProfile();
  }, []);

  return (
    <div>
      {!stock && (
        <>
          <SearchBar />
          <BasicTable stocks={stocks} delete={deleteStock} view={viewStock} />
        </>
      )}
      {stock && <StockDataTable stock={stock} back={() => setStock(null)} />}
    </div>
  );
}
