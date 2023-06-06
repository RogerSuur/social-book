import SingleChatItem from "../components/SingleChatItem";
import React, { useState, useEffect } from "react";
import useWebSocketConnection from "../hooks/useWebSocketConnection";
import Notification from "../components/Notification";
import axios from "axios";

const ChatTest = ({}) => {
  const [openChats, setOpenChats] = useState([]);

  console.log(openChats);

  const chatlist = [
    { first_name: "Chill", last_name: "Bill", userid: 1 },
    { first_name: "Scary", last_name: "Mary", userid: 2 },
  ];

  // Function to toggle a chat
  const toggleChat = (chatId) => {
    if (openChats.includes(chatId)) {
      setOpenChats(openChats.filter((id) => id !== chatId));
    } else {
      setOpenChats([...openChats, chatId]);
    }
  };

  const renderedChats = chatlist.map((chat, index) => (
    <li key={index}>
      <SingleChatItem
        index={index}
        chat={chat}
        toggleChat={toggleChat}
        isOpen={openChats.includes(index)}
      />
    </li>
  ));

  return (
    <div>
      <ul>{renderedChats}</ul>
    </div>
  );
};

export default ChatTest;
