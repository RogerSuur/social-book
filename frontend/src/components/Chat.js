import SingleChatlistItem from "./SingleChatlistItem";
import React, { useState, useEffect } from "react";
import useWebSocketConnection from "../hooks/useWebSocketConnection";
import Chatbox from "./Chatbox";
import MessageNotification from "./MessageNotification";
import { WS_URL } from "../utils/routes";

const Chat = () => {
  const [openChat, setOpenChat] = useState(null);
  const [userChatlist, setUserChatlist] = useState([]);
  const [groupChatlist, setGroupChatlist] = useState([]);
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

  const resetUnreadCount = (userid) => {
    setUserChatlist((prevChatlist) =>
      prevChatlist.map((chat) =>
        chat.user_id === userid ? { ...chat, unread_count: 0 } : chat
      )
    );
  };

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
    open.every((value, index) => value === checker[index]);

  const updateChatlist = (chatToFind) => {
    const chatlist = chatToFind?.[0] > 0 ? userChatlist : groupChatlist;

    const userChat = chatlist?.find((chat) =>
      checkChat(
        [
          chat?.user_id ? chat?.user_id : 0,
          chat?.group_id ? chat?.group_id : 0,
        ],
        chatToFind
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
        user_id: group_id > 0 ? 0 : sender_id,
        group_id,
        timestamp,
        avatar_image,
        name: group_name ? group_name : sender_name,
        unread_count: 1,
      };

      newChat?.group_id > 0
        ? setGroupChatlist((prevChatlist) => [newChat, ...prevChatlist])
        : setUserChatlist((prevChatlist) => [newChat, ...prevChatlist]);
    } else {
      const filteredChatlist = chatlist?.filter(
        (chat) =>
          !checkChat(
            [
              chat?.user_id ? chat?.user_id : 0,
              chat?.group_id ? chat?.group_id : 0,
            ],
            chatToFind
          )
      );

      userChat.unread_count += 1;
      chatToFind?.[1] > 0
        ? setGroupChatlist([userChat, ...filteredChatlist])
        : setUserChatlist([userChat, ...filteredChatlist]);
    }
  };

  useEffect(() => {
    switch (lastJsonMessage?.type) {
      case "chatlist":
        setUser(lastJsonMessage?.data?.user_id);
        setUserChatlist([...lastJsonMessage?.data?.user_chatlist]);
        setGroupChatlist([...lastJsonMessage?.data?.group_chatlist]);
        break;
      case "message":
        updateChatlist([
          lastJsonMessage?.data?.group_id > 0
            ? 0
            : lastJsonMessage?.data?.sender_id,
          lastJsonMessage?.data?.group_id,
        ]);
        break;
      default:
        break;
    }
  }, [lastJsonMessage]);

  const openedChatbox = (
    <Chatbox
      toggleChat={toggleChat}
      chat={openChat}
      user={user}
      updateChatlist={updateChatlist}
      resetUnreadCount={resetUnreadCount}
    />
  );

  const renderedChats = (chatlist) =>
    chatlist.map((chat, index) => (
      <div className="hov" key={index}>
        <li>
          <SingleChatlistItem chat={chat} toggleChat={toggleChat} />
          {chat?.user_id > 0 && chat.unread_count > 0 && (
            <span className="chat-unread-count">{chat.unread_count}</span>
          )}
        </li>
      </div>
    ));

  return (
    <>
      <MessageNotification />
      <div className="chat-sidebar">
        <p>Private Chats</p>
        <ul className="pepe">{renderedChats(userChatlist)}</ul>
        <p>Group Chats</p>
        <ul className="pepe">{renderedChats(groupChatlist)}</ul>
      </div>
      {openChat && openedChatbox}
    </>
  );
};

export default Chat;
