import React, { useState, useRef, useEffect, useCallback } from "react";
import Comments from "./Comments";
import { makeRequest } from "../services/makeRequest";
import { Link } from "react-router-dom";

const FeedPosts = ({ url, reload }) => {
  const observer = useRef();
  const [posts, setPosts] = useState([]);
  const [error, setError] = useState(null);
  const [isPostsLoading, setPostsLoading] = useState(false);
  const [offset, setOffset] = useState(0);
  const [hasMore, setHasMore] = useState(false);

  const handlePageChange = (postId) => {
    setOffset(postId);
  };

  useEffect(() => {
    setPosts([]);
    setOffset(0);
  }, [reload]);

  useEffect(() => {
    const abortController = new AbortController();
    const loadPosts = async () => {
      try {
        const response = await makeRequest(`${url}/${offset}`, {
          signal: abortController.signal,
        });
        setPosts((prevPosts) => {
          return [...prevPosts, ...response];
        });
        setHasMore(response.length > 0);
      } catch (error) {
        setError(error.message);
      }
    };
    loadPosts();

    return () => {
      abortController.abort();
    };
  }, [offset, reload]);

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
        handlePageChange(postId);
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
