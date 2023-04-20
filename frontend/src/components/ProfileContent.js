import React, { useState, useEffect } from "react";
import UserPosts from "./UserPosts.js";
import Followers from "./Followers.js";
import Following from "./Following.js";
import ProfileInfo from "./ProfileInfo.js";

const ProfileContent = ({ selected }) => {
  console.log(selected, "Profile content");

  const [displayedContent, setDisplayedContent] = useState(null);

  useEffect(() => {
    switch (selected) {
      case "your-posts":
        setDisplayedContent(<UserPosts />);
        break;
      case "followers":
        console.log("FOLLOWERS SELECTED");
        setDisplayedContent(<Followers />);
        break;
      case "following":
        setDisplayedContent(<Following />);
        break;
      default:
        setDisplayedContent(<ProfileInfo />);
        break;
    }
  }, [selected]);

  return <div className="profile-content">{displayedContent}</div>;
};

export default ProfileContent;
