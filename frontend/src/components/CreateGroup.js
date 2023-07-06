import { useState } from "react";
import Modal from "./Modal";
import AvatarUpdater from "./AvatarUpdater.js";

const CreateGroup = () => {
  const [modalOpen, setModalOpen] = useState(false);
  const [groupCreateForm, setGroupCreateForm] = useState({
    Title: "",
    Description: "",
    inviteUsers: false,
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
    setGroupCreateForm((prevState) => ({
      ...prevState,
      [name]: value,
    }));
  };

  const handleInviteUsersChange = (event) => {
    const { checked } = event.target;
    setGroupCreateForm((prevState) => ({
      ...prevState,
      inviteUsers: checked,
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
              value={groupCreateForm.Title}
              onChange={handleChange}
            />
          </label>
          <br />
          <label>
            Description:
            <textarea
              name="description"
              value={groupCreateForm.Description}
              onChange={handleChange}
            ></textarea>
          </label>
          <br />
          <label>
            Invite Users:
            <input
              type="checkbox"
              name="inviteUsers"
              checked={groupCreateForm.inviteUsers}
              onChange={handleInviteUsersChange}
            />
          </label>
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
          <br />
          <button type="submit">Create</button>
        </form>
      </Modal>
      {/* <button onClick={openModal}>Upload New Image</button> */}
    </>
  );
};

export default CreateGroup;
