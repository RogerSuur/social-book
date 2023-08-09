import React, { useState, useRef, useEffect, useCallback } from "react";
import Comments from "./Comments";
import { makeRequest } from "../services/makeRequest";
import { Link } from "react-router-dom";
import Container from "react-bootstrap/Container";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import Image from "react-bootstrap/Image";
import { LinkContainer } from "react-router-bootstrap";

const FeedPosts = ({ url, reload }) => {
  const observer = useRef();
  const [posts, setPosts] = useState([]);
  const [error, setError] = useState(null);
  const [isPostsLoading, setPostsLoading] = useState(false);
  const [offset, setOffset] = useState(0);

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
      <Container
        fluid
        className="content-area border"
        key={id}
        ref={isLastPost ? lastPostElementRef : null}
        data-post-id={id}
      >
        <div className="d-flex justify-content-between">
          <LinkContainer to={`/profile/${userId}`}>
            <span>{userName}</span>
          </LinkContainer>
          {groupId > 0 && (
            <LinkContainer to={`/groups/${groupId}`}>
              <span>Posted in {groupName}</span>
            </LinkContainer>
          )}
        </div>
        {imagePath && (
          <Image
            fluid
            className="post-img"
            src={`${process.env.PUBLIC_URL}/images/${imagePath}`}
          />
        )}
        <div>{content}</div>
        <Col>{new Date(createdAt).toLocaleString("et-EE")}</Col>
        <Col className="comment-section">
          <Comments postId={id} commentCount={commentCount} />
        </Col>
      </Container>
    );
  };

  const renderedPosts = posts?.map(renderPost);
  return (
    <Container fluid>
      {isPostsLoading && <div className="spinner" />}
      {renderedPosts}
    </Container>
  );
};

export default FeedPosts;
