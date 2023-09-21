import React from "react";
import { Col } from "react-bootstrap";

const ChatMessage = ({ msg, own }) => {
  const getTime = (datetime) =>
    new Date(datetime).toLocaleTimeString([], {
      hour: "2-digit",
      minute: "2-digit",
      hour12: false,
    });

  const message = own ? (
    <Col md={{ span: 9, offset: 3 }} variant="success" className="text-justify">
      <div className="border-success rounded bg-success">
        <span className="p-1">{msg.body}</span>
        <span className="text-secondary text-muted small text-end ps-2">
          {getTime(msg.timestamp)}
        </span>
      </div>
    </Col>
  ) : (
    <Col md="9" className="message text-start m-0 p-0">
      <div className="border rounded bg-primary">
        <span>
          {msg.group_id > 0 && msg.sender_name} {msg.body}
          <span className="own-time text-secondary text-muted small text-end ps-2">
            {getTime(msg.timestamp)}
          </span>
        </span>
      </div>
    </Col>
  );
  return message;
};

export default ChatMessage;
