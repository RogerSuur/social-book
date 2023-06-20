import React from "react";
import { useState } from "react";
import axios from "axios";

const CreatePost = (props) => {
  const initialFormData = {
    content: "",
    imagePath: "",
    privacyType: 1,
  };

  const [formData, setFormData] = useState(initialFormData);

  // errordata state
  const [errMsg, setErrMsg] = useState("");

  const handleChange = (event) => {
    const { name, value, type } = event.target;

    setFormData((prevFormData) => {
      return {
        ...prevFormData,
        [name]: type === "radio" ? parseInt(value) : value,
      };
    });
    // console.log(formData);
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    try {
      const response = await axios.post(
        "http://localhost:8000/post",
        JSON.stringify(formData),
        { withCredentials: true },
        {
          headers: { "Content-Type": "application/json" },
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
  };

  return (
    <>
      <div className="content-area">
        {errMsg && <h2>{errMsg}</h2>}
        <form onSubmit={handleSubmit}>
        <textarea className="area-text"
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

          <button className="post-button">Post</button>
        </form>
      </div>
    </>
  );
};

export default CreatePost;
