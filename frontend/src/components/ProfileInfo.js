import React, { useState, useEffect } from "react";
import {
  useParams,
  useOutletContext,
  useNavigate,
  Link,
} from "react-router-dom";
import axios from "axios";
import useWebSocketConnection from "../hooks/useWebSocketConnection";
import { PROFILE_URL, FOLLOW_URL } from "../utils/routes";
import ImageHandler from "../utils/imageHandler.js";
import Modal from "../components/Modal.js";
import Posts from "../pages/PostsPage.js";

const ProfileInfo = () => {
  const [user, setUser] = useState({});
  const [followers, setFollowers] = useState([]);
  const [following, setFollowing] = useState([]);
  const [modalOpen, setModalOpen] = useState(false);
  const [postsModalOpen, setPostsModalOpen] = useState(false);
  const [activeTab, setActiveTab] = useState(true);
  const { id } = useParams();
  const { socketUrl } = useOutletContext();
  const { sendJsonMessage } = useWebSocketConnection(socketUrl);
  const [errMsg, setErrMsg] = useState("");
  const navigate = useNavigate();

  useEffect(() => {
    const loadUser = async () => {
      try {
        await axios
          .get(PROFILE_URL + id, {
            withCredentials: true,
          })
          .then((response) => {
            if (response?.data?.isOwnProfile === true) {
              navigate("/profile", { replace: true });
            } else {
              setUser(response?.data?.user);
            }
          });
      } catch (err) {
        if (!err?.response) {
          setErrMsg("No Server Response");
        } else if (err.response?.status === 303) {
          navigate("/profile", { replace: true });
        } else {
          setErrMsg("Internal Server Error");
        }
      }
    };
    loadUser();
  }, [id]);

  useEffect(() => {
    const loadFollow = async () => {
      try {
        await axios
          .get(FOLLOW_URL + id, {
            withCredentials: true,
          })
          .then((response) => {
            setFollowers(response?.data?.followers);
            setFollowing(response?.data?.following);
          });
      } catch (err) {
        if (!err?.response) {
          setErrMsg("No Server Response");
        } else {
          setErrMsg("Internal Server Error");
        }
      }
    };
    loadFollow();
  }, [setActiveTab]);

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

  const birthdayConverter = (date) => {
    if (!date) {
      return;
    }
    const [day, month, year] = date?.split("/");
    return new Date(year, month - 1, day).toLocaleDateString("en-UK", {
      month: "short",
      day: "numeric",
      year: "numeric",
    });
  };

  const handleModalClose = () => {
    setModalOpen(false);
  };

  const handleModalClick = (follow) => {
    setActiveTab(follow);
    setModalOpen(true);
  };

  const handlePostsModalClose = () => {
    setPostsModalOpen(false);
  };

  const handlePostsModalClick = () => {
    setPostsModalOpen(true);
  };

  const userList = (follow) => {
    const users = follow ? [...following] : [...followers];

    return users?.map((user, index) => (
      <li key={index}>
        <Link to={`/profile/${user.userId}`}>
          {ImageHandler(user.imagePath, "defaultuser.jpg", "profile-image")}
          <p>{user.name}</p>
        </Link>
      </li>
    ));
  };

  const image = () =>
    ImageHandler(user?.avatarImage, "defaultuser.jpg", "profile-image");

  return (
    <>
      {user && (
        <div className="profile-area1">
          <div className="row">
            <div>{image()}</div>
            <h1 className="column-title">{user.firstName}'s profile</h1>
          </div>
          <Modal open={modalOpen} onClose={handleModalClose}>
            <ul>
              <li onClick={() => setActiveTab(true)}>Following</li>
              <li onClick={() => setActiveTab(false)}>Followers</li>
            </ul>
            <ul>{userList(activeTab)}</ul>
          </Modal>
          <button onClick={() => handleModalClick(true)}>Following</button>
          <button onClick={() => handleModalClick(false)}>Followers</button>
          <Modal open={postsModalOpen} onClose={handlePostsModalClose}>
            <Posts />
          </Modal>
          <button onClick={() => handlePostsModalClick()}>Posts</button>

          <div className="row">
            <div className="column-title">First Name</div>
            <div className="column">{user.firstName}</div>
          </div>
          <div className="row">
            <div className="column-title">Last Name</div>
            <div className="column">{user.lastName}</div>
          </div>
          <div className="row">
            <div className="column-title">Nickname</div>
            <div className="column">{user.nickname}</div>
          </div>
          {(user?.isPublic || user?.isFollowed) && (
            <>
              <div className="row">
                <div className="column-title">Email</div>
                <div className="column">{user.email}</div>
              </div>
              <div className="row">
                <div className="column-title">Birthday</div>
                <div className="column">
                  {birthdayConverter(user?.birthday)}
                </div>
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
            </>
          )}

          <div className="row">
            <div className="column-title">Profile Type</div>
            <div className="column">{user.isPublic ? "Public" : "Private"}</div>
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
