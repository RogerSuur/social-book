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
  Stack,
  Alert,
} from "react-bootstrap";
import { ImageFill } from "react-bootstrap-icons";
import GenericModal from "../components/GenericModal";
import Scrollbars from "react-custom-scrollbars-2";

const CreateComment = ({ postId, onCommentsUpdate }) => {
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
    setErrMsg("");
  };

  const handleChange = (event) => {
    setErrMsg("");
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
    <Container className="mt-3 pt-3 pb-3 bg-light border rounded">
      <Row>
        {errMsg && (
          <Alert variant="danger" className="text-center">
            {errMsg}
          </Alert>
        )}
        {selectedImage && (
          <Image
            fluid
            src={URL.createObjectURL(selectedImage)}
            alt="Selected"
          />
        )}
      </Row>
      <Row>
        <Form onSubmit={handleSubmit}>
          <Stack direction="horizontal">
            <InputGroup className="me-2">
              <Form.Control
                as="textarea"
                placeholder="Write a comment"
                onChange={handleChange}
                value={formData.content}
                name="content"
              />
            </InputGroup>
            <div>
              <GenericModal
                variant="flush"
                img={<ImageFill />}
                buttonText="Add an image"
              >
                <ImageUploadModal onUploadSuccess={handleImageUpload} />
              </GenericModal>
            </div>
            <Button type="submit">Post</Button>
          </Stack>
        </Form>
      </Row>
    </Container>
  );
};

export default CreateComment;
