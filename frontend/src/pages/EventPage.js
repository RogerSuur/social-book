import React, { useState, useEffect } from "react";
import { useParams, useNavigate, Link } from "react-router-dom";
import { EVENT_URL } from "../utils/routes";
import axios from "axios";
import ImageHandler from "../utils/imageHandler";
import GroupSidebar from "../components/GroupSidebar";
import Modal from "../components/Modal.js";

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

  const handleResponse = async (attending) => {
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

  const handleModalClick = (attending) => {
    setActiveTab(attending);
    setModalOpen(true);
  };

  const countUsers = (attending) => {
    const userArray = event?.members?.map((member) => member.isAttending);

    return userArray?.reduce(
      (count, obj) => (obj === attending ? count + 1 : count),
      0
    );
  };

  console.log("ACTIVE", activeTab);
  console.log("EVENT: ", event);

  const renderedEvent = (
    <div className="event-mid">
      {/* <div className="event-stuff"> 
     {/* event img here 
    </div>*/}
      <button className="event-but" onClick={() => handleResponse(true)}>
        Going
      </button>
      <button className="event-but" onClick={() => handleResponse(false)}>
        Not going
      </button>
      <div className="event-stuff">
        <p>{event?.title}</p>
        <p>{event?.description}</p>
        <p>Start: {timeConverter(event?.eventTime)}</p>
        <p>End: {timeConverter(event?.eventEndTime)}</p>
        <Modal open={modalOpen} onClose={handleModalClose}>
          <ul>
            <li className="pepe" onClick={() => setActiveTab(true)}>
              Going
            </li>
            <li className="pepe" onClick={() => setActiveTab(false)}>
              Not Going
            </li>
          </ul>
          <div className="pepe">{userList(activeTab)}</div>
        </Modal>
        <button className="event-but" onClick={() => handleModalClick(true)}>
          Going {countUsers(true)}
        </button>
        <button className="event-but" onClick={() => handleModalClick(false)}>
          Not going {countUsers(false)}
        </button>
      </div>
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
