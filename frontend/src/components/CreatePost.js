import React from "react";
import { useState } from "react";
import axios from "axios";

const CreatePost = (props) => {
  const [formData, setFormData] = useState({
    content: "",
  });
  const [errMsg, setErrMsg] = useState("");

  const handleChange = (event) => {
    const { name, value, type, checked } = event.target;

    setFormData((prevFormData) => {
      return {
        ...prevFormData,
        [name]: type === "checkbox" ? checked : value,
      };
    });
  };

  const handleOptions = (selected) => {
    setFormData((prevFormData) => {
      return {
        ...prevFormData,
      };
    });
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
      props.handler();
    } catch (err) {
      if (!err?.response) {
        setErrMsg("No Server Response");
      } else {
        setErrMsg("Internal Server Error");
      }
    }

    setFormData({
      title: "",
      content: "",
    });
  };

  return (
    <>
      <div className="content-area">
        {errMsg && <h2>{errMsg}</h2>}
        <form onSubmit={handleSubmit}>
          <textarea
            placeholder="Write something..."
            onChange={handleChange}
            value={formData.body}
            name="content"
            required
          />
          <button>Post</button>
        </form>
      </div>
    </>
  );
};

export default CreatePost;
