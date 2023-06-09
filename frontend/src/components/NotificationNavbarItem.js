import useWebSocketConnection from "../hooks/useWebSocketConnection";
import NotificationList from "../components/NotificationList.js";
import { useState, useEffect } from "react";
import { WS_URL } from "../utils/routes";

const NotificationNavbarItem = () => {
  const [toggle, setToggle] = useState(false);
  const { sendJsonMessage, lastJsonMessage } = useWebSocketConnection(WS_URL);
  const [notifications, setNotifications] = useState([]);

  const loadNotifications = () => {
    sendJsonMessage({
      type: "notifications",
    });
  };

  useEffect(() => {
    loadNotifications();
  }, []);

  useEffect(() => {
    if (lastJsonMessage?.type === "notifications") {
      setNotifications(lastJsonMessage?.data?.notifications);
    } else if (
      lastJsonMessage?.type !== "message" &&
      lastJsonMessage?.type !== "chatlist"
    ) {
    }
  }, [lastJsonMessage]);

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
      {toggle && <NotificationList setToggle={setToggle} />}
    </>
  );
};

export default NotificationNavbarItem;
