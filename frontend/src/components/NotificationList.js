import React, { useState, useEffect } from "react";
import useWebSocketConnection from "../hooks/useWebSocketConnection";
import Notification from "../components/Notification";
import { useOutletContext } from "react-router-dom";
import axios from "axios";

const NOTIFICATIONS_URL = "http://localhost:8000/notifications/";
const WS_URL = `ws://localhost:8000/ws`;

const NotificationList = () => {
  const [notifications, setNotifications] = useState([]);
  const { lastJsonMessage } = useWebSocketConnection(WS_URL);

  useEffect(() => {
    const loadNotifications = async () => {
      await axios
        .get(NOTIFICATIONS_URL, {
          withCredentials: true,
        })
        .then((response) => {
          setNotifications(response.data);
        });
    };
    loadNotifications();
  }, []);

  useEffect(() => {
    if (lastJsonMessage && lastJsonMessage.type !== "new_message") {
      setNotifications((prevNotifications) => {
        return [lastJsonMessage, ...prevNotifications];
      });
    }
  }, [lastJsonMessage]);

  const renderedNotifications = notifications.map((notification, index) => (
    <li key={index}>
      <Notification notification={notification} />
    </li>
  ));

  return (
    <>
      {notifications.length === 0
        ? "You have no notifications"
        : renderedNotifications}
    </>
  );
};

export default NotificationList;
