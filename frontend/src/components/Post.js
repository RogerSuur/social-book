import React, { useState } from "react";
import Comments from "../components/Comments";
import { Container, Row, Col, Image } from "react-bootstrap";
import { LinkContainer } from "react-router-bootstrap";

const Post = ({ post, isLastPost, lastPostElementRef }) => {
  const {
    id,
    userId,
    imagePath,
    userName,
    content,
    createdAt,
    commentCount,
    groupId,
    groupName,
  } = post;

  return (
    <Container
      fluid
      className="mt-3 mb-5"
      key={id}
      ref={isLastPost ? lastPostElementRef : null}
      data-post-id={id}
    >
      {groupId > 0 && (
        <LinkContainer className="float-end" to={`/groups/${groupId}`}>
          <>{groupName}</>
        </LinkContainer>
      )}
      <Row>
        {imagePath && (
          <Image
            fluid
            className="post-img"
            src={`${process.env.PUBLIC_URL}/images/${imagePath}`}
          />
        )}
      </Row>
      <Row>
        <Col>{content}</Col>
      </Row>
      <Row>
        <Col xs="4">{new Date(createdAt).toLocaleString("et-EE")}</Col>
        <Col className="text-end">
          <LinkContainer to={`/profile/${userId}`}>
            <span>{userName}</span>
          </LinkContainer>
        </Col>{" "}
      </Row>
      <hr />

      <Row>
        <Comments postId={id} commentCount={commentCount} />
      </Row>
    </Container>
  );
};

export default Post;
