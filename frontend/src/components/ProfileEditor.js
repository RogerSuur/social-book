import React, { useState, useEffect } from "react";
import { useForm } from "react-hook-form";
import axios from "axios";
import Modal from "./Modal.js";
import AvatarUpdater from "./AvatarUpdater.js";
import ImageHandler from "../utils/imageHandler.js";
import Container from "react-bootstrap/Container";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import FloatingLabel from "react-bootstrap/FloatingLabel";
import Alert from "react-bootstrap/Alert";
import FeedPosts from "../components/FeedPosts.js";

import {
  PROFILE_URL,
  PROFILE_UPDATE_URL,
  PROFILE_POSTS_URL,
  AVATAR_UPDATER_URL,
  FOLLOWERS_URL,
  FOLLOWING_URL,
} from "../utils/routes.js";
import GenericUserList from "./GenericUserList.js";

const ProfileEditor = () => {
  const [user, setUser] = useState({});
  const [errMsg, setErrMsg] = useState("");
  const [modalOpen, setModalOpen] = useState("");
  const [activeTab, setActiveTab] = useState(true);

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

  console.log("EDITOR PROFILE: ", user);

  const image = ImageHandler(
    user?.avatarImage,
    "defaultuser.jpg",
    "profile-img"
  );

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
      await axios.post(PROFILE_UPDATE_URL, JSON.stringify(data), {
        withCredentials: true,
        headers: { "Content-Type": "application/json" },
      });
    } catch (err) {
      if (!err?.response) {
        setErrMsg("No Server Response");
      } else if (err.response?.status > 200) {
        setErrMsg("Internal Server Error");
      }
    }
  };

  const handleModalClose = () => {
    setModalOpen("");
  };

  const handleModalClick = (follow) => {
    setModalOpen("follow");
    setActiveTab(follow);
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

  const userList = activeTab ? (
    <GenericUserList url={FOLLOWERS_URL} />
  ) : (
    <GenericUserList url={FOLLOWING_URL} />
  );

  return (
    <Container fluid>
      {user && (
        <>
          <Form onSubmit={handleSubmit(onSubmit)}>
            <Row className="gap-2">
              <Col sm>
                <div className="profile-img">{image}</div>
                <Modal open={modalOpen === "avatar"} onClose={handleModalClose}>
                  <AvatarUpdater
                    url={AVATAR_UPDATER_URL}
                    onUploadSuccess={handleModalClose}
                  />
                </Modal>
                <Button onClick={() => setModalOpen("avatar")}>
                  Upload New Image
                </Button>
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
                    <FloatingLabel
                      className="mb-3"
                      controlId="floatingNickname"
                      label="Nickname (optional)"
                    >
                      <Form.Control
                        placeholder="Enter your nickname"
                        {...register("nickname", {
                          maxLength: {
                            value: 32,
                            message:
                              "A nickname should not be longer than 32 characters long",
                          },
                          pattern: {
                            value: /^[a-zA-Z0-9._ ]{0,32}$/,
                            message:
                              "A nickname can only contain letters, numbers, spaces, dots (.) and underscores (_)",
                          },
                        })}
                      />
                      {errors.nickname && (
                        <Alert variant="danger">
                          {errors.nickname.message}
                        </Alert>
                      )}
                    </FloatingLabel>
                  </Col>
                  <div className="mb-3">
                    <Form.Check
                      type="default-checkbox"
                      label="Profile is public"
                      {...register("isPublic")}
                    />
                  </div>
                </Row>
              </Col>
            </Row>

            <Row>
              <Row>
                <Col>
                  <FloatingLabel
                    className="mb-3"
                    controlId="about"
                    label="About you (optional)"
                  >
                    <Form.Control
                      type="textarea"
                      placeholder="Write something about yourself"
                      {...register("about")}
                    />
                  </FloatingLabel>
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
            </Row>
            <Row className="d-grip gap-2">
              <Modal open={modalOpen === "follow"} onClose={handleModalClose}>
                <ul>
                  <li onClick={() => setActiveTab(true)}>Following</li>
                  <li onClick={() => setActiveTab(false)}>Followers</li>
                </ul>
                <ul>{userList}</ul>
              </Modal>
              <Col md as={Button} onClick={() => handleModalClick(true)}>
                Following
              </Col>
              <Col md as={Button} onClick={() => handleModalClick(false)}>
                Followers
              </Col>
              <Modal open={modalOpen === "posts"} onClose={handleModalClose}>
                <FeedPosts url={PROFILE_POSTS_URL} />
              </Modal>
              <Col md as={Button} onClick={() => setModalOpen("posts")}>
                Posts
              </Col>
            </Row>
            <Button disabled={!isDirty}>Save changes</Button>
          </Form>
        </>
      )}
    </Container>
  );
};

export default ProfileEditor;
