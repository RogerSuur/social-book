import React, { useState, useEffect } from "react";
import { useParams, useNavigate, Link } from "react-router-dom";
import { EVENT_URL } from "../utils/routes";
import axios from "axios";
import ImageHandler from "../utils/imageHandler";
import GroupSidebar from "../components/GroupSidebar";
import Modal from "../components/Modal.js";
import { set } from "react-hook-form";

const EventPage = () => {
  const [event, setEvent] = useState({});
  const [error, setError] = useState("");
  const [modalOpen, setModalOpen] = useState(false);
  const [activeTab, setActiveTab] = useState(true);
  const navigate = useNavigate();
  const { id } = useParams();

  useEffect(() => {
    const loadEvent = async () => {
      try {
        await axios
          .get(EVENT_URL + id, {
            withCredentials: true,
          })
          .then((response) => {
            console.log("RESP: ", response.data);
            setEvent(response.data);
          });
      } catch (err) {
        if (!err?.response) {
          setError("No Server Response");
        } else if (err.response?.status === 404) {
          navigate("404", { replace: true });
        } else if (err.response?.status > 200) {
          setError("Internal Server Error");
        }
      }
    };
    loadEvent();
  }, [id]);

  console.log("EVENT DATA: ", event);

  const userList = (attendance) => {
    const users = event?.members?.filter(
      (member) => member.isAttending === attendance
    );

    return users?.map((member, index) => (
      <li key={index}>
        <Link to={`/profile/${member.userId}`}>
          {ImageHandler(member.imagePath, "defaultuser.jpg", "profile-image")}
          <p>{member.name}</p>
        </Link>
      </li>
    ));
  };

  const timeConverter = (datetime) =>
    new Date(datetime).toLocaleTimeString("en-UK", {
      month: "short",
      day: "2-digit",
      year: "2-digit",
      hour: "numeric",
      minute: "2-digit",
    });

  const handleModalClose = () => {
    setModalOpen(false);
  };

  const handleModalClick = () => {
    setModalOpen(true);
  };

  console.log("ACTIVE", activeTab);

  const renderedEvent = (
    <div>
      <p>{event?.title}</p>
      <p>{event?.description}</p>
      <p>Start: {timeConverter(event?.eventTime)}</p>
      <p>End: {timeConverter(event?.eventEndTime)}</p>
      <Modal open={modalOpen} onClose={handleModalClose}>
        <ul>
          <li onClick={() => setActiveTab(true)}>Going</li>
          <li onClick={() => setActiveTab(false)}>Not Going</li>
        </ul>
        <ul>{userList(activeTab)}</ul>
      </Modal>
      <button onClick={handleModalClick}>See who is going</button>
    </div>
  );

  return (
    <>
      <GroupSidebar />
      {error ? <div>{error}</div> : <div>{renderedEvent}</div>}
    </>
  );
};

export default EventPage;
