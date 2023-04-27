import React, { useState, useEffect } from "react";
import { useForm } from "react-hook-form";
import axios from "axios";
import Modal from "../components/Modal.js";
import AvatarUpdater from "../components/AvatarUpdater.js";
import FileUploader from "./FileUploader.js";
import UploadAndDisplayImage from "./UploadAndDisplayImage.js";

const PROFILE_URL = "http://localhost:8000/profile";
const PROFILE_UPDATE_URL = "";

const ProfileInfo = (props) => {
  const [user, setUser] = useState({});
  const [errMsg, setErrMsg] = useState("");
  const values = user;
  const {
    register,
    handleSubmit,
    watch,
    formState: { errors, isDirty },
  } = useForm({
    mode: "onBlur",
    values,
    criteriaMode: "all",
  });

  useEffect(() => {
    const loadUser = async () => {
      await axios
        .get(PROFILE_URL, {
          withCredentials: true,
        })
        .then((response) => setUser(response.data.user));
    };
    loadUser();
  }, []);

  const onSubmit = async (data) => {
    console.log(data);
    // try {
    //   const response = await axios.post(
    //     PROFILE_UPDATE_URL,
    //     JSON.stringify(data),
    //     {
    //       headers: { "Content-Type": "application/json" },
    //       withCredentials: true,
    //     }
    //   );

    //   console.log(JSON.stringify(response));

    //   navigate("/login", { replace: true });
    // } catch (err) {
    //   if (!err?.response) {
    //     setErrMsg("No Server Response");
    //   } else if (err.response?.status === 400) {
    //     const data = err.response.data.slice(0, -1);
    //     if (data === "nickname") {
    //       setErrMsg("The nickname has already been taken");
    //     } else if (data === "email") {
    //       setErrMsg("Please use another email address");
    //     } else if (data === "password") {
    //       setErrMsg(
    //         "Your password should have at least one lowercase and one uppercase letter, a number and a symbol"
    //       );
    //     }
    //   } else {
    //     setErrMsg("Internal Server Error");
    //   }
    // }
  };

  // console.log(isDirty, "DIRTY");
  // console.log(user, "USER");

  return (
    <>
      {user && (
        <div className="content-area">
          <form onSubmit={handleSubmit(onSubmit)}>
            <div className="row">
              <div className="column">
                <img
                  style={{
                    width: "20vw",
                    height: "20vw",
                    objectFit: "cover",
                    objectPosition: "0% 100%",
                  }}
                  src={
                    "https://www.pixelstalk.net/wp-content/uploads/2016/05/Free-Cool-Backgrounds.jpg"
                  }
                  alt={`${user.firstName}'s image`}
                ></img>
              </div>

              <h1 className="column-title">{user.firstName}'s profile</h1>
            </div>
            {/* <Modal text={"Update Image"}>
              <AvatarUpdater />
            </Modal> */}
            <Modal text={"Upload New Image"}>
              <AvatarUpdater />
            </Modal>
            <div className="row">
              <div className="column-title">First Name</div>
              <div className="column">{user.firstName}</div>
            </div>
            <div className="row">
              <div className="column-title">Last Name</div>
              <div className="column">{user.lastName}</div>
            </div>
            <div className="row">
              <div className="column-title">Email</div>
              <div className="column">{user.email}</div>
            </div>
            {/* <div className="row">
              <div className="column-title">New Password</div>
              <div className="column">
                <input
                  type="password"
                  placeholder="Password"
                  {...register("password", {
                    required: "Please enter your password",
                    minLength: {
                      value: 8,
                      message:
                        "The password should be at least 8 characters long",
                    },
                    pattern: {
                      value:
                        /^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#$%^&*])/,
                      message:
                        "The password should have at least one lowercase and one uppercase letter, a number and a symbol",
                    },
                  })}
                />
                {errors.password && <p>{errors.password.message}</p>}
              </div>
            </div>
            <div className="row">
              <div className="column-title">Confirm Password</div>
              <div className="column">
                <input
                  type="password"
                  placeholder="Confirm password"
                  {...register("confirmPassword", {
                    exclude: true,
                    required: "Please enter your password again",
                    validate: (value) =>
                      value === watch("password") ||
                      "The passwords do not match",
                  })}
                />
                {errors.confirmPassword && (
                  <p>{errors.confirmPassword.message}</p>
                )}
              </div>
            </div> */}
            <div className="row">
              <div className="column-title">Birthday</div>
              <div className="column">{user.birthday}</div>
            </div>
            <div className="row">
              <div className="column-title">Nickname</div>
              <div className="column">
                <input
                  placeholder="Nickname"
                  {...register("nickname", {
                    maxLength: {
                      value: 32,
                      message:
                        "The nickname should not be longer than 32 characters long",
                    },
                    pattern: {
                      value: /^[a-zA-Z0-9._ ]{0,32}$/,
                      message:
                        "The nickname can only contain letters, numbers, spaces, dots (.) and underscores (_)",
                    },
                  })}
                />
                {errors.nickname && <p>{errors.nickname.message}</p>}
                <br />
              </div>
            </div>
            <div className="row">
              <div className="column-title">About Me</div>
              <div className="column">
                <textarea
                  placeholder="Write something about yourself"
                  {...register("about")}
                />
                <br />
              </div>
            </div>
            <div className="row">
              <div className="column-title">User Joined at</div>
              <div className="column">{user.createdAt}</div>
            </div>
            <div className="row">
              <div className="column-title">User Profile is public</div>
              <div className="column">
                <input type="checkbox" {...register("isPublic")} />
              </div>
            </div>
            <button disabled={!isDirty}>Save changes</button>
          </form>
        </div>
      )}
    </>
  );
};

export default ProfileInfo;
