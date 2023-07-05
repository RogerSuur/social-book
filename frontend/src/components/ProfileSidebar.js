import React, { useState } from "react";

const ProfileSidebar = ({ onSelect }) => {
  const [selected, setSelected] = useState([]);

  const sidebarItems = ["General Info", "Your Posts", "Following", "Followers"];

  const handleClick = (i) => {
    const item = i.replace(" ", "-").toLowerCase();
    setSelected(item);
    onSelect(item);
  };

  const listSidebarItems = sidebarItems.map((item, index) => (
    <li className="bigger" key={index} onClick={() => handleClick(item)}>
      {item}
    </li>
  ));

  return <div className="profile-side">{listSidebarItems}</div>;
};

export default ProfileSidebar;
