import SingleChatlistItem from "./SingleChatlistItem";
import React, { useState, useEffect } from "react";
import useWebSocketConnection from "../hooks/useWebSocketConnection";
import Chatbox from "./Chatbox";
import MessageNotification from "./MessageNotification";
import { WS_URL } from "../utils/routes";

const Chat = () => {
  const [openChat, setOpenChat] = useState(null);
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

  const toggleChat = (chat) => {
    if (!chat) {
      setOpenChat(null);
    } else if (
      openChat?.user_id !== chat.user_id ||
      openChat?.group_id !== chat.group_id
    ) {
      setOpenChat(chat);
    }
  };

  const checkChat = (open, checker) =>
    open.every((value, index) => {
      return value === checker[index];
    });

  useEffect(() => {
    switch (lastJsonMessage?.type) {
      case "chatlist":
        setChatlist([...lastJsonMessage?.data?.chatlist]);
        setUser(lastJsonMessage?.data?.user_id);
        break;
      case "message":
        const chatToFind = [
          lastJsonMessage?.data?.group_id > 0
            ? 0
            : lastJsonMessage?.data?.sender_id,
          lastJsonMessage?.data?.group_id,
        ];

        const userChat = chatlist.find((chat) =>
          checkChat([chat?.user_id, chat?.group_id], chatToFind)
        );

        if (!userChat) {
          const {
            sender_id,
            sender_name,
            group_id,
            group_name,
            timestamp,
            avatar_image,
          } = lastJsonMessage?.data;

          const newChat = {
            user_id: group_id ? 0 : sender_id,
            group_id,
            timestamp,
            avatar_image,
            name: group_name ? group_name : sender_name,
          };

          setChatlist((prevChatlist) => [newChat, ...prevChatlist]);
        } else {
          const filteredChatlist = chatlist.filter(
            (chat) => !checkChat([chat?.user_id, chat?.group_id], chatToFind)
          );
          setChatlist([userChat, ...filteredChatlist]);
        }
        break;
      default:
        break;
    }
  }, [lastJsonMessage]);

  const openedChatbox = (
    <Chatbox toggleChat={toggleChat} chat={openChat} user={user} />
  );

  const renderedChats = chatlist.map((chat, index) => (
    <div className="hov" key={index}>
      <li>
        <SingleChatlistItem chat={chat} toggleChat={toggleChat} />
      </li>
      {/* {checkChat([chat?.user_id, chat?.group_id], openChat) && (
        <>
          <Chatbox toggleChat={toggleChat} chat={chat} user={user} />
        </>
      )} */}
    </div>
  ));

  return (
    <>
      <MessageNotification />
      <div className="chat-sidebar">
        <ul className="pepe">{renderedChats}</ul>
      </div>
      {openChat && openedChatbox}
    </>
  );
};

export default Chat;
