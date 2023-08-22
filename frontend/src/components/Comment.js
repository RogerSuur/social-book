import React from "react";
import { Image } from "react-bootstrap";
import { LinkContainer } from "react-router-bootstrap";

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
            <small>{new Date(comment.createdAt).toLocaleString("et-EE")}</small>
          </p>
        </div>
        <div className="column">
          <LinkContainer to={`/profile/${comment?.userId}`}>
            <>{comment?.userName}</>
          </LinkContainer>
        </div>
      </div>
      <hr />
    </>
  );
};

export default Comment;
