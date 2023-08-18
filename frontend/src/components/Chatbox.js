import React, { useState, useEffect, useCallback, useRef } from "react";
import { WS_URL } from "../utils/routes";
import useWebSocketConnection from "../hooks/useWebSocketConnection";
import { LinkContainer } from "react-router-bootstrap";
import InfiniteScroll from "react-infinite-scroller";
import ImageHandler from "../utils/imageHandler";
import {
  Row,
  CloseButton,
  Form,
  Button,
  Image,
  Col,
  Container,
  Stack,
  Card,
  Alert,
  Badge,
} from "react-bootstrap";
import { Send } from "react-bootstrap-icons";

const Chatbox = ({
  toggleChat,
  chat,
  user,
  updateChatlist,
  resetUnreadCount,
}) => {
  const [messageHistory, setMessageHistory] = useState([]);
  const { sendJsonMessage, lastJsonMessage } = useWebSocketConnection(WS_URL);
  const messageboxRef = useRef(null);
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

  const handleScrolling = () => {
    console.log("SCROLLING");
    console.log("SCROLLTOP: ", messageboxRef?.current?.scrollTop);
    console.log("SCROLLHEIGHT: ", messageboxRef?.current?.scrollHeight);
    console.log("CLIENTHEIGHT: ", messageboxRef?.current?.clientHeight);

    const lastMessage = messageHistory?.[messageHistory.length - 1]?.message_id;
    console.log(
      "MESSAGE HISTORY: ",
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
        sendJsonMessage({
          type: "messages_read",
          data: { message_id: lastMessage },
        });
        resetUnreadCount([chat.user_id, chat.group_id]);
      }
    }
  };

  const image =
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
            <>
              <Col
                md={{ span: 9, offset: 3 }}
                variant="success"
                key={index}
                className="own-message text-end"
              >
                <span className="p-1 bg-success">{msg.body}</span>
              </Col>
              <p className="own-time text-secondary text-muted text-end p-0 m-0">
                {getTime(msg.timestamp)}
              </p>
            </>
          );
        default:
          return (
            <Col md="9" key={index} className="message text-start">
              <span className="p-1 bg-primary">
                {msg.group_id > 0 && msg.sender_name} {msg.body}
              </span>
              <p className="message-time text-secondary text-muted">
                {getTime(msg.timestamp)}
              </p>
            </Col>
          );
      }
    }
  });

  const handleSubmit = (event) => {
    event.preventDefault();
    if (!message?.data?.body) {
      return;
    }
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
      <LinkContainer to={`/profile/${chat.user_id}`}>
        <>{chat.name}</>
      </LinkContainer>
    ) : (
      <LinkContainer to={`/groups/${chat.group_id}`}>
        <>{chat.name}</>
      </LinkContainer>
    );

  const scrollToBottom = () => {
    messageboxRef.current.scrollTop = messageboxRef.current.scrollHeight;
  };

  const chatbox = (
    <Card>
      <Card.Header>
        <Stack direction="horizontal">
          <div className="me-auto">{image}</div>
          <div className="text-center">{chatName}</div>
          <CloseButton
            className="ms-auto align-self-center"
            onClick={closeChat}
          />
        </Stack>
      </Card.Header>
      <Card.Body
        className="message-history h-25"
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
      </Card.Body>
      <Card.Footer>
        <Form onSubmit={handleSubmit}>
          <Stack direction="horizontal" gap={2}>
            <Form.Control
              type="text"
              placeholder="Message"
              onChange={handleChange}
              name="message"
              value={message.data.body}
              autoFocus
            />
            <Button type="submit">
              <Send />
            </Button>
          </Stack>
        </Form>
      </Card.Footer>
    </Card>
  );

  return <div className="chatbox">{chatbox}</div>;
};

export default Chatbox;

{
  /* <div className="chatbox">
        <Container fluid>
          <Row>
            <Col className="m-2">
              <div>{image}</div>
            </Col>
            <Col>
              {chatName}
              <CloseButton onClick={closeChat} />
            </Col>
          </Row>
          <Row
            className="message-history h-25"
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
          </Row>
          <Row className="message-box">
            <Form onSubmit={handleSubmit}>
              <Stack direction="horizontal">
                <Form.Control
                  type="text"
                  placeholder="Message"
                  onChange={handleChange}
                  name="message"
                  value={message.data.body}
                  autoFocus
                  required
                />
                <Button type="submit" variant="flush">
                  <Image src={`${process.env.PUBLIC_URL}/send-icon.png`} />
                </Button>{" "}
              </Stack>
            </Form>
          </Row>
        </Container>
      </div> */
}
