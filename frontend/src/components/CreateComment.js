import React, { useState } from "react";
import axios from "axios";
import ImageUploadModal from "./ImageUploadModal";

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
    <>
      <div
        style={{
          maxWidth: "100%",
          width: "100%",
          maxHeight: "200px",
          overflow: "hidden",
        }}
      >
        {errMsg && <h3>{errMsg}</h3>}

        {/* Display selected image(s) */}
        {selectedImage && (
          <img src={URL.createObjectURL(selectedImage)} alt="Selected" />
        )}

        <form onSubmit={handleSubmit}>
          <textarea
            style={{ width: "90%", height: "100%", resize: "none" }}
            placeholder="Your comment goes here"
            onChange={handleChange}
            value={formData.content}
            name="content"
          />
          <button>Post</button>
        </form>
        <button onClick={handleModalClick}>Add an image</button>
        <ImageUploadModal
          open={showModal}
          onClose={handleModalClose}
          onImageUpload={handleImageUpload}
        />
      </div>
    </>
  );
};

export default CreateComment;
