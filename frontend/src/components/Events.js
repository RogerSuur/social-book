import { useState, useEffect } from "react";
import axios from "axios";
import { GROUP_EVENTS_URL } from "../utils/routes";
import CreateEvent from "./CreateEvent";

const Events = ({ groupId }) => {
  const [eventsData, setEventsData] = useState([]);

  const [loadNewEvents, setloadNewEvents] = useState(0);

  const handleEventUpdate = () => {
    setloadNewEvents((prevCount) => prevCount + 1);
  };

  useEffect(() => {
    const fetchEvents = async () => {
      try {
        console.log("REQUEST URL", GROUP_EVENTS_URL + groupId);
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

  const eventsDataMap = eventsData.map((event, index) => (
    <div key={index}>
      {/* LOO SIIA LINK EVENTI LEHELE */}
      {/* <Link to={`/profile/${member.Id}`}> */}
      <h1>{event.Title}</h1>
      <p>{event.Description}</p>
      <p>{event.EventTime}</p>
      {/* </Link> */}
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
