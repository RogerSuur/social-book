import Modal from "./Modal";
import AvatarUpdater from "./AvatarUpdater.js";
import { useState, useEffect } from "react";
import axios from "axios";
import GenericGroupList from "./GenericGroupList";
import { USER_CREATED_GROUPS_URL } from "../utils/routes";

const CreateGroup = () => {
  const [modalOpen, setModalOpen] = useState(false);
  const [errMsg, setErrMsg] = useState("");
  const [groupCreateForm, setGroupCreateForm] = useState({
    title: "",
    description: "",
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

  const handleSubmit = async (event) => {
    event.preventDefault();
    // Send the groupCreateForm data to the backend handler
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
      <i className="iconoir-add-circle" onClick={openModal} />
      <Modal open={modalOpen} onClose={closeModal}>
        <form className="form-style" onSubmit={handleSubmit}>
          <div className="popup-group">
            Title:
            <input
              type="text"
              name="title"
              value={groupCreateForm.title}
              onChange={handleChange}
              required
            />
          </div>
          <br />
          <div>
            Description:
            <textarea
              className="group-textarea"
              name="description"
              value={groupCreateForm.description}
              onChange={handleChange}
              required
            ></textarea>
          </div>
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

          <button className="group-button" type="submit">Create</button>
        </form>
      </Modal>
      {/* <button onClick={openModal}>Upload New Image</button> */}
    </>
  );
};

export default CreateGroup;
