import { useLocation, Navigate, Outlet } from "react-router-dom";
import { useEffect, useState } from "react";
import useAuth from "../hooks/useAuth";
import axios from "axios";
import useWebSocketConnection from "../hooks/useWebSocketConnection";
import Chat from "./Chat";
import { WS_URL } from "../utils/routes";
import Login from "../pages/LoginPage";
import Container from "react-bootstrap/Container";
import GroupSidebar from "../components/GroupSidebar";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";

const AUTH_URL = "http://localhost:8000/auth";

const RequireAuth = () => {
  const { auth, setAuth } = useAuth();
  const [socketUrl] = useState(WS_URL);
  const { lastJsonMessage } = useWebSocketConnection(socketUrl);
  const [users, setUsers] = useState([]);

  const location = useLocation();

  const from = location.state?.from?.pathname;

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
        await axios.get(AUTH_URL, {
          withCredentials: true,
        });

        // console.log(JSON.stringify(response), "RESPONSE!!!!!!!!!!!!!!!!!!!");
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

  useEffect(() => {
    if (lastJsonMessage && lastJsonMessage.type === "chatlist") {
      setUsers(lastJsonMessage.data);
    }
  }, [lastJsonMessage]);

  return auth ? (
    <Container fluid>
      <Row>
        <Col className="sidebar d-none d-sm-flex" id="group-sidebar" xs="3">
          <GroupSidebar />
        </Col>
        <Col className="p-3" xs="12" sm={{ span: "7", offset: "3" }}>
          <Outlet
            context={{
              socketUrl,
              users,
              setUsers,
            }}
          />
        </Col>
        <Col id="chat-sidebar" xs="2" className="sidebar p-0 d-none d-sm-flex">
          <Chat chatlist={users} />
        </Col>
      </Row>
    </Container>
  ) : (
    <>
      {from && <Navigate to="/login" state={{ from: location }} replace />}
      <Login />
    </>
  );
};

export default RequireAuth;
