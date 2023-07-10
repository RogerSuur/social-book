import Modal from "./Modal";
import AvatarUpdater from "./AvatarUpdater.js";
import { useState, useEffect } from "react";
import axios from "axios";

const CreateGroup = () => {
  const [modalOpen, setModalOpen] = useState(false);
  const [errMsg, setErrMsg] = useState("");
  const [groupCreateForm, setGroupCreateForm] = useState({
    title: "",
    description: "",
    image: null,
  });

  const openModal = () => {
    setModalOpen(true);
  };

  const closeModal = () => {
    setModalOpen(false);
  };

  const handleChange = (event) => {
    const { name, value } = event.target;
    console.log("name", name, "value", value);
    setGroupCreateForm((prevState) => ({
      ...prevState,
      [name]: value,
    }));
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    // Send the groupCreateForm data to the backend handler
    console.log(groupCreateForm);

    try {
      const response = await axios.post(
        "http://localhost:8000/creategroup",
        JSON.stringify(groupCreateForm),
        { withCredentials: true },
        {
          headers: { "Content-Type": "application/json" },
        }
      );

      setErrMsg(response.data?.message);

      // props.onPostsUpdate();
      //ACtion for creating new group
    } catch (err) {
      if (!err?.response) {
        setErrMsg("No Server Response");
      } else {
        setErrMsg("Internal Server Error");
      }
    }

    closeModal();
  };

  return (
    <>
      <p>Here you can create a group</p>
      <button onClick={openModal}>Create New group</button>
      <Modal open={modalOpen} onClose={closeModal}>
        <form onSubmit={handleSubmit}>
          <label>
            Title:
            <input
              type="text"
              name="title"
              value={groupCreateForm.title}
              onChange={handleChange}
              required
            />
          </label>
          <br />
          <label>
            Description:
            <textarea
              name="description"
              value={groupCreateForm.description}
              onChange={handleChange}
              required
            ></textarea>
          </label>
          <br />

          <br />
          {/* <label>
            Image:
            <input
              type="file"
              name="image"
              accept="image/*"
              onChange={handleImageUpload}
            />
          </label> */}

          <button type="submit">Create</button>
        </form>
      </Modal>
      {/* <button onClick={openModal}>Upload New Image</button> */}
    </>
  );
};

export default CreateGroup;
