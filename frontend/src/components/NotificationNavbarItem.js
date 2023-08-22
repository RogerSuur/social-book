import useWebSocketConnection from "../hooks/useWebSocketConnection";
import NotificationList from "../components/NotificationList.js";
import axios from "axios";
import React, { useState, useEffect } from "react";
import { WS_URL, NOTIFICATIONS_URL } from "../utils/routes";
import NotificationPopup from "../components/NotificationPopup";
import { Badge, Row, Col, Image } from "react-bootstrap";

const NotificationNavbarItem = () => {
  const [toggle, setToggle] = useState(false);
  const [newNotification, setNewNotification] = useState(null);
  const { lastJsonMessage } = useWebSocketConnection(WS_URL);
  const [notifications, setNotifications] = useState([]);

  useEffect(() => {
    if (lastJsonMessage && lastJsonMessage.type === "notification") {
      setNotifications((prevNotifications) => {
        return [lastJsonMessage?.data, ...prevNotifications];
      });
    }
  }, [lastJsonMessage]);

  useEffect(() => {
    const loadNotifications = async () => {
      try {
        await axios
          .get(NOTIFICATIONS_URL, {
            withCredentials: true,
          })
          .then((response) => {
            setNotifications(response.data);
          });
      } catch (err) {
        if (err.response?.status > 200) {
          console.log(err);
        }
      }
    };

    loadNotifications();
  }, []);

  console.log(notifications, "NOTLIST");
  useEffect(() => {
    const exceptions = ["message", "chatlist", "message_history"];

    if (!exceptions.includes(lastJsonMessage?.type)) {
      setNewNotification(lastJsonMessage?.data);
    }
  }, [lastJsonMessage]);

  const handleToggle = () => {
    setToggle(!toggle);
  };

  const onPopupClose = () => {
    setNewNotification(null);
  };

  const notificationCount = notifications.length;

  return (
    <>
      <Row>
        <Col>
          <Image
            src={`${process.env.PUBLIC_URL}/notification_bell.png`}
            onClick={handleToggle}
          />
          {notificationCount > 0 && (
            <span className="position-absolute">
              <Badge pill bg="danger">
                {notificationCount}
              </Badge>
            </span>
          )}
        </Col>
      </Row>

      {newNotification && (
        <NotificationPopup
          notification={newNotification}
          onPopupClose={onPopupClose}
        />
      )}
      {toggle && (
        <NotificationList
          notifications={notifications}
          setNotifications={setNotifications}
          setToggle={setToggle}
        />
      )}
    </>
  );
};

export default NotificationNavbarItem;
