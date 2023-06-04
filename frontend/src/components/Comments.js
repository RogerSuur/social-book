import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import axios from "axios";
import CreateComment from "./CreateComment";

const Comments = (postId) => {
  const [comments, setComments] = useState([]);
  // const { postid: id } = props.postid;
  const [first, setFirst] = useState(true);
  const [loading, setLoading] = useState(false);
  const [hasMore, setHasMore] = useState(false);
  let offset = 0;

  const loader = () => {
    setLoading((prevLoading) => !prevLoading);
  };

  useEffect(() => {
    // console.log("useeffect comments postId", postId.postid);
    const loadComments = async () => {
      try {
        await axios
          .get(`http://localhost:8000/comments/${postId.postid}/0`, {
            withCredentials: true,
          })
          .then((response) => {
            setComments(response.data);
            setFirst(false);
          });
      } catch (err) {
        if (err.response?.status === 404) {
          setFirst(true);
        }
      }
    };

    loadComments();
  }, [loading]);

  return (
    <>
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
      </div>
    </>
  );
};

export default Comments;
