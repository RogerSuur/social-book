import { useState, useEffect } from "react";
import axios from "axios";
import { Link } from "react-router-dom";
import { GROUP_EVENTS_URL } from "../utils/routes";
import CreateEvent from "./CreateEvent";
import GenericEventList from "./GenericEventList";
import { ListGroup } from "react-bootstrap";
import GenericModal from "./GenericModal";
import { ShortDatetime } from "../utils/datetimeConverters";

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

  const eventsMap = events?.map((event, index) => (
    <li key={index}>
      <Link to={`/event/${event.id}`}>
        <h1>{event.title}</h1>
      </Link>
      <p>{event.description}</p>
      <p>Begins {ShortDatetime(event.eventTime)}</p>
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
