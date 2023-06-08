import { WS_URL } from "../utils/routes";
import useWebSocketConnection from "../hooks/useWebSocketConnection";
import React, { useState } from "react";

const Chatbox = ({ toggleChat, chat }) => {
  const [messageHistory, setMessageHistory] = useState([]);
  const { sendJsonMessage, lastJsonMessage } = useWebSocketConnection(WS_URL);
  const [message, setMessage] = useState({
    type: "message",
    data: {
      body: "",
    },
  });

  console.log(message, "MESS");

  const closeChat = () => {
    toggleChat(0);
  };

  const handleChange = (event) => {
    const { value } = event.target;

    setMessage((prevMessage) => {
      return {
        ...prevMessage,
        data: { body: value },
      };
    });
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    if (chat?.group_id > 0) {
      console.log({
        ...message,
        data: { ...message.data, recepient_id: chat.group_id },
      });
      // sendJsonMessage({
      //   ...message,
      //   data: { ...message.data, recepient_id: chat.group_id },
      // });
    } else {
      console.log(
        {
          ...message,
          data: { ...message.data, recepient_id: chat.userid },
        },
        "SENDING"
      );
      // sendJsonMessage({
      //   ...message,
      //   data: { ...message.data, recepient_id: chat.userid },
      // });
    }
    setMessage({ ...message, data: { body: "" } });
  };

  const chatbox = (
    <div className="chatbox">
      <div className="chat-title">
        {chat.first_name} {chat.last_name}
        <button onClick={closeChat}>Close</button>
      </div>
      <div className="message-history">Messages</div>
      <div className="message-box">
        <form onSubmit={handleSubmit}>
          <input
            type="text"
            placeholder="Message"
            onChange={handleChange}
            name="message"
            value={message.data.body}
            autoFocus
            required
          />
          <button>Send</button>
        </form>
      </div>
    </div>
  );

  return <>{chatbox}</>;
};

export default Chatbox;
