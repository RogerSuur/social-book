import React, { useState, useEffect } from "react";
import axios from "axios";
import Select from "react-select";
import ImageUploadModal from "./ImageUploadModal";
import {
  Form,
  Image,
  InputGroup,
  Alert,
  Button,
  Container,
  Col,
  Stack,
  Badge,
} from "react-bootstrap";
import PostButton from "../components/PostButton.js";
import { FOLLOWERS_URL } from "../utils/routes";
import GenericModal from "../components/GenericModal";
import { ImageFill } from "react-bootstrap-icons";

const CreatePost = ({ onPostsUpdate, handleClose }) => {
  const initialFormData = {
    content: "",
    image: null,
    privacyType: 0,
    selectedReceivers: [],
  };

  const [formData, setFormData] = useState(initialFormData);
  const [followers, setFollowers] = useState([]);
  const [errMsg, setErrMsg] = useState("");

  const handleImageUpload = (image) => {
    setFormData((prevFormData) => ({
      ...prevFormData,
      image: image,
    }));
  };
  console.log(formData);
  // errordata state

  const handleChange = (event) => {
    const { name, value, type } = event.target;

    setFormData((prevFormData) => ({
      ...prevFormData,
      [name]: type === "radio" ? parseInt(value) : value,
    }));
  };

  useEffect(() => {
    const fetchFollowers = async () => {
      try {
        const response = await axios.get(FOLLOWERS_URL, {
          withCredentials: true,
        });
        setFollowers(response.data);
      } catch (err) {
        console.error(err);
      }
    };
    if (formData.privacyType === 3) {
      fetchFollowers();
    }
  }, [formData.privacyType]);

  const handleSubmit = async (event) => {
    event.preventDefault();

    if (
      formData?.privacyType === 0 ||
      (formData?.content === "" && formData?.image === null)
    ) {
      return;
    }

    const formDataWithImage = new FormData();
    formDataWithImage.append("content", formData.content);
    formDataWithImage.append("privacyType", formData.privacyType);
    formDataWithImage.append(
      "selectedReceivers",
      JSON.stringify(formData.selectedReceivers)
    );

    if (formData?.image) {
      formDataWithImage.append("image", formData?.image); // Append the image file if it exists
    }

    console.log("data when submitted", formDataWithImage);

    try {
      const response = await axios.post(
        "http://localhost:8000/post",
        formDataWithImage,
        // JSON.stringify(formData),
        {
          withCredentials: true,
          headers: { "Content-Type": "multipart/form-data" },
        }
      );

      setErrMsg(response.data?.message);

      onPostsUpdate();
      handleClose();
    } catch (err) {
      if (!err?.response) {
        setErrMsg("No Server Response");
      } else {
        setErrMsg("Internal Server Error");
      }
    }

    setFormData(initialFormData);
  };

  const handleSelectChange = (selectedOptions) => {
    const selectedValues = selectedOptions.map((option) =>
      option.value.toString()
    );
    setFormData((prevFormData) => ({
      ...prevFormData,
      selectedReceivers: selectedValues,
    }));
  };

  const followersOptions = followers.map((follower) => ({
    value: follower.id,
    label: `${follower.firstName} ${follower.lastName}`,
  }));

  return (
    <>
      {errMsg && (
        <Alert variant="danger" className="text-center">
          {errMsg}
        </Alert>
      )}

      <div className="post-img mb-3">
        {formData?.image && (
          <Image
            src={URL.createObjectURL(formData?.image)}
            fluid
            alt="Selected"
          />
        )}
      </div>

      <Form onSubmit={handleSubmit}>
        <Stack direction="horizontal">
          <Col>
            <Stack direction="horizontal" gap="2">
              <InputGroup>
                <Form.Control
                  as="textarea"
                  placeholder="Write what's on your mind"
                  onChange={handleChange}
                  value={formData.content}
                  name="content"
                />
                <div>
                  <GenericModal img={<ImageFill />} buttonText="Add image">
                    <ImageUploadModal onUploadSuccess={handleImageUpload} />
                  </GenericModal>
                </div>
              </InputGroup>
              <Col as={PostButton} className="text-center" />
            </Stack>

            <Col className="mb-3">
              <Form.Check
                inline
                label="Public"
                name="privacyType"
                type="radio"
                id="public"
                value={1}
                onChange={handleChange}
              />
              <Form.Check
                inline
                label="Private"
                name="privacyType"
                type="radio"
                id="private"
                value={2}
                onChange={handleChange}
              />
              <Form.Check
                inline
                label="Sub-private"
                name="privacyType"
                type="radio"
                id="subPrivate"
                value={3}
                onChange={handleChange}
              />
            </Col>
            <Col>
              {formData.privacyType === 3 && (
                <Select
                  options={followersOptions}
                  isMulti
                  onChange={handleSelectChange}
                />
              )}
            </Col>
          </Col>
        </Stack>
      </Form>
    </>
  );
};

export default CreatePost;
