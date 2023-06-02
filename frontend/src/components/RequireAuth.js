import { useLocation, Navigate, Outlet } from "react-router-dom";
import { useEffect, useState } from "react";
import useAuth from "../hooks/useAuth";
import axios from "axios";
import { useWebSocketConnection } from "../utils/useWebSocketConnection";

const AUTH_URL = "http://localhost:8000/auth";
const WS_URL = `ws://localhost:8000/ws`;

const RequireAuth = () => {
  const { auth, setAuth } = useAuth();
  const [socketUrl] = useState(WS_URL);
  const [users, setUsers] = useState([]);
  const [sender_id, setSenderId] = useState(0);

  const location = useLocation();

  console.log(location, "LOC AUTH");

  // const updateUsers = (message) => {
  //   const { user_id, username, online, datetime } = message;

  //   setUsers((prevUsers) =>
  //     [
  //       ...prevUsers.filter(
  //         (user) => user.user_id !== user_id && user.user_id !== sender_id
  //       ),
  //       { user_id, username, online, datetime },
  //     ].sort((a, b) =>
  //       a.datetime < b.datetime
  //         ? 1
  //         : b.datetime < a.datetime
  //         ? -1
  //         : a.username.toLowerCase() > b.username.toLowerCase()
  //         ? 1
  //         : a.username.toLowerCase() < b.username.toLowerCase()
  //         ? -1
  //         : 0
  //     )
  //   );
  // };

  useEffect(() => {
    const authorisation = async () => {
      try {
        const response = await axios.get(AUTH_URL, {
          withCredentials: true,
        });

        console.log(JSON.stringify(response));
        setAuth(true);
      } catch (err) {
        if (!err?.response) {
          setAuth(false);
        } else if (err.response?.status === 401) {
          setAuth(false);
        } else {
          setAuth(false);
        }
      }
    };

    authorisation();
  }, [location]);

  // const { lastMessage, sendJsonMessage } = useWebSocket(socketUrl, {
  //   onOpen: console.log("opened"),
  //   share: true,
  // });

  // useEffect(() => {
  //   sendJsonMessage("hello");
  // }, [location]);

  useWebSocketConnection(socketUrl, "msg");

  return auth ? (
    <Outlet
      context={{
        socketUrl,
        users,
        setUsers,
        sender_id,
      }}
    />
  ) : (
    <Navigate to="/login" state={{ from: location }} replace />
  );
};

export default RequireAuth;
