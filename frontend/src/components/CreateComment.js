import React, { useState } from "react";
import axios from "axios";

const CreateComment = ({ postid }) => {
  // const { postid } = useParams();
  // console.log(postid);
  const [formData, setFormData] = useState({
    postId: postid,
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

  const handleSubmit = async (event) => {
    event.preventDefault();
    console.log(formData);

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
      //props.handler();
      if (!errMsg) {
        setFormData({
          postId: postid,
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
