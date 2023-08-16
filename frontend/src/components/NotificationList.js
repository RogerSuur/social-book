import React, { useEffect, useRef } from "react";
import Notification from "../components/Notification";
import { Dropdown, ListGroup, Col, Row, Container } from "react-bootstrap";

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
    <Container>
      <ListGroup.Item variant="success" action key={index}>
        <Notification
          notification={notification}
          onClose={handleNotificationClose}
        />
      </ListGroup.Item>
    </Container>
  ));

  return (
    <>
      {notifications.length > 0 && (
        <ListGroup ref={ref} className="position-absolute top-100 w-25 start-0">
          <Col>{renderedNotifications}</Col>
        </ListGroup>
      )}
    </>
  );
};

export default NotificationList;
