import { useState, useEffect } from "react";
import axios from "axios";
import Select from "react-select";

const CreateGroupPost = ({ groupId, onPostsUpdate }) => {
  const initialFormData = {
    content: "",
    imagePath: "",
    privacyType: 0,
    selectedReceivers: [],
  };

  const [formData, setFormData] = useState(initialFormData);

  // errordata state
  const [errMsg, setErrMsg] = useState("");

  const handleChange = (event) => {
    const { name, value, type } = event.target;

    setFormData((prevFormData) => ({
      ...prevFormData,
      [name]: type === "radio" ? parseInt(value) : value,
    }));
  };

  const handleSubmit = async (event) => {
    event.preventDefault();

    try {
      const response = await axios.post(
        `http://localhost:8000/groups/${groupId}/post`,
        JSON.stringify(formData),
        { withCredentials: true },
        {
          headers: { "Content-Type": "application/json" },
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
  };

  return (
    <>
      <div className="group-post-area">
        {errMsg && <h2>{errMsg}</h2>}
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

          <button className="post-button">Post</button>
        </form>
      </div>
    </>
  );
};

export default CreateGroupPost;
