import Modal from "./Modal";
import AvatarUpdater from "./AvatarUpdater.js";
import { useState, useEffect } from "react";
import axios from "axios";

const CreateGroup = () => {
  const [modalOpen, setModalOpen] = useState(false);
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

  const handleImageUpload = (event) => {
    const file = event.target.files[0];
    setGroupCreateForm((prevState) => ({
      ...prevState,
      image: file,
    }));
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    // Send the groupCreateForm data to the backend handler
    console.log(groupCreateForm);
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
          <AvatarUpdater onUploadSuccess={closeModal} />

          <button type="submit">Create</button>
        </form>
      </Modal>
      {/* <button onClick={openModal}>Upload New Image</button> */}
    </>
  );
};

export default CreateGroup;
