import React, { useState, useEffect } from "react";
import Comments from "./Comments";
import List from "./List";
import CreateComment from "./CreateComment";

const FeedPosts = (props) => {
  const { posts } = props;
  // const [postid, setPostId] = useState(null);
  const mapFeedPosts = (post) => {
    // console.log("FEEDposts numcomments", numComments);

    return (
      <>
        <div className="content-area" key={post.id}>
          Postid{post.id}
          <div className="row">UserId{post.userId}</div>
          <div className="row">{post.content}</div>
          <div className="row">
            {new Date(post.createdAt).toLocaleString("et-EE")}
          </div>
          {post.commentCount !== 0 ? (
            <div className="row">
              <p>{post.commentCount} comments</p>
              <Comments postid={post.id} />
            </div>
          ) : (
            <p>Be the first to leave a comment</p>
          )}
          <CreateComment />
        </div>
      </>
    );
  };

  const renderedPosts = posts.map(mapFeedPosts);

  return <div>{renderedPosts}</div>;
};

export default FeedPosts;
