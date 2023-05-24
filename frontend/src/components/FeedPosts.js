import React, { useState, useEffect } from "react";
import Comments from "./Comments";
import List from "./List";
import CreateComment from "./CreateComment";

const FEEDPOSTS_URL = `http://localhost:8000/feedposts/`;

const FeedPosts = (props) => {
  const { offset } = props;
  const [postid, setPostId] = useState(null);
  const mapFeedPosts = (post) => {
    console.log("COMMENTCOUNT", post.commentCount);
    const numComments = 3;

    return (
      <>
        <div className="content-area" key={post.id}>
          <div className="row">{post.userId}</div>
          <div className="row">{post.content}</div>
          <div className="row">
            {new Date(post.createdAt).toLocaleString("et-EE")}
          </div>
          {/* ADD A LINK TO comments if not empty*/}
          {post.commentCount !== 0 && (
            <div className="row">
              <a href={`/comments/${post.id}`}>{post.commentCount} comments</a>
              {/* <Comments postid={post.id}/> */}
              <Comments />
            </div>
          )}
          <CreateComment />
        </div>
      </>
    );
  };

  return <List url={`${FEEDPOSTS_URL}${offset}`} mapFunction={mapFeedPosts} />;
};

export default FeedPosts;
