import React from "react";
import { useState, useEffect } from "react";
import { Link, useNavigate, useLocation } from "react-router-dom";
import axios from "axios";
import useAuth from "../hooks/useAuth";

const LOGIN_URL = "http://localhost:8000/login";

const Login = () => {
  const { setAuth } = useAuth();

  const navigate = useNavigate();
  const location = useLocation();
  const from =
    location.state?.from?.pathname !== "/logout"
      ? location.state?.from?.pathname || "/profile"
      : "/profile";

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

      console.log(JSON.stringify(response, null, 2));
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
      <div style={{ display: "flex", justifyContent: "center", alignItems: "center", height: "100vh" }}>
  <div style={{ width: "500px", height: "450px", border: "1px solid #ccc", borderRadius: "20px", padding: "20px" }}>
    <form onSubmit={handleSubmit}>
    <h1 >Sign in</h1>
    
      <label className="log-in">Email or username</label>
      
      <input  className= "login-input"
        type="text"
        placeholder="Email or username"
        onChange={handleChange}
        name="username"
        value={formData.username}
        required
        autoFocus
        
      />
      <br />
       <label  className="log-in">Password</label>
      <input  className= "login-input"
        type="password"
        placeholder="Password"
        onChange={handleChange}
        name="password"
        value={formData.password}
        required
       
      />
      
      <div className="center1">
      <button className = "log-button">
        Sign In
      </button>
      </div>
    </form>
    <div style={{ marginTop: "10px", textAlign: "center" }}>
      Do not have an account? <Link to={`/signup`}>Sign up</Link>
    </div>
  </div>
</div>
    </>
  );
};

export default Login;
