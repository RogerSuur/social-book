import React, { useState, useEffect } from "react";
import axios from "axios";
import CreateComment from "./CreateComment";
import { Link } from "react-router-dom";

const Comments = ({ postId, commentCount, showCreateComment }) => {
  const [comments, setComments] = useState([]);
  const [error, setError] = useState(null);
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
          setError(err.message);
        }
      }
    };

    loadComments();

    return () => {
      abortController.abort();
    };
  }, [offset, loading]);

  function showMoreComments() {
    if (commentCountUpdate > 4) {
      setOffset(offset + 1);
      setCommentsToShow(commentsToShow - 5);
    }
  }

  console.log("COMMENTS: ", comments);

  return (
    <>
      {error ? (
        <div className="error">{error}</div>
      ) : (
        <div className="content-area">
          {commentCountUpdate > 0 && (
            <div className="row">
              <div className="column">
                {comments.map((comment, index) => (
                  <div key={index}>
                    <div key={comment.id}>{comment.content}</div>
                    <div className="row">
                      <div className="column">
                        <p>
                          <small>
                            {new Date(comment.createdAt).toLocaleString(
                              "et-EE"
                            )}
                          </small>
                        </p>
                      </div>
                      <div className="column">
                        <Link to={`/profile/${comment?.userId}`}>
                          {comment?.userName}
                        </Link>
                      </div>
                    </div>
                    <hr />
                  </div>
                ))}
              </div>
            </div>
          )}
          {commentsToShow > 5 && (
            <p>
              <button onClick={showMoreComments}>
                {commentsToShow - 5} more comment
                {commentsToShow - 5 === 1 ? "" : "s"}
              </button>
            </p>
          )}
          {!commentCountUpdate && <p>Be the first to leave a comment</p>}
          {
            <CreateComment
              postId={postId}
              onCommentsUpdate={handleCommentsUpdate}
            />
          }
        </div>
      )}
    </>
  );
};

export default Comments;
