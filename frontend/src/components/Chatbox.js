import { WS_URL } from "../utils/routes";
import useWebSocketConnection from "../hooks/useWebSocketConnection";
import React, { useState, useEffect, useCallback, useRef } from "react";
import { Link } from "react-router-dom";
import InfiniteScroll from "react-infinite-scroller";
import ImageHandler from "../utils/imageHandler";

const Chatbox = ({
  toggleChat,
  chat,
  user,
  updateChatlist,
  resetUnreadCount,
}) => {
  const [messageHistory, setMessageHistory] = useState([]);
  const { sendJsonMessage, lastJsonMessage } = useWebSocketConnection(WS_URL);
  const [hasMoreMessages, setHasMoreMessages] = useState(true);
  const [lastMessageRead, setLastMessageRead] = useState(0);
  const [scrollToBottomNeeded, setScrollToBottomNeeded] = useState(false);
  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState({
    type: "message",
    data: {
      body: "",
    },
  });

  console.log(messageHistory);
  useEffect(() => {
    setMessageHistory([]);
    setHasMoreMessages(true);
  }, [chat]);

  const messageboxRef = useRef(null);

  const debounce = (f, ms) => {
    let timeout;
    return function executedFunction() {
      const context = this;
      const args = arguments;
      const later = function () {
        timeout = null;
        f.apply(context, args);
      };
      clearTimeout(timeout);
      timeout = setTimeout(later, ms);
    };
  };

  const sendMessageRead = () => {
    console.log("SCROLLING");
    console.log("SCROLLTOP: ", messageboxRef?.current?.scrollTop);
    console.log("SCROLLHEIGHT: ", messageboxRef?.current?.scrollHeight);
    console.log("CLIENTHEIGHT: ", messageboxRef?.current?.clientHeight);

    const lastMessage = messageHistory?.[messageHistory.length - 1]?.message_id;
    console.log(
      "MESSAGE HUISTORY: ",
      messageHistory?.[messageHistory.length - 1]
    );

    console.log("LAST MESSAGE: ", lastMessage);
    console.log("lastReadMessage: ", lastMessageRead);

    if (lastMessage && lastMessage !== lastMessageRead) {
      setLastMessageRead(lastMessage);
      if (
        messageboxRef?.current?.scrollHeight -
          messageboxRef?.current?.clientHeight <=
        messageboxRef?.current?.scrollTop + 1
      ) {
        console.log("SENDing");
        sendJsonMessage({
          type: "messages_read",
          data: { message_id: lastMessage },
        });
        resetUnreadCount([chat.user_id, chat.group_id]);
      }
    }
  };

  const handleScrolling = () => {
    debounce(sendMessageRead(), 200);
  };

  const image = () =>
    chat?.user_id > 0
      ? ImageHandler(chat?.avatar_image, "defaultuser.jpg", "chatbox-img")
      : ImageHandler("", "defaultgroup.png", "chatbox-img");

  const loadMessages = useCallback(async () => {
    if (loading) {
      return;
    }

    setLoading(true);

    const offset = messageHistory.length > 0 ? messageHistory[0].id : 0;

    sendJsonMessage({
      type: "request_message_history",
      data: { id: chat.user_id, group_id: chat.group_id, last_message: offset },
    });
  }, [loading, hasMoreMessages]);

  useEffect(() => {
    switch (lastJsonMessage?.type) {
      case "message_history":
        if (lastJsonMessage?.data.length > 0) {
          setMessageHistory((prevMessageHistory) => [
            ...lastJsonMessage?.data,
            ...prevMessageHistory,
          ]);
        }

        if (lastJsonMessage?.data.length < 10) {
          setHasMoreMessages(false);
        }

        setLoading(false);
        break;
      case "message":
        if (
          (lastJsonMessage?.data?.sender_id === chat.user_id &&
            lastJsonMessage?.data?.group_id === 0) ||
          lastJsonMessage?.data?.recipient_id === chat.user_id ||
          lastJsonMessage?.data?.group_id === chat.group_id
        ) {
          setMessageHistory((prevMessageHistory) => [
            ...prevMessageHistory,
            lastJsonMessage?.data,
          ]);
        }
        break;
      default:
        break;
    }
  }, [lastJsonMessage]);

  const closeChat = () => {
    toggleChat(null);
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

  const getTime = (datetime) =>
    new Date(datetime).toLocaleTimeString([], {
      hour: "2-digit",
      minute: "2-digit",
      hour12: false,
    });

  const renderedMessages = messageHistory.map((msg, index) => {
    if (msg) {
      switch (msg.sender_id) {
        case user:
          return (
            <p key={index} className="own-message">
              {msg.body}
              <span className="own-time"> {getTime(msg.timestamp)}</span>
            </p>
          );
        default:
          return (
            <p key={index} className="message">
              {msg.group_id > 0 && msg.sender_name} {msg.body}
              <span className="message-time">{getTime(msg.timestamp)}</span>
            </p>
          );
      }
    }
  });

  const handleSubmit = (event) => {
    event.preventDefault();
    let msg = {
      ...message,
      data: {
        ...message.data,
        sender_id: user,
        recipient_id: chat?.user_id,
        group_id: chat?.group_id,
      },
    };

    sendJsonMessage(msg);

    setMessageHistory((prevMessageHistory) => [
      ...prevMessageHistory,
      { ...msg.data, timestamp: new Date().toISOString() },
    ]);

    updateChatlist([
      chat?.user_id ? chat?.user_id : 0,
      chat?.group_id ? chat?.group_id : 0,
    ]);

    setMessage({ ...message, data: { body: "" } });
    setScrollToBottomNeeded(true);
    resetUnreadCount([chat.user_id, chat.group_id]);
  };

  useEffect(() => {
    if (scrollToBottomNeeded) {
      scrollToBottom();
      setScrollToBottomNeeded(false);
    }
  }, [scrollToBottomNeeded]);

  const chatName =
    chat?.user_id > 0 ? (
      <Link to={`/profile/${chat.user_id}`}>{chat.name}</Link>
    ) : (
      <Link to={`/groups/${chat.group_id}`}>{chat.name}</Link>
    );

  // const handleEmojiClick = (emojiData) => {
  //   const { emoji } = emojiData;

  //   setMessage((prevMessage) => ({
  //     ...prevMessage,
  //     data: { ...prevMessage.data, body: prevMessage.data.body + emoji },
  //   }));
  // };

  const scrollToBottom = () => {
    messageboxRef.current.scrollTop = messageboxRef.current.scrollHeight;
  };

  const chatbox = (
    <div className="chatbox">
      <div className="chat-title">
        {image()}
        {chatName}
        <img
          className="exit-but"
          src={`${process.env.PUBLIC_URL}/Vectorexit.png`}
          onClick={closeChat}
        />
      </div>
      <div
        className="message-history"
        ref={messageboxRef}
        onScroll={handleScrolling}
      >
        <InfiniteScroll
          pageStart={0}
          isReverse={true}
          loadMore={loadMessages}
          hasMore={hasMoreMessages}
          useWindow={false}
        >
          {renderedMessages}
        </InfiniteScroll>
      </div>
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
          <button>
            <img
              className="send-icon"
              src={`${process.env.PUBLIC_URL}/send-icon.png`}
            />
          </button>
        </form>
      </div>
    </div>
  );

  return <>{chatbox}</>;
};

export default Chatbox;
