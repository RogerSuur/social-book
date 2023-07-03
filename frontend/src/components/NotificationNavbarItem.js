import useWebSocketConnection from "../hooks/useWebSocketConnection";
import NotificationList from "../components/NotificationList.js";
import axios from "axios";
import { useState, useEffect } from "react";
import { WS_URL, NOTIFICATIONS_URL } from "../utils/routes";
import { useLocation } from "react-router-dom";

const NotificationNavbarItem = () => {
  const [toggle, setToggle] = useState(false);
  const location = useLocation();
  // const [newNotification, setNewNotification] = useState();
  // const [notificationTimer, setNotificationTimer] = useState(false);
  const { lastJsonMessage } = useWebSocketConnection(WS_URL);
  const [notifications, setNotifications] = useState([]);

  useEffect(() => {
    if (lastJsonMessage && lastJsonMessage.type === "notification") {
      setNotifications((prevNotifications) => {
        return [lastJsonMessage, ...prevNotifications];
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
  }, [location]);

  // useEffect(() => {
  //   const exceptions = ["message", "chatlist", "message_history"];

  //   if (!exceptions.includes(lastJsonMessage?.type)) {
  //     setNewNotification(lastJsonMessage?.data);
  //     setNotificationTimer(true);

  //     const timer = setTimeout(() => {
  //       setNotificationTimer(false);
  //     }, 5000);

  //     return () => clearTimeout(timer);
  //   }
  // }, [lastJsonMessage]);

  const handleToggle = () => {
    setToggle(!toggle);
  };

  const notificationCount = notifications.length;

  return (
    <>
      <li onClick={handleToggle}>
        Notifications
        {notificationCount > 0 && (
          <div className="notification-count">{notificationCount}</div>
        )}
      </li>
      {/* {notificationTimer && (
        <NotificationPopup notification={newNotification} />
      )} */}
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
