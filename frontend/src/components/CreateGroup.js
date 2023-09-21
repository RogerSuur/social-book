import { useState, useEffect } from "react";
import axios from "axios";
import { Container, Form, Button, FloatingLabel, Alert } from "react-bootstrap";

const CreateGroup = ({ onGroupCreated, handleClose }) => {
  const [modalOpen, setModalOpen] = useState(false);
  const [errMsg, setErrMsg] = useState("");
  const [formData, setFormData] = useState({
    title: "",
    description: "",
  });

  const handleChange = (event) => {
    const { name, value } = event.target;
    setFormData((prevState) => ({
      ...prevState,
      [name]: value,
    }));
  };

  const handleSubmit = async (event) => {
    event.preventDefault();

    if (formData?.title === "" || formData?.description === "") {
      return;
    }
    // Send the formData data to the backend handler
    try {
      const response = await axios.post(
        "http://localhost:8000/creategroup",
        JSON.stringify(formData),
        { withCredentials: true },
        {
          headers: { "Content-Type": "application/json" },
        }
      );
      setErrMsg(response.data?.message);
      // props.onPostsUpdate();
      //ACtion for creating new group
      onGroupCreated();
      handleClose();
    } catch (err) {
      if (!err?.response) {
        setErrMsg("No Server Response");
      } else {
        setErrMsg("Internal Server Error");
      }
    }
  };

  return (
    <>
      {errMsg && (
        <Alert variant="danger" className="text-center">
          {errMsg}
        </Alert>
      )}
      <Form onSubmit={handleSubmit}>
        <FloatingLabel
          className="mb-3"
          controlId="floatingTitle"
          label="Group name"
        >
          <Form.Control
            placeholder="Description"
            name="title"
            value={formData.title}
            onChange={handleChange}
            autoFocus
          />
        </FloatingLabel>
        <FloatingLabel
          className="mb-3"
          controlId="floatingDescription"
          label="Description"
        >
          <Form.Control
            as="textarea"
            placeholder="Description"
            name="description"
            value={formData.description}
            onChange={handleChange}
          />
        </FloatingLabel>
        <Button type="submit">Create</Button>
      </Form>
    </>
  );
};

export default CreateGroup;
