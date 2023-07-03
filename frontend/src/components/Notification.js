import { WS_URL } from "../utils/routes";
import useWebSocketConnection from "../hooks/useWebSocketConnection";
import { Link } from "react-router-dom";

const Notification = ({ notification, onClose }) => {
  const { sendJsonMessage } = useWebSocketConnection(WS_URL);

  const handleReject = () => {
    const msg = {
      type: "response",
      data: { id: notification?.data?.notification_id, reaction: false },
    };
    sendJsonMessage(msg);
    onClose(notification?.data?.notification_id);
  };

  const handleAccept = () => {
    const msg = {
      type: "response",
      data: { id: notification?.data?.notification_id, reaction: true },
    };
    sendJsonMessage(msg);
    onClose(notification?.data?.notification_id);
  };

  const acceptButton = (text = "Accept") => (
    <button onClick={handleAccept}>{text}</button>
  );

  const rejectButton = (text = "Reject") => (
    <button onClick={handleReject}>{text}</button>
  );

  const followRequestNotification = () => {
    return (
      <>
        <Link to={`/profile/${notification?.data?.sender_id}`}>
          {notification?.data?.sender_name}
        </Link>{" "}
        wants to follow you
        {acceptButton()}
        {rejectButton()}
      </>
    );
  };

  const groupInviteNotification = () => {
    return (
      <>
        <Link to={`/profile/${notification?.data?.sender_id}`}>
          {notification?.data?.sender_name}
        </Link>{" "}
        invites you to join the group{" "}
        <Link to={`/groups/${notification?.data?.group_id}`}>
          {notification?.data?.group_name}
        </Link>
        {acceptButton()}
        {rejectButton()}
      </>
    );
  };

  const groupJoinNotification = () => {
    return (
      <>
        <Link to={`/profile/${notification?.data?.sender_id}`}>
          {notification?.data?.sender_name}
        </Link>{" "}
        wants to join your group{" "}
        <Link to={`/groups/${notification?.data?.group_id}`}>
          {notification?.data?.group_name}
        </Link>
        {acceptButton()}
        {rejectButton()}
      </>
    );
  };

  const eventNotification = () => {
    return (
      <>
        <Link to={`/events/${notification?.data?.event_id}`}>
          {notification?.data?.event_name}
        </Link>{" "}
        is going to take place at{" "}
        {new Date(notification?.data?.event_datetime).toLocaleString("et-EE")}
        {acceptButton("Going")}
        {rejectButton("Not going")}
      </>
    );
  };

  const notificationMessage = () => {
    switch (notification?.data?.notification_type) {
      case "follow_request":
        return followRequestNotification();
      case "group_invite":
        return groupInviteNotification();
      case "group_join":
        return groupJoinNotification();
      case "event":
        return eventNotification();
      default:
        break;
    }
  };

  return <div>{notificationMessage()}</div>;
};

export default Notification;
