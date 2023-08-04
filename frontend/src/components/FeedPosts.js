import React, { useState, useRef, useCallback } from "react";
import Comments from "./Comments";
import { Link } from "react-router-dom";

const FeedPosts = ({ posts, onLoadMore, hasMore }) => {
  const observer = useRef();
  const [isPostsLoading, setPostsLoading] = useState(false);

  async function toggleSpinner() {
    setPostsLoading((prev) => !prev);
    setTimeout(function () {
      setPostsLoading((prev) => !prev);
    }, 800);
  }

  const lastPostElementRef = useCallback((node) => {
    if (observer.current) {
      observer.current.disconnect();
    }

    observer.current = new IntersectionObserver((entries) => {
      if (entries[0].isIntersecting) {
        toggleSpinner();
        const postId = node.getAttribute("data-post-id"); // Get the post ID from the element attribute
        onLoadMore(postId);
      }
    });

    if (node) {
      observer.current.observe(node);
    }
  }, []);

  const renderPost = (post, index) => {
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
    const isLastPost = index === posts.length - 1;

    return (
      <div
        className="content-area"
        key={id}
        ref={isLastPost ? lastPostElementRef : null}
        data-post-id={id}
      >
        <div>
          <Link to={`/groups/${groupId}`}>{groupName}</Link>
        </div>
        <div className="row3">
          <Link to={`/profile/${userId}`}>{userName}</Link>
        </div>
        {imagePath && (
          <img
            className="profile-pic"
            src={`${process.env.PUBLIC_URL}/images/${imagePath}`}
          />
        )}
        <div className="row2">{content}</div>;
        <div className="row">{new Date(createdAt).toLocaleString("et-EE")}</div>
        <div className="comment-section">
          <Comments postId={id} commentCount={commentCount} />
        </div>
      </div>
    );
  };

  const renderedPosts = posts?.map(renderPost);
  return (
    <>
      <div>
        {isPostsLoading && <div className="spinner" />}
        {renderedPosts}
        {!hasMore && "No more posts to show"}
      </div>
    </>
  );
};

export default FeedPosts;
