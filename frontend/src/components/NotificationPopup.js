import React from "react";
import Notification from "../components/Notification";

const NotificationPopup = ({ notification }) => {
  return <Notification notification={notification} popup={true} />;
};

export default NotificationPopup;
