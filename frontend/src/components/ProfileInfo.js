import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import axios from "axios";
import useWebSocketConnection from "../hooks/useWebSocketConnection";
import { useOutletContext } from "react-router-dom";
import { PROFILE_URL } from "../utils/routes";

const ProfileInfo = () => {
  const [user, setUser] = useState({});
  const { id } = useParams();
  const [errMsg, setErrMsg] = useState("");
  const { socketUrl } = useOutletContext();
  const { sendJsonMessage } = useWebSocketConnection(socketUrl);

  useEffect(() => {
    const loadUser = async () => {
      await axios
        .get(PROFILE_URL + id, {
          withCredentials: true,
        })
        .then((response) => {
          setUser(response.data.user);
        });
    };
    loadUser();
  }, [id]);

  console.log(user, "OTHER USER");

  const handleFollow = () => {
    sendJsonMessage({
      type: "follow_request",
      data: { id: user.id },
    });
  };

  const handleUnfollow = () => {
    sendJsonMessage({
      type: "unfollow",
      data: { id: user.id },
    });
  };

  const imageHandler = () => {
    const source = user?.avatarImage
      ? `${process.env.PUBLIC_URL}/images/${user.avatarImage}`
      : `${process.env.PUBLIC_URL}/images/${user.avatarImage}/defaultuser.jpg`;

    const image = <img className="profile-image" src={source} />;
    console.log(source);
    return image;
  };

  return (
    <>
      {user && (
        <div className="profile-area">
          <div className="row">
            <div>{imageHandler()}</div>
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
            <div className="column">
              {new Date(user.birthday).toLocaleDateString("en-UK", {
                month: "short",
                day: "numeric",
                year: "numeric",
              })}
            </div>
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
            <div className="column-title">User Joined</div>
            <div className="column">
              {new Date(user.createdAt).toLocaleDateString("en-UK", {
                month: "short",
                day: "numeric",
                year: "numeric",
              })}
            </div>
          </div>
          <div className="row">
            <div className="column-title">User Profile is public</div>
            <div className="column">{user.isPublic ? "Yes" : "No"}</div>
          </div>
          <button disabled={user.isFollowed} onClick={handleFollow}>
            Follow
          </button>
          <button disabled={!user.isFollowed} onClick={handleUnfollow}>
            Unfollow
          </button>
        </div>
      )}
    </>
  );
};

export default ProfileInfo;
