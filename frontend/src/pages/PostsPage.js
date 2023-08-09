import React, { useState } from "react";
import FeedPosts from "../components/FeedPosts";
import CreatePost from "../components/CreatePost";
import GroupSidebar from "../components/GroupSidebar";
import { FEEDPOSTS_URL } from "../utils/routes";
import Col from "react-bootstrap/Col";
import Row from "react-bootstrap/Row";

const PostsPage = () => {
  const [reload, setReload] = useState(false);

  const handlePostUpdate = () => {
    setReload(!reload);
  };

  return (
    <>
      <Row>
        <Col xs="8">
          <CreatePost onPostsUpdate={handlePostUpdate} />

          <FeedPosts url={FEEDPOSTS_URL} reload={reload} />
        </Col>
      </Row>
    </>
  );
};

export default PostsPage;
