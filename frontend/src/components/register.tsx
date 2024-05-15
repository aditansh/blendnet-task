"use client";

import React, { useState } from "react";
import { useFormik } from "formik";
import {
  Alert,
  Box,
  Button,
  CircularProgress,
  IconButton,
  InputAdornment,
  Link,
  MenuItem,
  Snackbar,
  Stack,
  TextField,
  Typography,
  useMediaQuery,
  useTheme,
} from "@mui/material";
import VisibilityIcon from "@mui/icons-material/Visibility";
import VisibilityOffIcon from "@mui/icons-material/VisibilityOff";
import * as yup from "yup";
import { useRouter } from "next/navigation";
import { APIResponse } from "@/types/common";

const registerSchema = yup.object({
  name: yup.string().required("Name is required"),
  email: yup.string().email("Invalid email").required("Email is required"),
  password: yup
    .string()
    .required("Password is required")
    .min(6, "Password must be at least 6 characters"),
  cpassword: yup
    .string()
    .required("Confirm Password is required")
    .oneOf([yup.ref("password")], "Passwords must match"),
});

function Register() {
  const router = useRouter();
  const theme = useTheme();
  const matches = useMediaQuery(theme.breakpoints.up("md"));
  const [error, setError] = useState("");
  const [message, setMessage] = useState("");
  const [open, setOpen] = useState(false);
  const [loading, setLoading] = useState(false);
  const [show, setShow] = useState(false);
  const [show2, setShow2] = useState(false);

  const handleClose = (
    event: React.SyntheticEvent | Event,
    reason?: string
  ) => {
    if (reason === "clickaway") {
      return;
    }

    setOpen(false);
  };

  //FORM HANDLING LIBRARY
  const formik = useFormik({
    initialValues: {
      name: "",
      email: "",
      password: "",
      cpassword: "",
    },
    validationSchema: registerSchema,
    validateOnBlur: true,
    validateOnChange: true,
    onSubmit: async (values) => {
      setMessage("");
      setError("");
      setLoading(true);
      let url = `${process.env.NEXT_PUBLIC_API_URL}/auth/register`;
      try {
        let response = await fetch(url, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            name: values.name,
            email: values.email,
            password: values.password,
          }),
        });
        let data: APIResponse = await response.json();
        if (data.status) {
          setLoading(false);
          setMessage("Account created successfully");
          setOpen(true);
          router.push("/login");
        } else {
          setLoading(false);
          setOpen(true);
          setError("Account creation failed");
        }
      } catch (err) {
        setLoading(false);
        setOpen(true);
        setError("Something went wrong");
      }
    },
  });

  const { values, errors, handleChange, handleSubmit } = formik;

  //to show and hide password
  const handleShowPassword = () => {
    setShow((prev) => !prev);
  };
  const handleShowPassword2 = () => {
    setShow2((prev) => !prev);
  };

  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        justifyContent: "center",
        height: "100vh",
        width: "100%",
      }}
    >
      <Box
        component="form"
        sx={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          justifyContent: "center",
          padding: matches ? 0 : "1.5rem",
          borderRadius: "1rem",
          height: "fit-content",
          width: matches ? "33%" : "100%",
          margin: "auto",
          boxSizing: "border-box",
        }}
        noValidate
        autoComplete="off"
        onSubmit={formik.handleSubmit}
      >
        {/* Form Heading */}
        <Box sx={{ mt: 1, mb: 2, mr: "auto" }}>
          <Typography
            sx={{
              textAlign: "left",
              fontSize: "30px",
              fontWeight: 700,
            }}
          >
            Hello there!
          </Typography>
          <Typography
            sx={{ textAlign: "left", fontSize: "18px", fontWeight: 600 }}
          >
            Create an account to get started
          </Typography>
        </Box>

        {/*First Name */}
        <TextField
          error={errors.name ? true : false}
          helperText={errors.name}
          required
          id="name"
          name="name"
          label="Name"
          placeholder="Name"
          onChange={handleChange}
          value={values.name}
          sx={{ width: "100%", mb: 2 }}
        />

        {/* Email ID */}
        <TextField
          error={errors.email ? true : false}
          helperText={errors.email}
          required
          id="email"
          name="email"
          label="Email ID"
          placeholder="Enter Email ID"
          onChange={handleChange}
          value={values.email}
          sx={{ width: "100%", mb: 2 }}
        />

        {/* Password */}
        <TextField
          error={errors.password ? true : false}
          helperText={errors.password}
          required
          id="password"
          name="password"
          label="Password"
          type={show ? "text" : "password"}
          placeholder="Password"
          onChange={handleChange}
          value={values.password}
          sx={{ width: "100%", mb: 2 }}
          InputProps={{
            endAdornment: (
              <InputAdornment position="end">
                <IconButton
                  aria-label="toggle password visibility"
                  onClick={handleShowPassword}
                  edge="end"
                >
                  {show ? <VisibilityIcon /> : <VisibilityOffIcon />}
                </IconButton>
              </InputAdornment>
            ),
          }}
        />

        {/* Confirm Password */}
        <TextField
          error={errors.cpassword ? true : false}
          helperText={errors.cpassword}
          required
          id="cpassword"
          name="cpassword"
          label="Confirm Password"
          type={show2 ? "text" : "password"}
          placeholder="Confirm Password"
          onChange={handleChange}
          value={values.cpassword}
          sx={{ width: "100%", mb: 2 }}
          InputProps={{
            endAdornment: (
              <InputAdornment position="end">
                <IconButton
                  aria-label="toggle cpassword visibility"
                  onClick={handleShowPassword2}
                  edge="end"
                >
                  {show2 ? <VisibilityIcon /> : <VisibilityOffIcon />}
                </IconButton>
              </InputAdornment>
            ),
          }}
        />

        {/* Submit and reset buttons */}
        <Button
          type="submit"
          variant="contained"
          size="medium"
          disabled={loading}
          sx={{ mb: 2 }}
        >
          {loading && (
            <>
              <CircularProgress thickness={6} color="inherit" size="1.2rem" />
              <Typography variant="subtitle2" style={{ marginLeft: "0.5rem" }}>
                Creating Account...
              </Typography>
            </>
          )}
          {!loading && "SUBMIT"}
        </Button>

        {/* Snackbar */}
        {!loading && (
          <Snackbar
            anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
            open={open}
            autoHideDuration={1000}
            onClose={handleClose}
          >
            <Alert
              onClose={handleClose}
              severity={error ? "error" : "success"}
              sx={{ width: "100%" }}
            >
              {message ? message : error}
            </Alert>
          </Snackbar>
        )}

        {/* Log in link */}
        <Link
          href="./login"
          variant="body2"
          sx={{
            fontWeight: 500,
            fontSize: "14px",
            textDecoration: "none",
          }}
        >
          Existing User? Log In
        </Link>
      </Box>
    </Box>
  );
}

export default Register;
