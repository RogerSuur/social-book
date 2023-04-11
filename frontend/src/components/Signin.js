import React from "react";
import { useState, useEffect } from "react";
import { Link, useNavigate, useLocation } from "react-router-dom";
import axios from "axios";
import useAuth from "../hooks/useAuth";

const LOGIN_URL = "http://localhost:8000/signin";

const Signin = () => {
  const { setAuth } = useAuth();

  const navigate = useNavigate();
  const location = useLocation();
  const from = location.state?.from?.pathname || "/posts";

  const [formData, setFormData] = useState({
    username: "",
    password: "",
  });
  const [errMsg, setErrMsg] = useState("");

  const handleChange = (event) => {
    const { name, value } = event.target;

    setFormData((prevFormData) => {
      return {
        ...prevFormData,
        [name]: value,
      };
    });
  };

  useEffect(() => {
    setErrMsg("");
  }, [formData]);

  const handleSubmit = async (event) => {
    event.preventDefault();

    try {
      const response = await axios.post(LOGIN_URL, JSON.stringify(formData), {
        headers: { "Content-Type": "application/json" },
        withCredentials: true,
      });

      console.log(JSON.stringify(response));
      setAuth(true);
      setFormData({
        username: "",
        password: "",
      });

      console.log(from, "FROM");

      navigate(from, { replace: true });
    } catch (err) {
      if (!err?.response) {
        setErrMsg("No Server Response");
      } else if (err.response?.status === 400) {
        setErrMsg("Missing Username or Password");
      } else if (err.response?.status === 401) {
        setErrMsg("Wrong username or password");
      } else {
        setErrMsg("Login Failed");
      }
    }
  };

  return (
    <>
      {errMsg && <h2>{errMsg}</h2>}
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="Email or username"
          onChange={handleChange}
          name="username"
          value={formData.username}
          required
          autoFocus
        />
        <input
          type="password"
          placeholder="Password"
          onChange={handleChange}
          name="password"
          value={formData.password}
          required
        />
        <button>Sign In</button>
      </form>
      <div>
        Do not have an account? <Link to={`/signup`}>Sign up</Link>
      </div>
    </>
  );
};

export default Signin;
