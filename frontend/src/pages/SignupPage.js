import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import axios from "axios";

const SIGNUP_URL = "http://localhost:8000/signup";

const Signup = () => {
  const [errMsg, setErrMsg] = useState("");
  const [formData, setFormData] = useState({
    email: "",
    password: "",
    firstName: "",
    lastName: "",
    dateOfBirth: new Date().toISOString().split("T")[0],
    nickname: "",
    about: "",
  });

  const navigate = useNavigate();

  const handleChange = (event) => {
    const { name, value, type, checked } = event.target;

    setFormData((prevFormData) => {
      return {
        ...prevFormData,
        [name]: type === "checkbox" ? checked : value,
      };
    });
  };

  const handleSubmit = async (event) => {
    event.preventDefault();

    try {
      const response = await axios.post(SIGNUP_URL, JSON.stringify(formData), {
        headers: { "Content-Type": "application/json" },
        withCredentials: true,
      });

      console.log(JSON.stringify(response));

      navigate("/login", { replace: true });
    } catch (err) {
      if (!err?.response) {
        setErrMsg("No Server Response");
      } else if (err.response?.status === 400) {
        const data = err.response.data.slice(0, -1);
        if (data === "nickname") {
          setErrMsg(
            "A username should contain only uppercase and lowercase letters, dots (.) or underscores(_). If it fits the description, the username has already been taken"
          );
        } else if (data === "email") {
          setErrMsg(
            "An email address should be in the form of example@example.com and contain only uppercase and lowercase letters, and symbols(. _%+-). If it fits the description, the email has already been taken"
          );
        } else if (data === "password") {
          setErrMsg(
            "Your password should have at least one lowercase and one uppercase letter, a number and a symbol"
          );
        }
      } else {
        setErrMsg("Internal Server Error");
      }
    }
  };

  return (
    <>
      {errMsg && <h3>{errMsg}</h3>}
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="First Name"
          onChange={handleChange}
          name="firstName"
          value={formData.firstName}
          required
        />
        <br />
        <input
          type="text"
          placeholder="Last Name"
          onChange={handleChange}
          name="lastName"
          value={formData.lastName}
          required
        />
        <br />
        <input
          type="email"
          placeholder="Email address"
          onChange={handleChange}
          name="email"
          value={formData.email}
          title="Email in the form of example@example.com"
          required
        />
        <br />
        <input
          type="password"
          placeholder="Password"
          onChange={handleChange}
          name="password"
          value={formData.password}
          title="Password should have at least one lowercase and one uppercase letter, a number and a symbol"
          minLength="8"
          maxLength="32"
          required
        />
        <br />
        <input
          type="date"
          placeholder="Date of Birth"
          onChange={handleChange}
          name="dateOfBirth"
          value={formData.dateOfBirth}
          required
        />
        <br />
        <input
          type="text"
          placeholder="Username"
          onChange={handleChange}
          name="nickname"
          value={formData.nickname}
          minLength="8"
          maxLength="30"
        />
        <br />
        <textarea
          placeholder="Write something about yourself"
          onChange={handleChange}
          value={formData.about}
          name="about"
        />
        <br />
        <button>Sign Up</button>
      </form>
      <div>
        Already have an account? <Link to={`/login`}>Sign in</Link>
      </div>
    </>
  );
};

export default Signup;
