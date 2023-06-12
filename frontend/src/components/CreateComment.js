import React, { useState } from "react";
import axios from "axios";

const CreateComment = ({ postId, onCommentsUpdate }) => {
  // const { postid } = useParams();

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
      onCommentsUpdate(1);

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
      <div className="content-area">
        {errMsg && <h3>{errMsg}</h3>}
        <form onSubmit={handleSubmit}>
          <textarea
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
