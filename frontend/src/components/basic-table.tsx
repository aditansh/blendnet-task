"use client";

import * as React from "react";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Paper from "@mui/material/Paper";
import { Delete, Visibility } from "@mui/icons-material";
import { useTheme } from "@mui/material";

export default function BasicTable({
  stocks,
  delete: deleteStock,
  view: viewStock,
}: {
  stocks: string[];
  delete: (key: string) => void;
  view: (key: string) => void;
}) {
  const theme = useTheme();
  const matches = theme.breakpoints.up("md");

  return (
    <TableContainer
      component={Paper}
      sx={{ display: "flex", justifyContent: "center" }}
    >
      <Table sx={{ width: matches ? "50%" : "100%" }} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>Stock Name</TableCell>
            <TableCell align="right"></TableCell>
            <TableCell align="right"></TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {stocks.map((stock) => (
            <TableRow
              key={stock}
              sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
            >
              <TableCell component="th" scope="row">
                {stock}
              </TableCell>
              <TableCell
                align="right"
                sx={{ cursor: "pointer" }}
                onClick={() => viewStock(stock)}
              >
                <Visibility />
              </TableCell>
              <TableCell
                align="right"
                sx={{ cursor: "pointer" }}
                onClick={() => deleteStock(stock)}
              >
                <Delete />
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}
