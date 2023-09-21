import React from "react";
import { Image } from "react-bootstrap";
import { LinkContainer } from "react-router-bootstrap";
import { ShortDatetime } from "../utils/datetimeConverters";
import { Container, Row, Col, Stack } from "react-bootstrap";

const Comment = ({ comment }) => {
  return (
    <>
      <Container fluid className="bg-light p-3 mt-3 border rounded">
        <Stack>
          <Row>
            <LinkContainer to={`/profile/${comment?.userId}`}>
              <strong>{comment?.userName}</strong>
            </LinkContainer>
          </Row>
          <Row>
            <div>{comment.content}</div>
          </Row>
        </Stack>
        {comment.imagePath && (
          <Row className="mt-1">
            <Image
              fluid
              src={`${process.env.PUBLIC_URL}/images/${comment.imagePath}`}
            />
          </Row>
        )}
      </Container>

      <p>
        <small className="text-muted">{ShortDatetime(comment.createdAt)}</small>
      </p>
    </>
  );
};

export default Comment;
