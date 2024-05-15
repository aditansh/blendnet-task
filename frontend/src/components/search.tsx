"use client";

import {
  Autocomplete,
  InputBase,
  TextField,
  alpha,
  styled,
} from "@mui/material";
import SearchIcon from "@mui/icons-material/Search";
import React, { useEffect, useState } from "react";
import { APIResponse, SearchResponse } from "@/types/common";
import { getToken } from "@/utils/getToken";
import { useRouter } from "next/navigation";

const Search = styled("div")(({ theme }) => ({
  position: "relative",
  borderRadius: theme.shape.borderRadius,
  backgroundColor: alpha(theme.palette.common.white, 0.15),
  "&:hover": {
    backgroundColor: alpha(theme.palette.common.white, 0.25),
  },
  marginLeft: 0,
  width: "100%",
  [theme.breakpoints.up("md")]: {
    width: "50%",
  },
}));

const SearchIconWrapper = styled("div")(({ theme }) => ({
  padding: theme.spacing(0, 2),
  height: "100%",
  position: "absolute",
  pointerEvents: "none",
  display: "flex",
  alignItems: "center",
  justifyContent: "center",
}));

const StyledInputBase = styled(TextField)(({ theme }) => ({
  color: "inherit",
  width: "100%",
  "& .MuiInputBase-input": {
    padding: "theme.spacing(1, 1, 1, 0) !important",
    // vertical padding + font size from searchIcon
    paddingLeft: `calc(1em + ${theme.spacing(4)}) !important`,
  },
}));

const SearchBar = () => {
  const router = useRouter();
  const [inputValue, setinputValue] = useState("");
  const [values, setValues] = useState<string[]>([]);
  const [stock, setStock] = useState<string | null>(null);

  const getSearch = async () => {
    let url = `${process.env.NEXT_PUBLIC_API_URL}/user/search`;
    try {
      let token: string | null = "";
      if (localStorage.getItem("token")) {
        token = await getToken();
        if (!token) router.push("/login");
      }
      let response = await fetch(url, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
          key: inputValue,
        }),
      });
      let data: APIResponse = await response.json();
      if (data.status) {
        const search = data.data as SearchResponse;
        setValues([]);
        for (let i = 0; i < search.bestMatches.length; i++) {
          const curr = search.bestMatches[i];
          if (curr) {
            const temp = `${curr["1. symbol"]} - ${curr["2. name"]}`;
            setValues((prev) => [...prev, temp]);
          }
        }
      } else {
      }
    } catch (err) {
      console.log(err);
    }
  };

  useEffect(() => {
    getSearch();
  }, [inputValue]);

  useEffect(() => {
    if (stock) {
      router.push(`/stock?symbol=${stock?.split(" - ")[0]}`);
    }
  }, [stock]);

  return (
    <div>
      <Autocomplete
        inputValue={inputValue}
        onInputChange={(event, newInputValue) => {
          setinputValue(newInputValue);
        }}
        value={stock}
        onChange={(event, newValue) => {
          setStock(newValue);
        }}
        id="combo-box-demo"
        options={values}
        getOptionLabel={(option) => option}
        sx={{ width: "100%", my: "20px", mx: "auto", display: "flex" }}
        renderInput={(params) => (
          <Search sx={{ mx: "auto" }}>
            <SearchIconWrapper>
              <SearchIcon />
            </SearchIconWrapper>
            <StyledInputBase {...params} variant="outlined" />
          </Search>
        )}
        open={inputValue.length > 0}
      />
    </div>
  );
};

export default SearchBar;
