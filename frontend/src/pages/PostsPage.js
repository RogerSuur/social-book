import React, { useState } from "react";
import FeedPosts from "../components/FeedPosts";
import CreatePost from "../components/CreatePost";
import GroupSidebar from "../components/GroupSidebar";
import { FEEDPOSTS_URL } from "../utils/routes";

const PostsPage = () => {
  const [reload, setReload] = useState(false);

  const handlePostUpdate = () => {
    setReload(!reload);
  };

  return (
    <>
      <GroupSidebar />
      <div className="content-as">
        <CreatePost onPostsUpdate={handlePostUpdate} />

        <FeedPosts url={FEEDPOSTS_URL} reload={reload} />
      </div>
    </>
  );
};

export default PostsPage;
