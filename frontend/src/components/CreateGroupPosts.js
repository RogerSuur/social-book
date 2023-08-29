import React, { useState } from "react";
import axios from "axios";
import { Form, Stack, Col, InputGroup } from "react-bootstrap";
import PostButton from "../components/PostButton";

const CreateGroupPost = ({ groupId, onPostsUpdate, handleClose }) => {
  const initialFormData = {
    content: "",
    imagePath: "",
    privacyType: 0,
    selectedReceivers: [],
  };

  const [formData, setFormData] = useState(initialFormData);
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
        console.log("ERROR: ", err);
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
        <Form onSubmit={handleSubmit}>
          <Col>
            <Stack direction="horizontal" gap="2">
              <InputGroup>
                <Form.Control
                  className="post-textarea"
                  type="textarea"
                  placeholder="Write what's on your mind"
                  onChange={handleChange}
                  value={formData.content}
                  name="content"
                />
              </InputGroup>
              <Col as={PostButton} className="text-center" />
            </Stack>
          </Col>
        </Form>
      </div>
    </>
  );
};

export default CreateGroupPost;
