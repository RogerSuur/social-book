import React, { useState, useEffect } from "react";
import Comments from "./Comments";
import CreateComment from "./CreateComment";

const FeedPosts = (props) => {
  const { posts } = props;
  // const [postid, setPostId] = useState(null);
  const renderPost = (post) => {
    const { id, userId, content, createdAt, commentCount } = post;

    return (
      <div className="content-area" key={id}>
        <div>Post ID: {id}</div>
        <div className="row">UserId{userId}</div>
        <div className="row">{content}</div>
        <div className="row">{new Date(createdAt).toLocaleString("et-EE")}</div>
        {commentCount !== 0 ? (
          <div className="row">
            <p>{commentCount} comments</p>
            <Comments postid={id} />
          </div>
        ) : (
          <p>Be the first to leave a comment</p>
        )}
        <CreateComment />
      </div>
    );
  };

  const renderedPosts = posts.map(renderPost);

  return <div>{renderedPosts}</div>;
};

export default FeedPosts;
