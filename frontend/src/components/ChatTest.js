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

  const toggleChat = (chatId) => {
    if (openChat !== chatId) {
      setOpenChat(chatId);
    }
  };

  const checkOpenChat = (open) =>
    open.every((value, index) => {
      return value === openChat[index];
    });

  useEffect(() => {
    switch (lastJsonMessage?.type) {
      case "chatlist":
        setChatlist([...lastJsonMessage?.data?.chatlist]);
        setUser(lastJsonMessage?.data?.user_id);
        break;
      case "message":
        break;
    }
  }, [lastJsonMessage]);

  const renderedChats = chatlist.map((chat, index) => (
    <>
      <li key={index + 1}>
        <SingleChatlistItem chat={chat} toggleChat={toggleChat} />
      </li>
      {checkOpenChat([chat?.user_id, chat?.group_id]) && (
        <>
          <Chatbox toggleChat={toggleChat} chat={chat} user={user} />
        </>
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
