import SingleChatlistItem from "./SingleChatlistItem";
import React, { useState, useEffect } from "react";
import useWebSocketConnection from "../hooks/useWebSocketConnection";
import Chatbox from "./Chatbox";
import { WS_URL } from "../utils/routes";
import { Container, ListGroup } from "react-bootstrap";
import Scrollbars from "react-custom-scrollbars-2";
import { EnvelopeFill } from "react-bootstrap-icons";

const Chat = ({ newMessages, setNewMessages }) => {
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

  const resetUnreadCount = (openChatbox) => {
    openChatbox[0]
      ? setUserChatlist((prevChatlist) =>
          prevChatlist.map((chat) =>
            checkChat([chat.user_id, chat.group_id], openChatbox)
              ? { ...chat, unread_count: 0 }
              : chat
          )
        )
      : openChatbox[1] &&
        setGroupChatlist((prevChatlist) =>
          prevChatlist.map((chat) =>
            checkChat([chat.user_id, chat.group_id], openChatbox)
              ? { ...chat, unread_count: 0 }
              : chat
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
      console.log("OPENING");
      setOpenChat(chat);
    }
  };

  const checkChat = (open, checker) => {
    console.log("OPEN", open, checker);
    return open.every((value, index) => value === checker[index]);
  };
  console.log("OPENCHAT", openChat);

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

    console.log("USERCHAT", userChat);

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
      console.log("USERCHAT BEFORE: ", userChat);

      userChat.unread_count += 1;
      console.log("USERCHAT: ", userChat);
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
        setNewMessages && setNewMessages(true);
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

  const navbarNotification = (unread) =>
    setNewMessages && unread > 0 && !newMessages && setNewMessages(true);

  const renderedChats = (chatlist) =>
    chatlist.map((chat, index) => {
      navbarNotification(chat?.unread_count);
      console.log("Group id", chat?.group_id, chat?.unread_count);
      console.log("Group id", chat?.group_id);
      return (
        <ListGroup.Item key={index} action onClick={() => toggleChat(chat)}>
          <SingleChatlistItem chat={chat} />
          {chat.unread_count > 0 && (
            <>
              {chat.unread_count}
              <EnvelopeFill color="red" />
            </>
          )}
        </ListGroup.Item>
      );
    });

  return (
    <Scrollbars>
      <Container fluid>
        {userChatlist?.length > 0 && (
          <ListGroup variant="flush">
            <h4>Private Chats</h4>
            {renderedChats(userChatlist)}
          </ListGroup>
        )}
        {groupChatlist?.length > 0 && (
          <ListGroup variant="flush">
            <h4>Group Chats</h4>
            {renderedChats(groupChatlist)}
          </ListGroup>
        )}
      </Container>
      {openChat && openedChatbox}
    </Scrollbars>
  );
};

export default Chat;
