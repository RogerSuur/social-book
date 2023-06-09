import React, { useState, useEffect } from "react";
import useWebSocketConnection from "../hooks/useWebSocketConnection";
import Notification from "../components/Notification";
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

  const handleNotificationClose = (id) => {
    setNotifications((prevNotifications) =>
      prevNotifications.filter((notification) => notification?.data?.id !== id)
    );
  };

  const notif = [
    {
      type: "follow_request",
      data: { id: 1, following_id: 3, username: "Jim Boles" },
    },
    {
      type: "group_invite",
      data: {
        id: 2,
        sender_id: 1,
        username: "Jo-Jo",
        group_id: 1,
        group_name: "Funky Animals",
      },
    },
    {
      type: "group_join",
      data: {
        id: 3,
        sender_id: 3,
        username: "Kevin Bacon",
        group_id: 2,
        group_name: "Bad Weather",
      },
    },
  ];

  useEffect(() => {
    if (lastJsonMessage && lastJsonMessage.type !== "message") {
      setNotifications((prevNotifications) => {
        return [lastJsonMessage, ...prevNotifications];
      });
    }
  }, [lastJsonMessage]);

  const renderedNotifications = notif.map((notification) => (
    <li key={notification?.data?.id}>
      <Notification
        notification={notification}
        onClose={handleNotificationClose}
      />
    </li>
  ));

  return (
    <>
      {notif.length === 0 ? "You have no notifications" : renderedNotifications}
    </>
  );
};

export default NotificationList;
