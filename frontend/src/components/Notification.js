const Notification = ({ notification }) => {
  let notificationMessage;

  switch (notification.type) {
    case "follow_request":
      notificationMessage = `${notification.data.username} wants to follow you`;
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
      notificationMessage = `${notification.data.username} invited you to the group ${notification.data.group_name}`;
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

  return <p>{notificationMessage}</p>;
};

export default Notification;
