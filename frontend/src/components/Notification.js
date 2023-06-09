import { WS_URL } from "../utils/routes";
import useWebSocketConnection from "../hooks/useWebSocketConnection";
import { Link } from "react-router-dom";

const Notification = ({ notification, onClose }) => {
  const { sendJsonMessage } = useWebSocketConnection(WS_URL);

  // const handleResponse = (e) => {
  //   let msg;

  //   switch (e) {
  //     case true:
  //       msg = { data: { id: notification?.data?.id, reaction: true } };
  //       break;
  //     case false:
  //       msg = {
  //         data: { id: notification?.data?.id, reaction: false },
  //       };
  //       break;
  //     default:
  //       break;
  //   }

  //   console.log({ type: "response", ...msg });
  //   sendJsonMessage({ type: "response", ...msg });
  // };

  const handleReject = () => {
    const msg = {
      type: "response",
      data: { id: notification?.data?.id, reaction: false },
    };
    console.log(msg);
    sendJsonMessage(msg);
    onClose(notification?.data?.id);
  };

  const handleAccept = () => {
    const msg = {
      type: "response",
      data: { id: notification?.data?.id, reaction: true },
    };
    console.log(msg);
    sendJsonMessage(msg);
    onClose(notification?.data?.id);
  };

  // const handleFollowRequest = (e) => {
  //   let msg;

  //   switch (e) {
  //     case true:
  //       msg = { type: "follow_accept", data: { id: notification?.data?.id, reaction: true } };
  //       break;
  //     case false:
  //       msg = { type: "follow_reject", data: { id: notification?.data?.id, reaction: false } };
  //       break;
  //     default:
  //       break;
  //   }

  //   console.log(msg);
  //   sendJsonMessage(msg);
  // };

  // const handleGroupInvite = (e) => {
  //   let msg;

  //   switch (e) {
  //     case true:
  //       msg = { type: "group_invite_accept", data: { id: notification?.data?.id, reaction: true } };
  //       break;
  //     case false:
  //       msg = { type: "group_invite_reject", data: { id: notification?.data?.id, reaction: false } };
  //       break;
  //     default:
  //       break;
  //   }

  //   console.log(msg);
  //   sendJsonMessage(msg);
  // };

  const acceptButton = <button onClick={handleAccept}>Accept</button>;

  const rejectButton = <button onClick={handleReject}>Reject</button>;

  const followRequestNotification = () => {
    return (
      <>
        <Link to={`/profile/${notification?.data?.following_id}`}>
          {notification?.data?.username}
        </Link>{" "}
        wants to follow you
        {acceptButton}
        {rejectButton}
      </>
    );
  };

  const groupInviteNotification = () => {
    return (
      <>
        <Link to={`/profile/${notification?.data?.sender_id}`}>
          {notification?.data?.username}
        </Link>{" "}
        invites you to join the group{" "}
        <Link to={`/groups/${notification?.data?.group_id}`}>
          {notification?.data?.group_name}
        </Link>
        {acceptButton}
        {rejectButton}
      </>
    );
  };

  const groupJoinNotification = () => {
    return (
      <>
        <Link to={`/profile/${notification?.data?.sender_id}`}>
          {notification?.data?.username}
        </Link>{" "}
        wants to join your group{" "}
        <Link to={`/groups/${notification?.data?.group_id}`}>
          {notification?.data?.group_name}
        </Link>
        {acceptButton}
        {rejectButton}
      </>
    );
  };

  let notificationMessage;

  switch (notification.type) {
    case "follow_request":
      notificationMessage = followRequestNotification();
      break;
    case "follow_accept":
      notificationMessage = `${notification.data.username} accepted your follow request`;
      break;
    case "follow_reject":
      notificationMessage = `${notification.data.username} rejected your follow request`;
      break;
    case "unfollow":
      notificationMessage = `${notification.data.username} unfollowed you`;
      break;
    case "group_invite":
      notificationMessage = groupInviteNotification();
      break;
    case "group_join":
      notificationMessage = groupJoinNotification();
      break;
    case "group_accept":
      console.log(notification, "group_accept");
      break;
    case "group_reject":
      console.log(notification, "group_reject");
      break;
    default:
      break;
  }

  return <div className="notification">{notificationMessage}</div>;
};

export default Notification;
