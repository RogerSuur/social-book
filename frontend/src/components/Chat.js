import React, { useState, useEffect } from "react";
import axios from "axios";
import { useOutletContext } from "react-router-dom";
import useWebSocket, { ReadyState } from "react-use-websocket";
import Users from "../components/Users";

const LOAD_MSG_HISTORY = "http://localhost:8000/loadmessages";

const Chat = () => {
  const [messageHistory, setMessageHistory] = useState([]);
  const { socketUrl } = useOutletContext();
  const { users, setUsers } = useOutletContext();
  const { sender_id } = useOutletContext();
  const [errMsg, setErrMsg] = useState("");
  const [scroller, setScroller] = useState(0);
  const [msg, setMsg] = useState({
    sender_id: sender_id,
    receiver_id: 0,
    body: "",
    history: 0,
  });

  const loadMessages = async () => {
    try {
      const response = await axios.post(
        LOAD_MSG_HISTORY,
        JSON.stringify(msg),
        {
          withCredentials: true,
        },
        {
          headers: { "Content-Type": "application/json" },
        }
      );

      setMessageHistory((prevMessageHistory) => {
        return [...response.data, ...prevMessageHistory];
      });
    } catch (err) {
      if (!err?.response) {
        setErrMsg("No Server Response");
      } else if (err.response?.status === 401) {
        setErrMsg("Unauthorized");
      } else {
        setErrMsg("Internal Server Error");
      }
    }
  };

  useEffect(() => {
    if (msg.receiver_id !== 0) {
      loadMessages();
    }
  }, [msg.receiver_id]);

  const getChat = (id) => {
    setMsg((prevMsg) => {
      setMessageHistory([]);

      return {
        ...prevMsg,
        receiver_id: id,
      };
    });
  };

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

  const handleScroll = (event) => {
    const { scrollTop, scrollHeight } = event.target;

    if (scrollTop === 0) {
      updateScroll(scrollHeight);
      debounce(loadMessages(), 1000);
      event.target.scrollTop = event.target.scrollHeight - scroller;
    }
  };

  const updateScroll = (num) => {
    setScroller(num);
  };

  const { sendJsonMessage, lastMessage, readyState } = useWebSocket(socketUrl, {
    share: true,
  });

  useEffect(() => {
    setMsg((prevMsg) => {
      return {
        ...prevMsg,
        sender_id: sender_id,
        history: messageHistory.length,
      };
    });
  }, [sender_id, messageHistory, setMessageHistory]);

  const connectionStatus = {
    [ReadyState.CONNECTING]: "Connecting",
    [ReadyState.OPEN]: "Open",
    [ReadyState.CLOSING]: "Closing",
    [ReadyState.CLOSED]: "Closed",
    [ReadyState.UNINSTANTIATED]: "Uninstantiated",
  }[readyState];

  const handleChange = (event) => {
    const { value } = event.target;

    setMsg((prevMsg) => {
      return {
        ...prevMsg,
        body: value,
      };
    });
  };

  const checkStyle = (id) => {
    if (id !== sender_id) {
      return { color: "red" };
    }
  };

  const showUsername = (id) => {
    if (id !== sender_id) {
      return "Them";
    }

    return "You";
  };

  useEffect(() => {
    if (lastMessage !== null) {
      const data = JSON.parse(lastMessage.data);

      if (data.id !== undefined && data.id !== 0) {
        if (
          (msg.receiver_id !== 0 && data.receiver_id === sender_id) ||
          data.receiver_id === msg.receiver_id
        ) {
          setMessageHistory((prevMessageHistory) => [
            ...prevMessageHistory,
            data,
          ]);
        }

        setUsers((prevUsers) =>
          [
            ...prevUsers.map((user) =>
              user.user_id === data.receiver_id ||
              user.user_id === data.sender_id
                ? { ...user, datetime: data.datetime }
                : user
            ),
          ].sort((a, b) =>
            a.datetime < b.datetime
              ? 1
              : b.datetime < a.datetime
              ? -1
              : a.username.toLowerCase() > b.username.toLowerCase()
              ? 1
              : a.username.toLowerCase() < b.username.toLowerCase()
              ? -1
              : 0
          )
        );
      }
    }
  }, [lastMessage, setMessageHistory]);

  const handleSubmit = (event) => {
    event.preventDefault();
    sendJsonMessage(msg);
    setMsg({ ...msg, body: "" });
  };

  return (
    <>
      <span>The WebSocket is currently {connectionStatus}</span>
      <div
        style={{
          display: "flex",
          width: "80%",
          height: "60vh",
          margin: "50px auto",
          border: "1px solid black",
        }}
      >
        <div
          style={{
            width: "100%",
            border: "1px solid black",
            overflowY: "scroll",
            flexBasis: "80%",
            fontSize: "30px",
          }}
          onScroll={handleScroll}
        >
          {messageHistory &&
            messageHistory.map((mess) => (
              <p key={mess.message_id} style={checkStyle(mess.sender_id)}>
                {showUsername(mess.sender_id)}: {mess.body} at{" "}
                {new Date(mess.datetime).toLocaleString("et-EE")}
                <br />
              </p>
            ))}
        </div>
        <Users data={getChat} users={users} sender_id={sender_id} />
      </div>

      <div>
        {msg.receiver_id !== 0 && (
          <div>
            <form onSubmit={handleSubmit}>
              <input
                type="text"
                placeholder="Message"
                onChange={handleChange}
                name="message"
                value={msg.body}
                disabled={readyState !== ReadyState.OPEN}
                autoFocus
                required
              />
              <button>Post</button>
            </form>
          </div>
        )}
      </div>
    </>
  );
};

export default Chat;
