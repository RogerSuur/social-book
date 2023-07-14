import React, { useState, useEffect } from "react";
import Posts from "../pages/PostsPage.js";
import GenericUserList from "../components/GenericUserList.js";
import ProfileEditor from "../components/ProfileEditor.js";
import { FOLLOWERS_URL, FOLLOWING_URL } from "../utils/routes.js";

const ProfileContent = ({ selected }) => {
  const [displayedContent, setDisplayedContent] = useState(null);

  useEffect(() => {
    switch (selected) {
      case "your-posts":
        setDisplayedContent(null);
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

  return (
    <div className="profile-content">
      {displayedContent}
      {selected === "your-posts" && (
        <Posts
          forGroupPage={false}
          showGroupSidebar={false}
          showCreatePost={false}
          url={"/profileposts"}
        />
      )}
    </div>
  );
};

export default ProfileContent;
