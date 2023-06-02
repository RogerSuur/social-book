import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import axios from "axios";
import { useWebSocketConnection } from "../hooks/useWebSocketConnection";
import { useOutletContext } from "react-router-dom";

const PROFILE_URL = "http://localhost:8000/profile/";
const PROFILE_FOLLOW_URL = "http://localhost:8000/profile/follow";

const ProfileInfo = () => {
  const [user, setUser] = useState({});
  const { id } = useParams();
  const [errMsg, setErrMsg] = useState("");
  const { socketUrl } = useOutletContext();

  useEffect(() => {
    const loadUser = async () => {
      await axios
        .get(`http://localhost:8000/profile/${id}`, {
          withCredentials: true,
        })
        .then((response) => {
          setUser(response.data.user);
          console.log(response.data.user, "USRRR");
        });
    };
    loadUser();
  }, []);

  const sendJsonMessage = useWebSocketConnection(socketUrl);

  const handleClick = (message) => {
    console.log("here");
    sendJsonMessage(message);
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
          <button disabled={user.follow} onClick={() => handleClick("follow")}>
            Follow
          </button>
          <button
            disabled={user.follow}
            onClick={() => handleClick("unfollow")}
          >
            Unfollow
          </button>
        </div>
      )}
    </>
  );
};

export default ProfileInfo;
