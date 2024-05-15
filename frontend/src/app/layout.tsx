"use client";

import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "@/styles/globals.css";
import { CssBaseline, ThemeProvider, createTheme } from "@mui/material";

const inter = Inter({ subsets: ["latin"] });

// export const metadata: Metadata = {
//   title: "Blendnet - Task",
//   description: "Blendnet - Task",
// };

const darkTheme = createTheme({
  palette: {
    mode: "dark",
  },
});

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <ThemeProvider theme={darkTheme}>
          <CssBaseline />
          {children}
        </ThemeProvider>
      </body>
    </html>
  );
}
