import React, { useState, useEffect } from "react";
import { useParams, useOutletContext, useNavigate } from "react-router-dom";
import axios from "axios";
import useWebSocketConnection from "../hooks/useWebSocketConnection";
import {
  PROFILE_URL,
  USER_FOLLOWING_URL,
  USER_FOLLOWERS_URL,
  USER_POSTS_URL,
} from "../utils/routes";
import ImageHandler from "../utils/imageHandler.js";
import FeedPosts from "../components/FeedPosts.js";
import { Container, Row, Col, Button } from "react-bootstrap";
import GenericUserList from "../components/GenericUserList";
import GenericModal from "../components/GenericModal";
import { BirthdayConverter, LongDate } from "../utils/datetimeConverters";

const ProfileInfo = () => {
  const [user, setUser] = useState({});
  const [isFollowed, setIsFollowed] = useState(false);
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

  const userList = (following) =>
    following ? (
      <GenericUserList url={USER_FOLLOWING_URL + id} />
    ) : (
      <GenericUserList url={USER_FOLLOWERS_URL + id} />
    );

  const image = ImageHandler(user?.imagePath, "defaultuser.jpg", "profile-img");

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
                    <p>{BirthdayConverter(user?.birthday)}</p>
                  </Col>
                  <Col>
                    <p>Joined</p>
                    <p>{LongDate(user.createdAt)}</p>
                  </Col>
                </Row>
                <Row className="d-grip gap-2">
                  <Col>
                    <GenericModal buttonText="Following">
                      {userList(true)}
                    </GenericModal>
                  </Col>
                  <Col>
                    <GenericModal buttonText="Followers">
                      {userList(false)}
                    </GenericModal>
                  </Col>
                  <Col>
                    <GenericModal buttonText="User's posts">
                      <FeedPosts url={USER_POSTS_URL + id} />
                    </GenericModal>
                  </Col>
                </Row>
              </>
            )}
          </Row>
        </>
      )}
    </Container>
  );
};

export default ProfileInfo;
