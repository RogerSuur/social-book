import React from "react";
import { useState } from "react";
import axios from "axios";

const CreatePost = (props) => {
  // state kus sees hoiame vormidatat
  const initialFormData = {
    content: "",
    imagePath: "",
    privacyType: 1,
  };

  const [formData, setFormData] = useState(initialFormData);

  // errordata state
  const [errMsg, setErrMsg] = useState("");

  // muudab meie form state valuesi
  const handleChange = (event) => {
    const { name, value, type } = event.target;

    setFormData((prevFormData) => {
      return {
        ...prevFormData,
        [name]: type === "radio" ? parseInt(value) : value,
      };
    });
    console.log(formData);
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
      // n'itab lihtsalt et on 'ra loadinud
      props.handler();
    } catch (err) {
      if (!err?.response) {
        setErrMsg("No Server Response");
      } else {
        setErrMsg("Internal Server Error");
      }
    }

    // teeb p'rast lihtsalt tyhjaks vormi
    setFormData(initialFormData);
  };

  return (
    <>
      <div className="content-area">
        {errMsg && <h2>{errMsg}</h2>}
        <form onSubmit={handleSubmit}>
          <textarea
            type="text"
            placeholder="Write something..."
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

          <button>Post</button>
        </form>
      </div>
    </>
  );
};

export default CreatePost;
