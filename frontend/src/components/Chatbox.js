import { WS_URL } from "../utils/routes";
import useWebSocketConnection from "../hooks/useWebSocketConnection";
import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";

const Chatbox = ({ toggleChat, chat }) => {
  const [messageHistory, setMessageHistory] = useState([]);
  const { sendJsonMessage, lastJsonMessage } = useWebSocketConnection(WS_URL);
  const [message, setMessage] = useState({
    type: "message",
    data: {
      body: "",
    },
  });

  const loadMessages = () => {
    sendJsonMessage({
      type: "request_message_history",
      data: { userid: chat.userid },
    });
  };

  useEffect(() => {
    loadMessages();
  }, []);

  useEffect(() => {
    if (lastJsonMessage && lastJsonMessage.type === "message_history") {
      setMessageHistory((prevMessageHistory) => [
        ...lastJsonMessage?.data?.messages,
        ...prevMessageHistory,
      ]);
    }
  }, [lastJsonMessage]);

  console.log(chat.userid, "CHATID");

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

  const renderedMessages = <p>Message</p>;

  const handleSubmit = (event) => {
    event.preventDefault();
    if (chat?.group_id > 0) {
      console.log({
        ...message,
        data: { ...message.data, recipient_id: chat.group_id },
      });
      // sendJsonMessage({
      //   ...message,
      //   data: { ...message.data, recipient_id: chat.group_id },
      // });
    } else {
      console.log(
        {
          ...message,
          data: { ...message.data, recipient_id: chat.userid },
        },
        "SENDING"
      );
      // sendJsonMessage({
      //   ...message,
      //   data: { ...message.data, recipient_id: chat.userid },
      // });
    }
    setMessage({ ...message, data: { body: "" } });
  };

  const chatbox = (
    <div className="chatbox">
      <div className="chat-title">
        <Link to={`/profile/${chat.userid}`}>
          {chat.first_name} {chat.last_name}
        </Link>
        <button onClick={closeChat}>Close</button>
      </div>
      <div className="message-history">{renderedMessages}</div>
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
