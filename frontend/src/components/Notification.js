import { WS_URL } from "../utils/routes";
import useWebSocketConnection from "../hooks/useWebSocketConnection";
import { Link } from "react-router-dom";
import { Button } from "react-bootstrap";
import { ShortDatetime } from "../utils/datetimeConverters";

const Notification = ({ notification, onClose, popup }) => {
  const { sendJsonMessage } = useWebSocketConnection(WS_URL);

  console.log(notification, "NOT BASIC");

  const handleReject = () => {
    const msg = {
      type: "response",
      data: { id: notification?.notification_id, reaction: false },
    };
    console.log(msg, "NOTIFICATION MESSAGE");
    sendJsonMessage(msg);
    onClose(notification?.notification_id);
  };

  const handleAccept = () => {
    const msg = {
      type: "response",
      data: { id: notification?.notification_id, reaction: true },
    };
    console.log(msg, "NOTIFICATION MESSAGE");

    sendJsonMessage(msg);
    onClose(notification?.notification_id);
  };

  const acceptButton = (text = "Accept") => (
    <Button onClick={handleAccept}>{text}</Button>
  );

  const rejectButton = (text = "Reject") => (
    <Button onClick={handleReject}>{text}</Button>
  );

  const followRequestNotification = () => {
    return (
      <>
        <Link to={`/profile/${notification?.sender_id}`}>
          {notification?.sender_name}
        </Link>{" "}
        wants to follow you
        {!popup && (
          <>
            {acceptButton()}
            {rejectButton()}
          </>
        )}
      </>
    );
  };

  const groupInviteNotification = () => {
    return (
      <>
        <Link to={`/profile/${notification?.sender_id}`}>
          {notification?.sender_name}
        </Link>{" "}
        invites you to join the group{" "}
        <Link to={`/groups/${notification?.group_id}`}>
          {notification?.group_name}
        </Link>
        {!popup && (
          <>
            {acceptButton()}
            {rejectButton()}
          </>
        )}
      </>
    );
  };

  const groupRequestNotification = () => {
    return (
      <>
        <Link to={`/profile/${notification?.sender_id}`}>
          {notification?.sender_name}
        </Link>{" "}
        wants to join your group{" "}
        <Link to={`/groups/${notification?.group_id}`}>
          {notification?.group_name}
        </Link>
        {!popup && (
          <>
            {acceptButton()}
            {rejectButton()}
          </>
        )}
      </>
    );
  };

  const eventNotification = () => {
    return (
      <>
        <Link to={`/event/${notification?.event_id}`}>
          {notification?.event_name}
        </Link>{" "}
        is going to take place at {ShortDatetime(notification?.event_datetime)}
        {!popup && (
          <>
            {acceptButton("Going")}
            {rejectButton("Not going")}
          </>
        )}
      </>
    );
  };

  const notificationMessage = () => {
    switch (notification?.notification_type) {
      case "follow_request":
        return followRequestNotification();
      case "group_invite":
        return groupInviteNotification();
      case "group_request":
        return groupRequestNotification();
      case "event_invite":
        return eventNotification();
      default:
        break;
    }
  };

  return notificationMessage();
};

export default Notification;
