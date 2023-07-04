import { WS_URL } from "../utils/routes";
import useWebSocketConnection from "../hooks/useWebSocketConnection";
import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";

const Chatbox = ({ toggleChat, chat, user }) => {
  const [messageHistory, setMessageHistory] = useState([]);
  const { sendJsonMessage, lastJsonMessage } = useWebSocketConnection(WS_URL);
  const [message, setMessage] = useState({
    type: "message",
    data: {
      body: "",
    },
  });

  const defaultImage = () =>
    chat.user_id ? "defaultuser.jpg" : "defaultgroup.png";

  const imageHandler = () => {
    const source = chat?.avatarImage
      ? `images/${chat.id}/${chat.avatarImage}`
      : defaultImage();

    const image = (
      <img
        style={{
          width: "20px",
          height: "20px",
        }}
        src={source}
      ></img>
    );
    return image;
  };

  const loadMessages = () => {
    sendJsonMessage({
      type: "request_message_history",
      data: { id: chat.user_id, group_id: chat.group_id },
    });
  };

  useEffect(() => {
    loadMessages();
  }, []);

  useEffect(() => {
    switch (lastJsonMessage?.type) {
      case "message_history":
        if (lastJsonMessage?.data.length > 0) {
          setMessageHistory((prevMessageHistory) => [
            lastJsonMessage?.data?.messages,
            ...prevMessageHistory,
          ]);
        }

        break;
      case "message":
        if (
          (lastJsonMessage?.data?.sender_id === chat.user_id &&
            lastJsonMessage?.data?.group_id === 0) ||
          lastJsonMessage?.data?.recipient_id === chat.user_id ||
          lastJsonMessage?.data?.group_id === chat.group_id
        ) {
          setMessageHistory((prevMessageHistory) => [
            ...prevMessageHistory,
            lastJsonMessage?.data,
          ]);
        }
    }
  }, [lastJsonMessage]);

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

  console.log(messageHistory, "HISTORY");

  const getTime = (datetime) =>
    datetime.toLocaleTimeString([], {
      hour: "2-digit",
      minute: "2-digit",
      hour12: false,
    });

  const renderedMessages = messageHistory.map((msg) => {
    if (msg) {
      switch (msg.sender_id) {
        case user:
          return (
            <p key={msg.id} className="own-message">
              {getTime(msg.timestamp)} {msg.body}
            </p>
          );
        default:
          return (
            <p key={msg.id} className="message">
              {msg.body} {getTime(msg.timestamp)}
            </p>
          );
      }
    }
  });

  const handleSubmit = (event) => {
    event.preventDefault();
    let msg = {
      ...message,
      data: { ...message.data, sender_id: user, recipient_id: 0, group_id: 0 },
    };
    if (chat?.group_id > 0) {
      msg.data.group_id = chat.group_id;
      sendJsonMessage(msg);
    } else {
      msg.data.recipient_id = chat.user_id;
      sendJsonMessage(msg);
    }
    setMessageHistory((prevMessageHistory) => [
      ...prevMessageHistory,
      { ...msg.data, timestamp: new Date() },
    ]);
    setMessage({ ...message, data: { body: "" } });
  };

  const chatName =
    chat?.user_id > 0 ? (
      <Link to={`/profile/${chat.user_id}`}>{chat.name}</Link>
    ) : (
      <Link to={`/groups/${chat.group_id}`}>{chat.name}</Link>
    );

  const chatbox = (
    <div className="chatbox">
      <div className="chat-title">
        {chatName}
        {imageHandler()}
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
