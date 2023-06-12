import React, { useState, useEffect } from "react";
import axios from "axios";
import CreateComment from "./CreateComment";

const Comments = ({ postId, commentCount, updateCommentCount }) => {
  const [comments, setComments] = useState([]);
  const [error, setError] = useState(null);
  // const { postid: id } = props.postid;
  const [commentsToShow, setCommentsToShow] = useState(commentCount);

  const [offset, setOffset] = useState(0);
  const [loading, setLoading] = useState(true);

  const handleCommentsUpdate = (add) => {
    console.log("handleCOmmentsUdate", add);
    setLoading(!loading);
    setCommentsToShow((prev) => prev + add);
    updateCommentCount(add);
    // setOffset(0);
  };

  useEffect(() => {
    const abortController = new AbortController();
    const loadComments = async () => {
      console.log(
        "loadmorecomments",
        "postId",
        postId,
        "offset",
        offset,
        "commentCount",
        commentCount,
        "commentsToShow",
        commentsToShow
      );
      // console.log("commentsToShow", commentsToShow);
      try {
        await axios
          .get(`http://localhost:8000/comments/${postId}/${offset}`, {
            withCredentials: true,
            signal: abortController.signal,
          })
          .then((response) => {
            // if (commentCount > 5) {
            setComments((prevComments) => {
              const commentIds = new Set(
                prevComments.map((comment) => comment.id)
              );
              console.log(commentIds);
              const newComments = response.data.filter(
                (comment) => !commentIds.has(comment.id)
              );
              console.log(commentIds, newComments);
              const updatedComments = [...newComments, ...prevComments];
              console.log(updatedComments);
              const sortedComments = updatedComments.sort(
                (a, b) => new Date(b.createdAt) - new Date(a.createdAt)
              );
              console.log(sortedComments);
              return sortedComments;
            });
            // } else {
            //   setComments(response.data);
            // }
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

  function showMoreComments(e) {
    console.log(e.target);
    if (commentCount > 4) {
      setOffset(offset + 1);
      setCommentsToShow(commentsToShow - 5);
    }
    // console.log("offset", offset);
    //here we need to increase the offset of comments
  }

  return (
    <>
      {error ? (
        <div className="error">{error}</div>
      ) : (
        <div className="content-area">
          <div className="row">
            <div className="column">
              {comments.map((comment) => (
                <>
                  <div key={comment.comment_id}>{comment.content}</div>
                  <div className="row">
                    <div className="column">
                      <p>
                        <small>
                          {new Date(comment.createdAt).toLocaleString("et-EE")}
                        </small>
                      </p>
                    </div>
                    <div className="column">{comment.userId}</div>
                  </div>

                  <hr />
                </>
              ))}
            </div>
          </div>
          {commentsToShow > 5 ? (
            <p>
              <a onClick={showMoreComments}>
                {commentsToShow - 5} more comment
                {commentsToShow - 5 === 1 ? "" : "s"}
              </a>
            </p>
          ) : null}
          <CreateComment
            postId={postId}
            onCommentsUpdate={handleCommentsUpdate}
          />
        </div>
      )}
    </>
  );
};

export default Comments;
