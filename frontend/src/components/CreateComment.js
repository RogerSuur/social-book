import React, { useState } from "react";
import axios from "axios";

const CreateComment = ({ postId, onCommentsUpdate }) => {
  const [formData, setFormData] = useState({
    postId: postId,
    content: "",
  });
  const [errMsg, setErrMsg] = useState("");

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

    try {
      const response = await axios.post(
        "http://localhost:8000/insertcomment",
        JSON.stringify(formData),
        { withCredentials: true },
        {
          headers: { "Content-Type": "application/json" },
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
  };

  return (
    <>
      <div style={{ maxWidth: '100%', width: '100%', maxHeight: '200px', overflow: 'hidden' }}>
        {errMsg && <h3>{errMsg}</h3>}
        <form onSubmit={handleSubmit}>
          <textarea
          style={{ width: '90%', height: '100%', resize: 'none' }}
            placeholder="Your comment goes here"
            onChange={handleChange}
            value={formData.content}
            name="content"
          />
          <button>Post</button>
        </form>
      </div>
    </>
  );
};

export default CreateComment;
