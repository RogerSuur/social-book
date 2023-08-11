import React, { useState, useEffect } from "react";
import {
  useParams,
  useOutletContext,
  useNavigate,
  Link,
} from "react-router-dom";
import axios from "axios";
import useWebSocketConnection from "../hooks/useWebSocketConnection";
import {
  PROFILE_URL,
  USER_FOLLOWING_URL,
  USER_FOLLOWERS_URL,
  USER_POSTS_URL,
} from "../utils/routes";
import ImageHandler from "../utils/imageHandler.js";
import Modal from "../components/Modal.js";
import FeedPosts from "../components/FeedPosts.js";
import Container from "react-bootstrap/Container";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import Button from "react-bootstrap/Button";

const ProfileInfo = () => {
  const [user, setUser] = useState({});
  const [followers, setFollowers] = useState([]);
  const [following, setFollowing] = useState([]);
  const [modalOpen, setModalOpen] = useState(false);
  const [isFollowed, setIsFollowed] = useState(false);
  const [postsModalOpen, setPostsModalOpen] = useState(false);
  const [activeTab, setActiveTab] = useState(true);
  const { id } = useParams();
  const { socketUrl } = useOutletContext();
  const { sendJsonMessage } = useWebSocketConnection(socketUrl);
  const [errMsg, setErrMsg] = useState("");
  const navigate = useNavigate();

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
    setIsFollowed(false);
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

  useEffect(() => {
    const loadUser = async () => {
      try {
        await axios
          .get(PROFILE_URL + `/${id}`, {
            withCredentials: true,
          })
          .then((response) => {
            if (response?.data?.user?.isOwnProfile === true) {
              navigate("/profile", { replace: true });
            } else {
              console.log("PROFILE INFO RESPONSE: ", response?.data);
              setUser(response?.data?.user);
              setIsFollowed(response?.data?.user?.isFollowed);
            }
          });
      } catch (err) {
        if (!err?.response) {
          setErrMsg("No Server Response");
        } else {
          setErrMsg("Internal Server Error");
        }
      }
    };
    loadUser();
  }, [id, isFollowed]);

  useEffect(() => {
    const loadFollowers = async () => {
      try {
        await axios
          .get(USER_FOLLOWERS_URL + id, {
            withCredentials: true,
          })
          .then((response) => {
            console.log("FOLLOWERS: ", response?.data);
            setFollowers(response?.data);
          });
      } catch (err) {
        if (!err?.response) {
          setErrMsg("No Server Response");
        } else {
          setErrMsg("Internal Server Error");
        }
      }
    };
    loadFollowers();
  }, [id]);

  useEffect(() => {
    const loadFollowing = async () => {
      try {
        await axios
          .get(USER_FOLLOWING_URL + id, {
            withCredentials: true,
          })
          .then((response) => {
            console.log("FOLLOWING: ", response?.data);
            setFollowing(response?.data);
          });
      } catch (err) {
        if (!err?.response) {
          setErrMsg("No Server Response");
        } else {
          setErrMsg("Internal Server Error");
        }
      }
    };
    loadFollowing();
  }, [id]);

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

  const userList = (follow) => {
    const users = follow ? [...following] : [...followers];

    return users?.map((user, index) => (
      <li key={index}>
        <Link to={`/profile/${user.userId}`}>
          {ImageHandler(user.imagePath, "defaultuser.jpg", "profile-image")}
          {user?.nickname ? (
            <p>{user.nickname}</p>
          ) : (
            <p>
              {user.firstName} {user.lastName}
            </p>
          )}
        </Link>
      </li>
    ));
  };

  const image = ImageHandler(
    user?.avatarImage,
    "defaultuser.jpg",
    "profile-img"
  );

  return (
    <Container fluid>
      {user && (
        <>
          <Row className="gap-2">
            <Col sm>
              <div className="profile-img">{image}</div>
            </Col>
            <Col sm>
              <Row className="d-grid gap-2">
                <Col xs="12">
                  <h1>
                    {user.firstName} {user.lastName}
                  </h1>
                </Col>
                <Col xs="12">
                  <div>also known as </div>
                  <h1 className="display-5">{user.nickname}</h1>
                </Col>
                <Row className="gap-2">
                  <Col
                    lg="5"
                    as={Button}
                    disabled={isFollowed}
                    onClick={handleFollow}
                  >
                    Follow
                  </Col>

                  <Col
                    lg="5"
                    as={Button}
                    disabled={!isFollowed}
                    onClick={handleUnfollow}
                  >
                    Unfollow
                  </Col>
                </Row>
              </Row>
            </Col>
          </Row>

          <Row>
            {(user?.isPublic || user?.isFollowed) && (
              <>
                <Row>
                  <Col>
                    <div>{user.about}</div>
                  </Col>
                </Row>
                <Row>
                  <Col>
                    <p>Email address</p>
                    <p>{user.email}</p>
                  </Col>
                  <Col>
                    <p>Profile Type</p>
                    <p>{user.isPublic ? "Public" : "Private"}</p>
                  </Col>
                </Row>
                <Row>
                  <Col>
                    <p>Born</p>
                    <p>{birthdayConverter(user?.birthday)}</p>
                  </Col>
                  <Col>
                    <p>Joined</p>
                    <p>
                      {new Date(user.createdAt).toLocaleDateString("en-UK", {
                        month: "short",
                        day: "numeric",
                        year: "numeric",
                      })}
                    </p>
                  </Col>
                </Row>
              </>
            )}
          </Row>
          <Row className="d-grip gap-2">
            <Modal open={modalOpen} onClose={handleModalClose}>
              <ul>
                <li onClick={() => setActiveTab(true)}>Following</li>
                <li onClick={() => setActiveTab(false)}>Followers</li>
              </ul>
              <ul>{userList(activeTab)}</ul>
            </Modal>
            <Col md as={Button} onClick={() => handleModalClick(true)}>
              Following
            </Col>
            <Col md as={Button} onClick={() => handleModalClick(false)}>
              Followers
            </Col>
            <Modal open={postsModalOpen} onClose={handlePostsModalClose}>
              <FeedPosts url={USER_POSTS_URL + id} />
            </Modal>
            <Col md as={Button} onClick={() => handlePostsModalClick()}>
              Posts
            </Col>
          </Row>
        </>
      )}
    </Container>
  );
};

export default ProfileInfo;
