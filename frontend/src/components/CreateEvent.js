import Modal from "./Modal";
import { useState } from "react";
import axios from "axios";

const CreateEvent = ({ onEventCreated, id }) => {
  const [modalOpen, setModalOpen] = useState(false);
  const [errMsg, setErrMsg] = useState("");
  const [createEventForm, setcreateEventForm] = useState({
    title: "",
    description: "",
    startTime: "",
    endTime: "",
    group_id: +id,
  });

  const openModal = () => {
    setModalOpen(true);
  };

  const closeModal = () => {
    setModalOpen(false);
  };

  const handleChange = (event) => {
    const { name, value } = event.target;
    setcreateEventForm((prevState) => ({
      ...prevState,
      [name]: value,
    }));
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    // Send the createEventForm data to the backend handler
    try {
      const response = await axios.post(
        "http://localhost:8000/creategroupevent",
        JSON.stringify(createEventForm),
        { withCredentials: true },
        {
          headers: { "Content-Type": "application/json" },
        }
      );
      setErrMsg(response.data?.message);
      onEventCreated();
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
          {errMsg && <p className="error">{errMsg}</p>}
          <form className="pop-form" onSubmit={handleSubmit}>
            Title:
            <label className="input-big">
              <input
                type="text"
                name="title"
                value={createEventForm.title}
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
                value={createEventForm.description}
                onChange={handleChange}
                required
              ></textarea>
            </label>
            <br />
            Start Time:
            <label>
              <input
                type="datetime-local"
                name="startTime"
                value={createEventForm.startTime}
                onChange={handleChange}
                required
              />
            </label>
            <br />
            End Time:
            <label>
              <input
                type="datetime-local"
                name="endTime"
                value={createEventForm.endTime}
                onChange={handleChange}
                required
              />
            </label>
            <br />
            <button className="create-but" type="submit">
              Create
            </button>
          </form>
        </Modal>
      </div>
    </>
  );
};

export default CreateEvent;
