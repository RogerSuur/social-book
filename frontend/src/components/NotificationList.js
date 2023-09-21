import React, { useEffect, useRef } from "react";
import Notification from "../components/Notification";
import { Dropdown, ListGroup, Col, Row, Container } from "react-bootstrap";
import Scrollbars from "react-custom-scrollbars-2";

const NotificationList = ({ notifications, setToggle, setNotifications }) => {
  const ref = useRef(null);

  const handleNotificationClose = (id) => {
    console.log(notifications, "NOTIFICATIONS");
    console.log(id, "ID");
    setNotifications((prevNotifications) =>
      prevNotifications.filter((notification) => {
        console.log("NOTIFICATION_ID: ", notification?.notification_id);
        console.log("PASSED ID: ", notification?.notification_id);
        return notification?.notification_id !== id;
      })
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

  const renderedNotifications = notifications.map((notification, index) => (
    <ListGroup.Item variant="success" action key={index}>
      <Notification
        notification={notification}
        onClose={handleNotificationClose}
      />
    </ListGroup.Item>
  ));

  return (
    <>
      {notifications.length > 0 && (
        <ListGroup
          ref={ref}
          className="scroll position-absolute top-0 w-25 start-0"
        >
          {renderedNotifications}
        </ListGroup>
      )}
    </>
  );
};

export default NotificationList;
