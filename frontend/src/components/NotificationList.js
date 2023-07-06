import React, { useEffect, useRef } from "react";
import Notification from "../components/Notification";

const NotificationList = ({ notifications, setToggle, setNotifications }) => {
  const ref = useRef(null);

  const handleNotificationClose = (id) => {
    setNotifications((prevNotifications) =>
      prevNotifications.filter(
        (notification) => notification?.data?.notification_id !== id
      )
    );
  };

  console.log(notifications, "NOTS IN LIST");

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

  const renderedNotifications = notifications.map((notification, index) => (
    <div key={index} className="notification">
      <li>
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
