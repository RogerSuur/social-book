import React, { useState, useEffect } from "react";
import axios from "axios";
import CreateComment from "./CreateComment";
import { Container, Image, Alert, Col, Button } from "react-bootstrap";
import GenericModal from "../components/GenericModal";
import Comment from "../components/Comment";

const Comments = ({ postId, commentCount }) => {
  const [comments, setComments] = useState([]);
  const [errMsg, setErrMsg] = useState(null);
  const [commentCountUpdate, setCommentsCountUpdate] = useState(commentCount);
  const [commentsToShow, setCommentsToShow] = useState(
    commentCount > 5 ? commentCount : 0
  );

  const [offset, setOffset] = useState(0);
  const [loading, setLoading] = useState(true);

  const handleCommentsUpdate = () => {
    setCommentsCountUpdate((prev) => prev + 1);
    setOffset(0);
    setLoading(!loading);
  };

  useEffect(() => {
    const abortController = new AbortController();
    const loadComments = async () => {
      try {
        await axios
          .get(`http://localhost:8000/comments/${postId}/${offset}`, {
            withCredentials: true,
            signal: abortController.signal,
          })
          .then((response) => {
            setComments((prevComments) => {
              const commentIds = new Set(
                prevComments.map((comment) => comment.id)
              );
              const newComments = response.data.filter(
                (comment) => !commentIds.has(comment.id)
              );
              const updatedComments = [...newComments, ...prevComments];
              const sortedComments = updatedComments.sort(
                (a, b) => new Date(b.createdAt) - new Date(a.createdAt)
              );
              return sortedComments;
            });
          });
      } catch (err) {
        if (err.response?.status === 404) {
          setErrMsg(err.message);
        }
      }
    };

    loadComments();

    return () => {
      abortController.abort();
    };
  }, [offset, loading]);

  const showMoreComments = () => {
    if (commentCountUpdate > 4) {
      setOffset(offset + 1);
      setCommentsToShow(commentsToShow - 5);
    }
  };

  const renderedComments = comments.map((comment, index) => (
    <Comment comment={comment} key={index} />
  ));

  return (
    <Container>
      {errMsg ? (
        <Alert variant="danger" className="text-center">
          {errMsg}
        </Alert>
      ) : (
        <div className="my-auto mt-2">
          <CreateComment
            postId={postId}
            onCommentsUpdate={handleCommentsUpdate}
          />
          {commentCountUpdate > 0 && <>{renderedComments}</>}
          {commentsToShow > 5 && (
            <div className="d-flex justify-content-center">
              <Button onClick={showMoreComments}>
                {commentsToShow - 5} more comment
                {commentsToShow - 5 === 1 ? "" : "s"}
              </Button>
            </div>
          )}
        </div>
      )}
    </Container>
  );
};

export default Comments;
