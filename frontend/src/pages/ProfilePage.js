import React, { useState } from "react";
import ProfileSidebar from "../components/ProfileSidebar.js";
import ProfileContent from "../components/ProfileContent.js";

const Profile = () => {
  const [selected, setSelected] = useState(null);

  const handleSelect = (item) => {
    setSelected(item);
  };

  return (
    <>
      <ProfileSidebar onSelect={handleSelect} />
      <ProfileContent selected={selected} />
    </>
  );
};

export default Profile;
