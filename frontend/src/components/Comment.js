import React from "react";
import { Image } from "react-bootstrap";
import { LinkContainer } from "react-router-bootstrap";
import { ShortDatetime } from "../utils/datetimeConverters";

const Comment = ({ comment }) => {
  return (
    <>
      {comment.imagePath && (
        <Image
          fluid
          className="profile-pic"
          src={`${process.env.PUBLIC_URL}/images/${comment.imagePath}`}
        />
      )}
      <div key={comment.id}>{comment.content}</div>
      <div className="row">
        <div className="column">
          <p>
            <small>{ShortDatetime(comment.createdAt)}</small>
          </p>
        </div>
        <LinkContainer to={`/profile/${comment?.userId}`}>
          <>{comment?.userName}</>
        </LinkContainer>
      </div>
      <hr />
    </>
  );
};

export default Comment;
