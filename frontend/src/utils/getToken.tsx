import { APIResponse } from "@/types/common";

export async function getToken() {
  const refreshToken = localStorage.getItem("token");
  if (!refreshToken) return null;

  let url = `${process.env.NEXT_PUBLIC_API_URL}/auth/refresh`;
  try {
    let response = await fetch(url, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        refreshToken,
      }),
    });
    let data: APIResponse = await response.json();
    if (data.status) {
        return data.data as string;
    } else {
        return null;
    }
  } catch (err) {
    return null;
  }
}
