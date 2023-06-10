import SingleChatlistItem from "../components/SingleChatlistItem";
import React, { useState, useEffect } from "react";
import useWebSocketConnection from "../hooks/useWebSocketConnection";
import Chatbox from "../components/Chatbox";
import MessageNotification from "../components/MessageNotification";
import { WS_URL } from "../utils/routes";

const ChatTest = ({}) => {
  const [openChat, setOpenChat] = useState([0, 0]);
  const [chatlist, setChatlist] = useState([]);
  const [user, setUser] = useState(0);
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
    { name: "Chill Bill", userid: 1 },
    { name: "Scary Mary", userid: 2 },
    { name: "bloodymary", userid: 3 },
    { name: "Best group", userid: 0, group_id: 4 },
    { name: "V Mary", userid: 5 },
    { name: "C Mary", userid: 6 },
    { name: "D Mary", userid: 7 },
    { name: "E Mary", userid: 8 },
    { name: "Better group", userid: 0, group_id: 2 },
    { name: "G Mary", userid: 10 },
    { name: "H Mary", userid: 11 },
    { name: "I Mary", userid: 12 },
    { name: "J Mary", userid: 13 },
    { name: "K Mary", userid: 14 },
    { name: "L Mary", userid: 15 },
    { name: "M Mary", userid: 16 },
    { name: "N Mary", userid: 17 },
    { name: "O Mary", userid: 18 },
    { name: "P Mary", userid: 19 },
    { name: "S Mary", userid: 20 },
    { name: "T Mary", userid: 21 },
    { name: "U Mary", userid: 22 },
    { name: "V Mary", userid: 22 },
  ];

  const toggleChat = (chatId) => {
    if (openChat !== chatId) {
      console.log(chatId, "chatID");
      setOpenChat(chatId);
    }
  };

  const useriddd = 1;

  const checkOpenChat = (open) =>
    open.every((value, index) => value === openChat[index]);

  useEffect(() => {
    switch (lastJsonMessage?.type) {
      case "chatlist":
        setChatlist([...lastJsonMessage?.data?.users]);
        setUser(...lastJsonMessage?.data?.userid);
        break;
      case "message":
        break;
    }
  }, [lastJsonMessage]);

  const renderedChats = chatter.map((chat, index) => (
    <>
      <li key={index + 1}>
        <SingleChatlistItem chat={chat} toggleChat={toggleChat} />
      </li>
      {(checkOpenChat([chat?.userid, 0]) ||
        checkOpenChat([0, chat?.group_id])) && (
        <Chatbox toggleChat={toggleChat} chat={chat} user={useriddd} />
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
