import React from "react";
import Comments from "../components/Comments";
import { Container, Row, Col, Image, Stack } from "react-bootstrap";
import { LinkContainer } from "react-router-bootstrap";
import GenericModal from "../components/GenericModal";

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
      <Stack direction="horizontal">
        <div></div>
        <Stack>
          <div>
            {new Date(createdAt).toLocaleDateString("en-UK", {
              month: "short",
              year: "2-digit",
              day: "numeric",
            })}
          </div>
          <div>
            <LinkContainer to={`/profile/${userId}`}>
              <span>{userName}</span>
            </LinkContainer>
          </div>
        </Stack>
        {groupId > 0 && (
          <LinkContainer to={`/groups/${groupId}`}>
            <div className="text-end">{groupName}</div>
          </LinkContainer>
        )}
      </Stack>

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
