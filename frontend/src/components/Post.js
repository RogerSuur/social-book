import React from "react";
import Comments from "../components/Comments";
import { Container, Row, Col, Image, Stack } from "react-bootstrap";
import { LinkContainer } from "react-router-bootstrap";
import GenericModal from "../components/GenericModal";
import { ShortDate } from "../utils/datetimeConverters";

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
      className="p-3 mt-3 mb-3 border rounded shadow"
      key={id}
      ref={isLastPost ? lastPostElementRef : null}
      data-post-id={id}
    >
      <Stack direction="horizontal">
        <Stack>
          <div>
            <LinkContainer to={`/profile/${userId}`}>
              <span>{userName}</span>
            </LinkContainer>
          </div>
          <div>{ShortDate(createdAt)}</div>
        </Stack>
        {groupId > 0 && (
          <LinkContainer to={`/groups/${groupId}`}>
            <div className="text-end">{groupName}</div>
          </LinkContainer>
        )}
      </Stack>
      <Row className="mb-3 mt-3">
        {imagePath && (
          <Row>
            <Image
              fluid
              className="post-img"
              src={`${process.env.PUBLIC_URL}/images/${imagePath}`}
            />
          </Row>
        )}
        <Row>
          <Col>{content}</Col>
        </Row>
      </Row>

      <Row>
        <Col>
          {commentCount === 0 ? (
            <Comments postId={id} commentCount={commentCount} />
          ) : (
            lastPostElementRef !== undefined && (
              <div className="text-end">
                <GenericModal
                  linkText={`${commentCount} ${
                    commentCount > 1 ? "comments" : "comment"
                  }`}
                  buttonText={`${userName}'s post`}
                >
                  <Post post={post} />
                  <Comments postId={id} commentCount={commentCount} />
                </GenericModal>
              </div>
            )
          )}
        </Col>
      </Row>
    </Container>
  );
};

export default Post;
