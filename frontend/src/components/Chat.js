import SingleChatlistItem from "./SingleChatlistItem";
import React, { useState, useEffect } from "react";
import useWebSocketConnection from "../hooks/useWebSocketConnection";
import Chatbox from "./Chatbox";
import MessageNotification from "./MessageNotification";
import { WS_URL } from "../utils/routes";

const Chat = ({}) => {
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
        const userChat = chatlist.find((chat) =>
          checkChat(
            [chat?.user_id, chat?.group_id],
            [lastJsonMessage?.data?.sender_id, lastJsonMessage?.data?.group_id]
          )
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
            user_id: sender_id,
            group_id,
            timestamp,
            avatar_image,
            name: sender_name ? sender_name : group_name,
          };

          setChatlist((prevChatlist) => [newChat, ...prevChatlist]);
        } else {
          const filteredChatlist = chatlist.filter(
            (chat) =>
              !checkChat(
                [chat?.user_id, chat?.group_id],
                [
                  lastJsonMessage?.data?.sender_id,
                  lastJsonMessage?.data?.group_id,
                ]
              )
          );
          setChatlist([userChat, ...filteredChatlist]);
        }
        break;
    }
  }, [lastJsonMessage]);

  const renderedChats = chatlist.map((chat, index) => (
    <div className="hov" key={index}>
      <li >
        <SingleChatlistItem chat={chat} toggleChat={toggleChat} />
      </li>
      {checkChat([chat?.user_id, chat?.group_id], openChat) && (
        <>
          <Chatbox toggleChat={toggleChat} chat={chat} user={user} />
        </>
      )}
    </div>
  ));

  return (
    <>
      <MessageNotification />
      <div className="chat-sidebar">
        <ul className="pepe">{renderedChats}</ul>
      </div>
    </>
  );
};

export default Chat;
