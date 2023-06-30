import React, { useState, useEffect, useRef } from "react";
import useWebSocketConnection from "../hooks/useWebSocketConnection";
import Notification from "../components/Notification";
import { WS_URL } from "../utils/routes";

const NotificationList = ({ notifications, setToggle, setNotifications }) => {
  const ref = useRef(null);
  const { lastJsonMessage } = useWebSocketConnection(WS_URL);

  const handleNotificationClose = (id) => {
    setNotifications((prevNotifications) =>
      prevNotifications.filter((notification) => notification?.data?.id !== id)
    );
  };

  useEffect(() => {
    const handleClickOutside = (event) => {
      if (ref.current && !ref.current.contains(event.target)) {
        setToggle(false);
      }
    };

    document.addEventListener("mousedown", handleClickOutside);

    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, [ref]);

  useEffect(() => {
    if (lastJsonMessage && lastJsonMessage.type !== "message") {
      setNotifications((prevNotifications) => {
        return [lastJsonMessage, ...prevNotifications];
      });
    }
  }, [lastJsonMessage]);

  const renderedNotifications = notifications.map((notification) => (
    <div className="notification">
      <li key={notification?.data?.id}>
        <Notification
          notification={notification}
          onClose={handleNotificationClose}
        />
      </li>
    </div>
  ));

  return (
    <div className="notification-list" ref={ref}>
      {notifications.length === 0
        ? "You have no notifications"
        : renderedNotifications}
    </div>
  );
};

export default NotificationList;

// const notif = [
//   {
//     type: "follow_request",
//     data: { id: 1, following_id: 3, username: "Jim Boles" },
//   },
//   {
//     type: "group_invite",
//     data: {
//       id: 2,
//       sender_id: 1,
//       username: "Jo-Jo",
//       group_id: 1,
//       group_name: "Funky Animals",
//     },
//   },
//   {
//     type: "group_join",
//     data: {
//       id: 3,
//       sender_id: 3,
//       username: "Kevin Bacon",
//       group_id: 2,
//       group_name: "Bad Weather",
//     },
//   },
//   {
//     type: "event",
//     data: {
//       id: 4,
//       event_id: 1,
//       event_name: "Fyre Party",
//       event_datetime: "2023-06-05 16:01:00.303095707+03:00",
//     },
//   },
// ];
