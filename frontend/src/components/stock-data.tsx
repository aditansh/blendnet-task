import { StockData } from "@/types/common";
import { Box, Button, Paper, useTheme } from "@mui/material";
import React from "react";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";

export default function StockDataTable({
  stock,
  back,
}: {
  stock: StockData;
  back?: () => void;
}) {
  const theme = useTheme();
  const matches = theme.breakpoints.up("md");

  return (
    <Box sx={{}}>
     {back && <Button onClick={back} variant="outlined" sx={{ m: 2 }}>
        Back
      </Button>}
      <TableContainer
        component={Paper}
        sx={{
          display: "flex",
        }}
      >
        <Table sx={{ width: "100%" }} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell>Time</TableCell>
              <TableCell>Open</TableCell>
              <TableCell>High</TableCell>
              <TableCell>Low</TableCell>
              <TableCell>Close</TableCell>
              <TableCell>Volume</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {Object.keys(stock["Time Series (5min)"]).map((key, i) => {
              const data = stock["Time Series (5min)"][key];
              if (!data) return null;
              return (
                <TableRow
                  key={i}
                  sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
                >
                  <TableCell component="th" scope="row">
                    {key}
                  </TableCell>
                  <TableCell component="th" scope="row">
                    {data["1. open"]}
                  </TableCell>
                  <TableCell>{data["2. high"]}</TableCell>
                  <TableCell>{data["3. low"]}</TableCell>
                  <TableCell>{data["4. close"]}</TableCell>
                  <TableCell>{data["5. volume"]}</TableCell>
                </TableRow>
              );
            })}
          </TableBody>
        </Table>
      </TableContainer>
    </Box>
  );
}
