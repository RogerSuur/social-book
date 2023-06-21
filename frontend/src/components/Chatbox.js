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

  console.log(chat);

  console.log(messageHistory);

  const defaultImage = () =>
    chat.userid ? "defaultuser.jpg" : "defaultgroup.png";

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

  const sms = [
    {
      id: 1, //message id
      sender_id: 2, // 0 if group
      sender_name: "Scary Mary", // either a  username (if exists) or firstname and lastname
      recipient_id: 1, // 0 if group
      recipient_name: "AnnieA", // either a username (if   exists) or firstname and lastname && empty if     group
      group_id: 0, // 0 if user
      group_name: "", //empty if user
      body: "message1",
      timestamp: "2023-06-05 16:01:00.303095707 +03:00",
    },
    {
      id: 2, //message id
      sender_id: 1, // 0 if group
      sender_name: "Scary Mary", // either a  username (if exists) or firstname and lastname
      recipient_id: 2, // 0 if group
      recipient_name: "AnnieA", // either a username (if   exists) or firstname and lastname && empty if     group
      group_id: 0, // 0 if user
      group_name: "", //empty if user
      body: "message2",
      timestamp: "2023-06-05 16:01:01.303095707 +03:00",
    },
    {
      id: 3, //message id
      sender_id: 2, // 0 if group
      sender_name: "Scary Mary", // either a  username (if exists) or firstname and lastname
      recipient_id: 1, // 0 if group
      recipient_name: "AnnieA", // either a username (if   exists) or firstname and lastname && empty if     group
      group_id: 0, // 0 if user
      group_name: "", //empty if user
      body: "message3",
      timestamp: "2023-06-05 16:01:02.303095707 +03:00",
    },
    {
      id: 4, //message id
      sender_id: 1, // 0 if group
      sender_name: "Scary Mary", // either a  username (if exists) or firstname and lastname
      recipient_id: 2, // 0 if group
      recipient_name: "AnnieA", // either a username (if   exists) or firstname and lastname && empty if     group
      group_id: 0, // 0 if user
      group_name: "", //empty if user
      body: "message4",
      timestamp: "2023-06-05 16:01:03.303095707 +03:00",
    },
    {
      id: 5, //message id
      sender_id: 2, // 0 if group
      sender_name: "Scary Mary", // either a  username (if exists) or firstname and lastname
      recipient_id: 1, // 0 if group
      recipient_name: "AnnieA", // either a username (if   exists) or firstname and lastname && empty if     group
      group_id: 0, // 0 if user
      group_name: "", //empty if user
      body: "message5",
      timestamp: "2023-06-05 16:01:04.303095707 +03:00",
    },
    {
      id: 6, //message id
      sender_id: 2, // 0 if group
      sender_name: "Scary Mary", // either a  username (if exists) or firstname and lastname
      recipient_id: 1, // 0 if group
      recipient_name: "AnnieA", // either a username (if   exists) or firstname and lastname && empty if     group
      group_id: 0, // 0 if user
      group_name: "", //empty if user
      body: "message6",
      timestamp: "2023-06-05 16:01:05.303095707 +03:00",
    },
    {
      id: 7, //message id
      sender_id: 1, // 0 if group
      sender_name: "Scary Mary", // either a  username (if exists) or firstname and lastname
      recipient_id: 2, // 0 if group
      recipient_name: "AnnieA", // either a username (if   exists) or firstname and lastname && empty if     group
      group_id: 0, // 0 if user
      group_name: "", //empty if user
      body: "message7",
      timestamp: "2023-06-05 16:01:06.303095707 +03:00",
    },
    {
      id: 8, //message id
      sender_id: 1, // 0 if group
      sender_name: "Scary Mary", // either a  username (if exists) or firstname and lastname
      recipient_id: 2, // 0 if group
      recipient_name: "AnnieA", // either a username (if   exists) or firstname and lastname && empty if     group
      group_id: 0, // 0 if user
      group_name: "", //empty if user
      body: "message8",
      timestamp: "2023-06-05 16:01:07.303095707 +03:00",
    },
    {
      id: 9, //message id
      sender_id: 1, // 0 if group
      sender_name: "Scary Mary", // either a  username (if exists) or firstname and lastname
      recipient_id: 2, // 0 if group
      recipient_name: "AnnieA", // either a username (if   exists) or firstname and lastname && empty if     group
      group_id: 0, // 0 if user
      group_name: "", //empty if user
      body: "message9",
      timestamp: "2023-06-05 16:01:08.303095707 +03:00",
    },
    {
      id: 10, //message id
      sender_id: 2, // 0 if group
      sender_name: "Scary Mary", // either a  username (if exists) or firstname and lastname
      recipient_id: 1, // 0 if group
      recipient_name: "AnnieA", // either a username (if   exists) or firstname and lastname && empty if     group
      group_id: 0, // 0 if user
      group_name: "", //empty if user
      body: "message10",
      timestamp: "2023-06-05 16:01:09.303095707 +03:00",
    },
  ];

  const loadMessages = () => {
    sendJsonMessage({
      type: "request_message_history",
      data: { id: chat.userid, group_id: chat.group_id },
    });
  };

  useEffect(() => {
    loadMessages();
  }, []);

  useEffect(() => {
    switch (lastJsonMessage?.type) {
      case "message_history":
        setMessageHistory((prevMessageHistory) => [
          ...lastJsonMessage?.data?.messages,
          ...prevMessageHistory,
        ]);
        break;
      case "message":
        if (
          (lastJsonMessage?.data?.sender_id === chat.userid &&
            lastJsonMessage?.data?.group_id === 0) ||
          lastJsonMessage?.data?.recipient_id === chat.userid ||
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

  const renderedMessages = sms.map((msg) => {
    switch (msg.sender_id) {
      case user:
        return <p className="own-message">{msg.body}</p>;
      default:
        return <p className="message">{msg.body}</p>;
    }
  });

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

  const chatName =
    chat?.userid > 0 ? (
      <Link to={`/profile/${chat.userid}`}>{chat.name}</Link>
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
