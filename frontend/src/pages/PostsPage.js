import React, { useState } from "react";
import FeedPosts from "../components/FeedPosts";
import CreatePost from "../components/CreatePost";
import { FEEDPOSTS_URL } from "../utils/routes";
import Col from "react-bootstrap/Col";
import Row from "react-bootstrap/Row";
import Container from "react-bootstrap/esm/Container";

const PostsPage = () => {
  const [reload, setReload] = useState(false);

  const handlePostUpdate = () => {
    setReload(!reload);
  };

  return (
    <Container fluid>
      <Row>
        <Col>
          <CreatePost onPostsUpdate={handlePostUpdate} />
        </Col>
      </Row>
      <Row>
        <Col>
          <FeedPosts url={FEEDPOSTS_URL} reload={reload} />
        </Col>
      </Row>
    </Container>
  );
};

export default PostsPage;
