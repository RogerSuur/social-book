import React from "react";
import Notification from "../components/Notification";

const NotificationPopup = ({ notification, onPopupClose }) => {
  return (
    <>
      <Notification notification={notification} popup={true} />
      <button onClick={onPopupClose}>X</button>
    </>
  );
};

export default NotificationPopup;
