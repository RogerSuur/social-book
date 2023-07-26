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

  const eventsDataMap = eventsData.map((event, index) => (
    <div key={index}>
      <Link to={`/event/${event.Id}`}>
        <h1>{event.Title}</h1>
      </Link>
      <p>{event.Description}</p>
      <p>Begins {timeConverter(event.EventTime)}</p>
    </div>
  ));

  return (
    <>
      <p>Here comes the events of the group</p>
      {eventsDataMap}
      <p>Here comes creating an event</p>
      <CreateEvent onEventCreated={handleEventUpdate} id={groupId} />
    </>
  );
};

export default Events;
