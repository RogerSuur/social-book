import { useLocation, Navigate, Outlet } from "react-router-dom";
import { useEffect, useState } from "react";
import useAuth from "../hooks/useAuth";
import axios from "axios";
import useWebSocketConnection from "../hooks/useWebSocketConnection";
import Chat from "./Chat";
import { WS_URL } from "../utils/routes";
import Login from "../pages/LoginPage";
import { Container, Row, Col } from "react-bootstrap";
import GroupSidebar from "../components/GroupSidebar";

const AUTH_URL = "http://localhost:8000/auth";

const RequireAuth = () => {
  const { auth, setAuth } = useAuth();
  const [socketUrl] = useState(WS_URL);
  const { lastJsonMessage } = useWebSocketConnection(socketUrl);
  const [users, setUsers] = useState([]);
  const [loading, setLoading] = useState(true);

  const location = useLocation();

  const from = location.state?.from?.pathname;

  useEffect(() => {
    const authorisation = async () => {
      try {
        await axios.get(AUTH_URL, {
          withCredentials: true,
        });
        console.log("AUTHENTICATION");
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
      setLoading(false);
    };

    authorisation();
  }, [location]);

  useEffect(() => {
    if (lastJsonMessage && lastJsonMessage.type === "chatlist") {
      setUsers(lastJsonMessage.data);
    }
  }, [lastJsonMessage]);

  return loading ? null : auth ? (
    <Container fluid className="bg-light">
      <Row>
        <Col className="sidebar p-0 d-none d-md-flex" id="group-sidebar" xs="3">
          <GroupSidebar />
        </Col>
        <Col
          xs="12"
          md={{ span: "6", offset: "3" }}
          className="mt-3 mb-3 justify-content-end"
        >
          <Outlet
            context={{
              socketUrl,
              users,
              setUsers,
            }}
          />
        </Col>
        <Col id="chat-sidebar" xs="3" className="sidebar p-0 d-none d-md-flex">
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
