import React, { useState, useRef, useCallback } from "react";
import Comments from "./Comments";

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
        onLoadMore();
      }
    });

    if (node) {
      observer.current.observe(node);
    }
  }, []);

  const renderPost = (post, index) => {
    const {
      id,
      // userId,
      imagePath,
      userName,
      content,
      createdAt,
      commentCount,
    } = post;
    const isLastPost = index === posts.length - 1;

    return (
      <div
        className="content-area"
        key={id}
        ref={isLastPost ? lastPostElementRef : null}
      >
        <div>Post ID: {id}</div>
        <div className="row">{userName}</div>
        <div className="row">
          {content}
          {imagePath}
        </div>
        <div className="row">{new Date(createdAt).toLocaleString("et-EE")}</div>
        <div className="row">
          <Comments postId={id} commentCount={commentCount} />
        </div>
      </div>
    );
  };

  const renderedPosts = posts.map(renderPost);
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
