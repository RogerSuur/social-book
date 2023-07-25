import React from "react";
import { useOutletContext } from "react-router-dom";
import useWebSocketConnection from "../hooks/useWebSocketConnection";

const GroupRequestButton = ({ groupid }) => {
  const { socketUrl } = useOutletContext();
  const { sendJsonMessage } = useWebSocketConnection(socketUrl);

  const handleGroupRequest = () => {
    sendJsonMessage({
      type: "group_request",
      data: { group_id: groupid },
    });
  };

  return <button onClick={handleGroupRequest}>Join Group</button>;
};

export default GroupRequestButton;
