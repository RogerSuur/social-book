import React, { useState, useEffect } from "react";
import UserPosts from "../components/UserPosts.js";
import GenericUserList from "../components/GenericUserList.js";
import ProfileEditor from "../components/ProfileEditor.js";
import { FOLLOWERS_URL, FOLLOWING_URL } from "../utils/routes.js";

const ProfileContent = ({ selected }) => {
  const [displayedContent, setDisplayedContent] = useState(null);

  useEffect(() => {
    switch (selected) {
      case "your-posts":
        setDisplayedContent(<UserPosts />);
        break;
      case "followers":
        setDisplayedContent(<GenericUserList url={FOLLOWERS_URL} />);
        break;
      case "following":
        setDisplayedContent(<GenericUserList url={FOLLOWING_URL} />);
        break;
      default:
        setDisplayedContent(<ProfileEditor />);
        break;
    }
  }, [selected]);

  return <div className="profile-content">{displayedContent}</div>;
};

export default ProfileContent;
