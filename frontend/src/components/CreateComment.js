import React, { useState } from "react";
import { useParams } from "react-router-dom";
import axios from "axios";

const CreateComment = (props) => {
  const { id } = useParams();
  const [formData, setFormData] = useState({
    post_id: id,
    body: "",
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

  const handleSubmit = async (event) => {
    event.preventDefault();

    try {
      const response = await axios.post(
        "http://localhost:8000/createcomment",
        JSON.stringify(formData),
        { withCredentials: true },
        {
          headers: { "Content-Type": "application/json" },
        }
      );

      setErrMsg(response.data?.message);
      props.handler();
      if (!errMsg) {
        setFormData({
          post_id: id,
          body: "",
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
      <div className="content-area">
        {errMsg && <h3>{errMsg}</h3>}
        <form onSubmit={handleSubmit}>
          <textarea
            placeholder="Your comment goes here"
            onChange={handleChange}
            value={formData.body}
            name="body"
          />
          <button>Post</button>
        </form>
      </div>
    </>
  );
};

export default CreateComment;
