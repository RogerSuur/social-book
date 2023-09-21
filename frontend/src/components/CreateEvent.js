import React, { useState } from "react";
import axios from "axios";
import { Form, Button, Alert, FloatingLabel } from "react-bootstrap";

const CreateEvent = ({ onEventCreated, id, handleClose }) => {
  const [errMsg, setErrMsg] = useState("");
  const [formData, setFormData] = useState({
    title: "",
    description: "",
    startTime: "",
    endTime: "",
    group_id: +id,
  });

  const handleChange = (event) => {
    const { name, value } = event?.target;
    setFormData((prevState) => ({
      ...prevState,
      [name]: value,
    }));
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    // Send the formData data to the backend handler
    try {
      const response = await axios.post(
        "http://localhost:8000/creategroupevent",
        JSON.stringify({
          ...formData,
          startTime: new Date(formData.startTime).toISOString(),
          endTime: new Date(formData.endTime).toISOString(),
        }),
        { withCredentials: true },
        {
          headers: { "Content-Type": "application/json" },
        }
      );
      setErrMsg(response.data?.message);
      onEventCreated();
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
        <FloatingLabel className="mb-3" controlId="floatingTitle" label="Name">
          <Form.Control
            placeholder="Event name"
            name="title"
            value={formData.title}
            onChange={handleChange}
            autoFocus
            required
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
            required
          />
        </FloatingLabel>
        <FloatingLabel
          className="mb-3"
          controlId="floatingStartTime"
          label="Start time"
        >
          <Form.Control
            type="datetime-local"
            placeholder="Start time"
            name="startTime"
            value={formData.startTime}
            onChange={handleChange}
            required
          />
        </FloatingLabel>
        <FloatingLabel
          className="mb-3"
          controlId="floatingEndTime"
          label="End time"
        >
          <Form.Control
            type="datetime-local"
            placeholder="End time"
            name="endTime"
            value={formData.endTime}
            onChange={handleChange}
            required
          />
        </FloatingLabel>
        <Button type="submit">Create</Button>
      </Form>
    </>
  );
};

export default CreateEvent;
