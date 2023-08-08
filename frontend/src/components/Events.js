import { useState, useEffect } from "react";
import axios from "axios";
import { Link } from "react-router-dom";
import { GROUP_EVENTS_URL } from "../utils/routes";
import CreateEvent from "./CreateEvent";

const Events = ({ groupId }) => {
  const [eventsData, setEventsData] = useState([]);

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
        setEventsData(response.data);
      } catch (err) {
        console.error(err);
      }
    };
    fetchEvents();
  }, [groupId, loadNewEvents]);

  console.log("GROUPEVENTS: ", eventsData);

  const timeConverter = (datetime) =>
    new Date(datetime).toLocaleTimeString("en-UK", {
      month: "short",
      day: "2-digit",
      year: "2-digit",
      hour: "numeric",
      minute: "2-digit",
    });

  const eventsDataMap = eventsData?.map((event, index) => (
    <li className="pepe">
    <li className="dif-link" key={index}>
      <Link to={`/event/${event.id}`}>
        <h1>{event.title}</h1>
      </Link>
      <p>{event.description}</p>
      <p>Begins {timeConverter(event.eventTime)}</p>
    </li>
    </li>
  ));

  return (
    <>
      <p>Here comes the events of the group</p>
      <ul>{eventsDataMap}</ul>
      <p>Create an event</p>
      <CreateEvent onEventCreated={handleEventUpdate} id={groupId} />
    </>
  );
};

export default Events;
