import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import axios from "axios";
import { useForm } from "react-hook-form";

const SIGNUP_URL = "http://localhost:8000/signup";

const Signup = () => {
  const [errMsg, setErrMsg] = useState("");
  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm({
    defaultValues: {
      dateOfBirth: new Date(
        new Date().getFullYear() - 13,
        new Date().getMonth(),
        new Date().getDate()
      )
        .toISOString()
        .split("T")[0],
    },
    criteriaMode: "all",
  });

  // const onSubmit = (data) => {
  //   console.log(data, "here");
  // };

  // const [formData, setFormData] = useState({
  //   email: "",
  //   password: "",
  //   firstName: "",
  //   lastName: "",
  //   dateOfBirth: new Date().toISOString().split("T")[0],
  //   nickname: "",
  //   about: "",
  // });

  const navigate = useNavigate();

  // const handleChange = (event) => {
  //   const { name, value, type, checked } = event.target;

  //   setFormData((prevFormData) => {
  //     return {
  //       ...prevFormData,
  //       [name]: type === "checkbox" ? checked : value,
  //     };
  //   });
  // };

  const onSubmit = async (data) => {
    try {
      const response = await axios.post(SIGNUP_URL, JSON.stringify(data), {
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
      <form onSubmit={handleSubmit(onSubmit)}>
        <input
          placeholder="First Name"
          {...register("firstName", {
            required: "Please enter your first name",
          })}
        />
        {errors.firstName && <p>{errors.firstName.message}</p>}
        <br />
        <input
          placeholder="Last Name"
          {...register("lastName", {
            required: "Please enter your last name",
          })}
        />
        {errors.lastName && <p>{errors.lastName.message}</p>}

        <br />
        <input
          placeholder="Email address"
          {...register("email", {
            required: "Please enter your email address",
            pattern: {
              value:
                /^[A-Z0-9][A-Z0-9._%+-]{0,63}@(?:[A-Z0-9-]{1,63}\.){1,15}[A-Z]{2,63}$/i,
              message:
                "The email address should be in form of example@example.com",
            },
          })}
        />
        {errors.email && <p>{errors.email.message}</p>}

        <br />
        <input
          type="password"
          placeholder="Password"
          {...register("password", {
            required: "Please enter your email address",
            minLength: {
              value: 8,
              message: "The password should be at least 8 characters long",
            },
            pattern: {
              value: /^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#$%^&*])/,
              message:
                "The password should have at least one lowercase and one uppercase letter, a number and a symbol",
            },
          })}
        />
        {errors.password && <p>{errors.password.message}</p>}
        <br />
        <input
          type="password"
          placeholder="Confirm password"
          {...register("confirm_password", {
            required: "Please enter your password again",
            validate: (value) =>
              value === watch("password") || "The passwords do not match",
          })}
        />
        {errors.confirm_password && <p>{errors.confirm_password.message}</p>}
        <br />
        <input
          type="date"
          {...register("dateOfBirth", {
            required: "Please enter your birth date",
            validate: (value) =>
              new Date(value) <
                new Date(
                  new Date().getFullYear() - 13,
                  new Date().getMonth(),
                  new Date().getDate()
                ) || "You must be 13 years of age or older to sign up",
          })}
        />
        {errors.dateOfBirth && <p>{errors.dateOfBirth.message}</p>}
        <br />
        <input
          placeholder="Nickname"
          {...register("nickname", {
            minLength: {
              value: 8,
              message: "The nickname should be at least 8 characters long",
            },
            pattern: {
              value: /^[a-zA-Z0-9._]{8,}$/,
              message:
                "The nickname can only contain letters, numbers, dots (.) and underscores (_)",
            },
          })}
        />
        {errors.nickname && <p>{errors.nickname.message}</p>}
        <br />
        <textarea
          placeholder="Write something about yourself"
          {...register("about")}
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
