import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import axios from "axios";

const PROFILE_URL = "http://localhost:8000/profile/";
const PROFILE_FOLLOW_URL = "http://localhost:8000/profile/follow";

const ProfileEditor = () => {
  const [user, setUser] = useState({});
  const { id } = useParams();
  const [errMsg, setErrMsg] = useState("");

  console.log(PROFILE_URL);

  useEffect(() => {
    const loadUser = async () => {
      await axios
        .get(PROFILE_URL + `${id}`, {
          withCredentials: true,
        })
        .then((response) => {
          setUser(response.data.user);
          console.log(response.data.user, "USRRR");
        });
    };
    loadUser();
  }, []);

  const handleFollow = async () => {
    // try {
    //   const response = await axios.post(
    //     PROFILE_FOLLOW_URL,
    //     JSON.stringify(data),
    //     {
    //       withCredentials: true,
    //       headers: { "Content-Type": "application/json" },
    //     }
    //   );
    //   console.log("RESPONSE:", JSON.stringify(response));
    // } catch (err) {
    //   if (!err?.response) {
    //     setErrMsg("No Server Response");
    //   } else if (err.response?.status > 200) {
    //     setErrMsg("Internal Server Error");
    //   }
    // }
  };

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
            <div className="column">{user.nickname}</div>
          </div>
          <div className="row">
            <div className="column-title">About Me</div>
            <div className="column">{user.about}</div>
          </div>
          <div className="row">
            <div className="column-title">User Joined at</div>
            <div className="column">{user.createdAt}</div>
          </div>
          <div className="row">
            <div className="column-title">User Profile is public</div>
            <div className="column">{user.isPublic}</div>
          </div>
          <button onClick={handleFollow}>Follow</button>
          <button onClick={handleFollow}>Unfollow</button>
        </div>
      )}
    </>
  );
};

export default ProfileEditor;
