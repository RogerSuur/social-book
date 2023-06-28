import React, { useState, useEffect } from "react";
import Posts from "../pages/PostsPage.js";
import Followers from "./Followers.js";
import Following from "./Following.js";
import ProfileEditor from "./ProfileEditor.js";

const ProfileContent = ({ selected }) => {
  const [displayedContent, setDisplayedContent] = useState(null);

  useEffect(() => {
    switch (selected) {
      case "your-posts":
        setDisplayedContent(null);
        break;
      case "followers":
        console.log("FOLLOWERS SELECTED");
        setDisplayedContent(<Followers />);
        break;
      case "following":
        setDisplayedContent(<Following />);
        break;
      default:
        setDisplayedContent(<ProfileEditor />);
        break;
    }
  }, [selected]);

  return (
    <div className="profile-content">
      {displayedContent}
      {selected === "your-posts" && <Posts />}
    </div>
  );
};

export default ProfileContent;
