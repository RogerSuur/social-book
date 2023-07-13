import { useState, useEffect } from "react";
import axios from "axios";
import { GROUP_EVENTS_URL } from "../utils/routes";

const Events = ({ groupId }) => {
  const [eventsData, setEventsData] = useState([]);

  useEffect(() => {
    const fetchEvents = async () => {
      try {
        console.log("REQUEST URL", GROUP_EVENTS_URL + groupId);
        const response = await axios.get(GROUP_EVENTS_URL + groupId, {
          withCredentials: true,
        });
        setEventsData(response.data);
        console.log(response.data);
      } catch (err) {
        console.error(err);
      }
    };
    fetchEvents();
  }, [groupId]);

  const eventsDataMap = eventsData.map((event, index) => (
    <div key={index}>
      {console.log(event)}
      {/* LOO SIIA LINK EVENTI LEHELE */}
      {/* <Link to={`/profile/${member.Id}`}> */}
      <h1>{event.title}</h1>
      <p>{event.description}</p>
      <p>{event.event_time}</p>
      {/* </Link> */}
    </div>
  ));

  return (
    <>
      <p>Here comes the events of the group</p>
      {eventsDataMap}
    </>
  );
};

export default Events;
