import React, { useState, useEffect } from "react";
import { useForm } from "react-hook-form";
import axios from "axios";
import Modal from "./Modal.js";
import AvatarUpdater from "./AvatarUpdater.js";

const PROFILE_URL = "http://localhost:8000/profile";
const PROFILE_UPDATE_URL = "http://localhost:8000/profile/update";

const ProfileEditor = (props) => {
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

  return (
    <>
     {user && (
  <div className="profile-area">
    <div className="top-part">
    
      
        <img
          className="profile-image"
          src={`images/${user.id}/${user.avatarImage}`}
          alt={`${user.firstName}`}
        />
      
      <div className="umber">
        <div className="profile-title-top">{user.firstName}'s profile</div>
        </div>
        <div className="profile-summary">{/* Summary content goes here */}</div>
      
    <div className="profile-actions">
      <Modal open={modalOpen} onClose={handleModalClose}>
        <AvatarUpdater onUploadSuccess={handleModalClose} />
      </Modal>
      <button onClick={handleModalClick}>Upload New Image</button>
    </div>
    <div className="profile-row-top">
      <div className="profile-title-top">About Me</div>
      <div className="profile-column-top">
        <textarea
          className="profile-textarea"
          placeholder="Write something about yourself"
          {...register("about")}
        />
        <br />
      </div>
    </div>
    </div>
    <div className="top-part">
    <div className="profile-row">
      <div className="profile-title">First Name</div>
      <div className="profile-column">{user.firstName}</div>
    </div>
    <div className="profile-row">
      <div className="profile-title">Last Name</div>
      <div className="profile-column">{user.lastName}</div>
    </div>
    <div className="profile-row">
      <div className="profile-title">Email</div>
      <div className="profile-column">{user.email}</div>
    </div>
    <div className="profile-row">
      <div className="profile-title">Birthday</div>
      <div className="profile-column">{user.birthday}</div>
    </div>
    <div className="profile-row">
      <div className="profile-title">Nickname</div>
      <div className="profile-column">
        <input
          className="profile-input"
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
    <div className="profile-row">
      <div className="profile-title">User Joined at</div>
      <div className="profile-column">{user.createdAt}</div>
    </div>
    <div className="profile-row">
      <div className="profile-title">User Profile is public</div>
      <div className="profile-column">
        <input className="profile-checkbox" type="checkbox" {...register("isPublic")} />
      </div>
    </div>
    
    <button className="profile-button" disabled={!isDirty}>Save changes</button>
    </div>
  </div>
)}
    </>
  );
};

export default ProfileEditor;
