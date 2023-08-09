import { useState, useEffect } from "react";
import axios from "axios";
import Select from "react-select";
import ImageUploadModal from "./ImageUploadModal";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import Container from "react-bootstrap/Container";
import PostButton from "../components/PostButton.js";

const CreatePost = ({ onPostsUpdate }) => {
  const initialFormData = {
    content: "",
    image: null,
    privacyType: 1,
    selectedReceivers: [],
  };

  const [formData, setFormData] = useState(initialFormData);
  const [followers, setFollowers] = useState([]);
  const [selectedImage, setSelectedImage] = useState(null);
  const [showModal, setShowModal] = useState(false);
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
        const response = await axios.get("http://localhost:8000/followers", {
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

    const formDataWithImage = new FormData();
    formDataWithImage.append("content", formData.content);
    formDataWithImage.append("privacyType", formData.privacyType);
    formDataWithImage.append(
      "selectedReceivers",
      JSON.stringify(formData.selectedReceivers)
    );

    if (selectedImage) {
      formDataWithImage.append("image", selectedImage); // Append the image file if it exists
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
    } catch (err) {
      if (!err?.response) {
        setErrMsg("No Server Response");
      } else {
        setErrMsg("Internal Server Error");
      }
    }

    setFormData(initialFormData);
    setSelectedImage(null);
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
    <Container>
      <div className="post-area">
        {errMsg && <h2>{errMsg}</h2>}

        {/* Display selected image(s) */}
        {selectedImage && (
          <img src={URL.createObjectURL(selectedImage)} alt="Selected" />
        )}
        <Form onSubmit={handleSubmit}>
          <Form.Control
            className="post-textarea"
            type="textarea"
            placeholder="Write what's on your mind"
            onChange={handleChange}
            value={formData.content}
            name="content"
            required
          />
          <div key={"inline-radio"} className="mb-3">
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
          </div>

          {formData.privacyType === 3 && (
            <>
              <legend>Choose receiver(s)</legend>
              <Select
                options={followersOptions}
                isMulti
                onChange={handleSelectChange}
              />
            </>
          )}
          <PostButton />
          <Button className="float-end" onClick={handleModalClick}>
            Add an image
          </Button>
          <ImageUploadModal
            open={showModal}
            onClose={handleModalClose}
            onImageUpload={handleImageUpload}
          />
        </Form>
      </div>
    </Container>
  );
};

export default CreatePost;
