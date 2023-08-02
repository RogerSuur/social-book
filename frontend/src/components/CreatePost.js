import { useState, useEffect } from "react";
import axios from "axios";
import Select from "react-select";
import ImageUploadModal from "./ImageUploadModal";

const CreatePost = (props) => {
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
      image: image, // Set the image in formData
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

    console.log("data when submitted", formData);
    console.log("type of image", typeof formData.image);

    try {
      const response = await axios.post(
        "http://localhost:8000/post",
        // formDataWithImage,
        JSON.stringify(formData),
        {
          withCredentials: true,
          headers: { "Content-Type": "multipart/form-data" }, // Set the correct Content-Type header
        }
      );

      setErrMsg(response.data?.message);

      props.onPostsUpdate();
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
    <>
      <div className="post-area">
        {errMsg && <h2>{errMsg}</h2>}

        {/* Display selected image(s) */}
        {selectedImage && (
          <img src={URL.createObjectURL(selectedImage)} alt="Selected" />
        )}

        <form onSubmit={handleSubmit}>
          <textarea
            className="area-text"
            type="text"
            placeholder="Write what's on your mind"
            onChange={handleChange}
            value={formData.content}
            name="content"
            required
          />
          <legend>Choose privacy type</legend>

          <input
            type="radio"
            id="public"
            name="privacyType"
            value={1}
            onChange={handleChange}
          />
          <label htmlFor="public">Public</label>

          <input
            type="radio"
            id="private"
            name="privacyType"
            value={2}
            onChange={handleChange}
          />
          <label htmlFor="private">Private</label>

          <input
            type="radio"
            id="subPrivate"
            name="privacyType"
            value={3}
            onChange={handleChange}
          />
          <label htmlFor="subPrivate">subPrivate</label>

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
          <button className="post-button">Post</button>
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

export default CreatePost;
