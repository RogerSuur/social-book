import SingleChatItem from "../components/SingleChatItem";
import React, { useState, useEffect } from "react";
import { useOutletContext } from "react-router-dom";
import useWebSocketConnection from "../hooks/useWebSocketConnection";
import Notification from "../components/Notification";
import axios from "axios";
import { WS_URL } from "../utils/routes";

const ChatTest = ({}) => {
  const [openChat, setOpenChat] = useState(0);
  const { sendJsonMessage, lastJsonMessage } = useWebSocketConnection(WS_URL);

  console.log(openChat);

  const chatlist = [
    { first_name: "Chill", last_name: "Bill", userid: 1 },
    { first_name: "Scary", last_name: "Mary", userid: 2 },
    { first_name: "Scary", last_name: "Mary", userid: 3 },
    { first_name: "Scary", last_name: "Mary", userid: 4 },
    { first_name: "Scary", last_name: "Mary", userid: 5 },
    { first_name: "Scary", last_name: "Mary", userid: 6 },
    { first_name: "Scary", last_name: "Mary", userid: 7 },
    { first_name: "Scary", last_name: "Mary", userid: 8 },
    { first_name: "Scary", last_name: "Mary", userid: 9 },
    { first_name: "Scary", last_name: "Mary", userid: 10 },
    { first_name: "Scary", last_name: "Mary", userid: 11 },
    { first_name: "Scary", last_name: "Mary", userid: 12 },
    { first_name: "Scary", last_name: "Mary", userid: 13 },
    { first_name: "Scary", last_name: "Mary", userid: 14 },
    { first_name: "Scary", last_name: "Mary", userid: 15 },
    { first_name: "Scary", last_name: "Mary", userid: 16 },
    { first_name: "Scary", last_name: "Mary", userid: 17 },
    { first_name: "Scary", last_name: "Mary", userid: 18 },
    { first_name: "Scary", last_name: "Mary", userid: 19 },
    { first_name: "Scary", last_name: "Mary", userid: 20 },
    { first_name: "Scary", last_name: "Mary", userid: 21 },
    { first_name: "Scary", last_name: "Mary", userid: 22 },
    { first_name: "Scary", last_name: "Mary", userid: 22 },
  ];

  // Function to toggle a chat
  const toggleChat = (chatId) => {
    if (openChat !== chatId) {
      setOpenChat(chatId);
    }
  };

  const renderedChats = chatlist.map((chat, index) => (
    <li key={index + 1}>
      <SingleChatItem
        index={index + 1}
        chat={chat}
        toggleChat={toggleChat}
        isOpen={openChat === index + 1}
      />
    </li>
  ));

  return (
    <div className="chat-sidebar">
      <ul>{renderedChats}</ul>
    </div>
  );
};

export default ChatTest;
