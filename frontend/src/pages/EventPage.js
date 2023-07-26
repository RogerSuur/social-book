import React, { useState, useEffect } from "react";
import { useParams, useNavigate, Link } from "react-router-dom";
import { EVENT_URL } from "../utils/routes";
import axios from "axios";
import ImageHandler from "../utils/imageHandler";

const EventPage = () => {
  const [event, setEvent] = useState({});
  const [error, setError] = useState("");
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

  //   const groupMembersMap = groupMembers.map((member, index) => (
  //     <div key={index}>
  //       <Link to={`/profile/${member.id}`}>
  //         {ImageHandler(member.imagePath, "defaultuser.jpg", "profile-image")}
  //         <p>{member.nickname}</p>
  //       </Link>
  //     </div>
  //   ));

  const renderedEvent = (
    <div>
      <p>{event?.name}</p>
      <p>{event?.body}</p>
      <p>{new Date(event?.datetime).toLocaleString("et-EE")}</p>
    </div>
  );

  return <>{error ? <div>{error}</div> : <div>{renderedEvent}</div>}</>;
};

export default EventPage;
