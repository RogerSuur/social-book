import React from "react";
import { Col } from "react-bootstrap";
import { ShortTime } from "../utils/datetimeConverters.js";

const ChatMessage = ({ msg, own }) => {
  const time = (
    <span className="small text-end ps-2">{ShortTime(msg.timestamp)}</span>
  );

  const message = own ? (
    <div className="own-message">
      <div className="message bg-primary text-light text-end ">
        {msg.body}
        {time}
      </div>
    </div>
  ) : (
    <>
      {msg.group_id > 0 && (
        <p className="m-0 small text-muted">{msg.sender_name}</p>
      )}
      <div className="message bg-secondary-subtle">
        {msg.body}
        {time}
      </div>
    </>
  );

  return message;
};

export default ChatMessage;
