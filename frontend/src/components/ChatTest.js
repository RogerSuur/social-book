import SingleChatlistItem from "../components/SingleChatlistItem";
import React, { useState, useEffect } from "react";
import useWebSocketConnection from "../hooks/useWebSocketConnection";
import Chatbox from "../components/Chatbox";
import MessageNotification from "../components/MessageNotification";
import { WS_URL } from "../utils/routes";

const ChatTest = ({}) => {
  const [openChat, setOpenChat] = useState(0);
  const [chatlist, setChatlist] = useState([]);
  const { sendJsonMessage, lastJsonMessage } = useWebSocketConnection(WS_URL);

  const loadChatlist = () => {
    sendJsonMessage({
      type: "request_chatlist",
    });
  };

  useEffect(() => {
    loadChatlist();
  }, []);

  const chatter = [
    { first_name: "Chill", last_name: "Bill", userid: 1 },
    { first_name: "Scary", last_name: "Mary", userid: 2 },
    { first_name: "B", last_name: "Mary", username: "bloodymary", userid: 3 },
    { first_name: "A", last_name: "Mary", userid: 4 },
    { first_name: "V", last_name: "Mary", userid: 5 },
    { first_name: "C", last_name: "Mary", userid: 6 },
    { first_name: "D", last_name: "Mary", userid: 7 },
    { first_name: "E", last_name: "Mary", userid: 8 },
    { first_name: "F", last_name: "Mary", userid: 9 },
    { first_name: "G", last_name: "Mary", userid: 10 },
    { first_name: "H", last_name: "Mary", userid: 11 },
    { first_name: "I", last_name: "Mary", userid: 12 },
    { first_name: "J", last_name: "Mary", userid: 13 },
    { first_name: "K", last_name: "Mary", userid: 14 },
    { first_name: "L", last_name: "Mary", userid: 15 },
    { first_name: "M", last_name: "Mary", userid: 16 },
    { first_name: "N", last_name: "Mary", userid: 17 },
    { first_name: "O", last_name: "Mary", userid: 18 },
    { first_name: "P", last_name: "Mary", userid: 19 },
    { first_name: "S", last_name: "Mary", userid: 20 },
    { first_name: "T", last_name: "Mary", userid: 21 },
    { first_name: "U", last_name: "Mary", userid: 22 },
    { first_name: "V", last_name: "Mary", userid: 22 },
  ];

  const toggleChat = (chatId) => {
    if (openChat !== chatId) {
      setOpenChat(chatId);
    }
  };

  useEffect(() => {
    switch (lastJsonMessage?.type) {
      case "chatlist":
        setChatlist([...lastJsonMessage.data]);
        break;
      case "message":
        break;
    }
  }, [lastJsonMessage]);

  const renderedChats = chatter.map((chat, index) => (
    <>
      <li key={index + 1}>
        <SingleChatlistItem
          userid={chat.userid}
          chat={chat}
          toggleChat={toggleChat}
        />
      </li>
      {openChat === chat.userid && (
        <Chatbox toggleChat={toggleChat} chat={chat} />
      )}
    </>
  ));

  return (
    <>
      <MessageNotification />
      <div className="chat-sidebar">
        <ul>{renderedChats}</ul>
      </div>
    </>
  );
};

export default ChatTest;
