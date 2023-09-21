import { useState, useEffect } from "react";
import axios from "axios";
import { Link } from "react-router-dom";
import { GROUP_EVENTS_URL } from "../utils/routes";
import CreateEvent from "./CreateEvent";
import GenericEventList from "./GenericEventList";
import { ListGroup } from "react-bootstrap";
import GenericModal from "./GenericModal";

const Events = ({ groupId }) => {
  const [events, setEvents] = useState([]);

  const [loadNewEvents, setLoadNewEvents] = useState(0);

  const handleEventUpdate = () => {
    setLoadNewEvents((prevCount) => prevCount + 1);
  };

  useEffect(() => {
    const fetchEvents = async () => {
      try {
        const response = await axios.get(GROUP_EVENTS_URL + groupId, {
          withCredentials: true,
        });
        setEvents(response.data);
      } catch (err) {
        console.error(err);
      }
    };
    fetchEvents();
  }, [groupId, loadNewEvents]);

  const timeConverter = (datetime) =>
    new Date(datetime).toLocaleTimeString("en-UK", {
      month: "short",
      day: "2-digit",
      year: "2-digit",
      hour: "numeric",
      minute: "2-digit",
    });

  const eventsMap = events?.map((event, index) => (
    <li key={index}>
      <Link to={`/event/${event.id}`}>
        <h1>{event.title}</h1>
      </Link>
      <p>{event.description}</p>
      <p>Begins {timeConverter(event.eventTime)}</p>
    </li>
  ));

  return (
    <>
      <ListGroup>
        <GenericEventList url={GROUP_EVENTS_URL + groupId} />
      </ListGroup>
      <GenericModal buttonText="Create an event">
        <CreateEvent onEventCreated={handleEventUpdate} id={groupId} />
      </GenericModal>
    </>
  );
};

export default Events;
