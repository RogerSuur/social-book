import Modal from "./Modal";
import { useState, useEffect } from "react";
import axios from "axios";

const CreateGroup = ({ onGroupCreated }) => {
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
      onGroupCreated();
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
      <div className="newModal">
        <i className="iconoir-add-circle" onClick={openModal} />
        <Modal open={modalOpen} onClose={closeModal}>
          <form className="pop-form" onSubmit={handleSubmit}>
            Title:
            <label className="input-big">
              <input
                type="text"
                name="title"
                value={groupCreateForm.title}
                onChange={handleChange}
                required
              />
            </label>
            <br />
            Description:
            <label>
              <textarea
                className="text-big"
                name="description"
                value={groupCreateForm.description}
                onChange={handleChange}
                required
              ></textarea>
            </label>
            <br />

            <button className="create-but" type="submit">Create</button>
          </form>
        </Modal>
      </div>
      {/* <button onClick={openModal}>Upload New Image</button> */}
    </>
  );
};

export default CreateGroup;
