import React, { useState } from "react";
import axios from "axios";
import ImageUploadModal from "./ImageUploadModal";
import {
  Container,
  Row,
  Col,
  Form,
  Button,
  Image,
  InputGroup,
} from "react-bootstrap";
import { ImageFill } from "react-bootstrap-icons";

const CreateComment = ({ postId, onCommentsUpdate }) => {
  const [showModal, setShowModal] = useState(false);
  const [selectedImage, setSelectedImage] = useState(null);
  const [formData, setFormData] = useState({
    postId: postId,
    content: "",
    image: null,
  });
  const [errMsg, setErrMsg] = useState("");

  const handleImageUpload = (image) => {
    setFormData((prevFormData) => ({
      ...prevFormData,
      image: image,
    }));
    setSelectedImage(image);
    setShowModal(false);
  };

  const handleModalClose = () => {
    setShowModal(false);
  };

  const handleModalClick = () => {
    setShowModal(true);
  };

  const handleChange = (event) => {
    const { name, value } = event.target;

    setFormData((prevFormData) => {
      return {
        ...prevFormData,
        [name]: value,
      };
    });
  };

  const handleSubmit = async (event) => {
    event.preventDefault();

    if (formData.content === "" && formData.selectedImage === null) {
      return;
    }

    const formDataWithImage = new FormData();
    formDataWithImage.append("content", formData.content);
    formDataWithImage.append("postId", formData.postId);

    if (selectedImage) {
      formDataWithImage.append("image", selectedImage); // Append the image file if it exists
    }

    try {
      const response = await axios.post(
        "http://localhost:8000/insertcomment",
        formDataWithImage,
        {
          withCredentials: true,
          headers: { "Content-Type": "multipart/form-data" },
        }
      );

      setErrMsg(response.data?.message);
      onCommentsUpdate();

      if (!errMsg) {
        setFormData({
          postId: postId,
          content: "",
        });
      }
    } catch (err) {
      if (!err?.response) {
        setErrMsg("No Server Response");
      } else if (err.response?.status === 400) {
        setErrMsg("The comment should be between 1 and 100 characters long");
      } else {
        setErrMsg("Internal Server Error");
      }
    }

    setSelectedImage(null);
  };

  return (
    <Container>
      <Row
        style={{
          maxWidth: "100%",
          width: "100%",
          overflow: "hidden",
        }}
      >
        {errMsg && <h3>{errMsg}</h3>}

        {/* Display selected image(s) */}
        {selectedImage && (
          <Image
            fluid
            src={URL.createObjectURL(selectedImage)}
            alt="Selected"
          />
        )}
      </Row>
      <Row>
        <Col>
          <Form onSubmit={handleSubmit}>
            <InputGroup>
              <Form.Control
                className="post-textarea"
                type="textarea"
                placeholder="Write a comment"
                onChange={handleChange}
                value={formData.content}
                name="content"
              />
              <Button className="float-end" onClick={handleModalClick}>
                <ImageFill />
              </Button>
            </InputGroup>

            <Button className="float-start" type="submit">
              Post
            </Button>

            {showModal && (
              <ImageUploadModal
                open={showModal}
                onClose={handleModalClose}
                onImageUpload={handleImageUpload}
              />
            )}
          </Form>
        </Col>
      </Row>
    </Container>
  );
};

export default CreateComment;
