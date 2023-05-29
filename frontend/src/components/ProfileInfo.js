import React, { useState, useEffect } from "react";
import { useForm } from "react-hook-form";
import axios from "axios";
import Modal from "../components/Modal.js";
import AvatarUpdater from "../components/AvatarUpdater.js";

const PROFILE_URL = "http://localhost:8000/profile";
const PROFILE_UPDATE_URL = "http://localhost:8000/profile/update";

const ProfileInfo = (props) => {
  const [user, setUser] = useState({});
  const [errMsg, setErrMsg] = useState("");
  const [modalOpen, setModalOpen] = useState(false);

  const values = user;
  const {
    register,
    handleSubmit,
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
        .then((response) => {
          setUser(response.data.user);
          // console.log(response.data.user, "USRRR");
        });
    };
    loadUser();
  }, [modalOpen]);

  const onSubmit = async (data) => {
    try {
      const response = await axios.post(
        PROFILE_UPDATE_URL,
        JSON.stringify(data),
        {
          withCredentials: true,
          headers: { "Content-Type": "application/json" },
        }
      );

      console.log("RESPONSE:", JSON.stringify(response));
    } catch (err) {
      if (!err?.response) {
        setErrMsg("No Server Response");
      } else if (err.response?.status > 200) {
        setErrMsg("Internal Server Error");
      }
    }
  };

  const handleModalClose = () => {
    setModalOpen(false);
  };

  const handleModalClick = () => {
    setModalOpen(true);
  };

  // console.log(isDirty, "DIRTY");
  // console.log(user, "USER");

  return (
    <>
      {user && (
        <div className="content-area">
          <div className="row">
            <div className="column">
              <img
                style={{
                  width: "20vw",
                  height: "20vw",
                  objectFit: "cover",
                  objectPosition: "0% 100%",
                }}
                src={`images/${user.id}/${user.avatarImage}`}
                alt={`${user.firstName}`}
              ></img>
            </div>

            <h1 className="column-title">{user.firstName}'s profile</h1>
          </div>
          <div>
            <Modal open={modalOpen} onClose={handleModalClose}>
              <AvatarUpdater onUploadSuccess={handleModalClose} />
            </Modal>
            <button onClick={handleModalClick}>Upload New Image</button>
          </div>

          <form onSubmit={handleSubmit(onSubmit)}>
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
