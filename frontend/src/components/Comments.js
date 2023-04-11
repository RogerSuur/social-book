import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import axios from "axios";
import CreateComment from "./CreateComment";

const Comments = () => {
  const [comments, setComments] = useState([]);
  const { id } = useParams();
  const [first, setFirst] = useState(false);
  const [loading, setLoading] = useState(false);

  const loader = () => {
    setLoading((prevLoading) => !prevLoading);
  };

  useEffect(() => {
    const loadComments = async () => {
      try {
        await axios
          .get(`http://localhost:8000/comments/${id}`, {
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
      <CreateComment handler={loader} />
      {first && <h2>Be the first to leave a comment</h2>}
      {comments.length > 0 && (
        <div className="content-area">
          <div className="row">
            <div className="column">
              {comments.map((comment) => (
                <>
                  <div key={comment.comment_id}>{comment.body}</div>
                  <div className="row">
                    <div className="column">
                      <p>
                        <small>
                          {new Date(comment.comment_datetime).toLocaleString(
                            "et-EE"
                          )}
                        </small>
                      </p>
                    </div>
                    <div className="column">{comment.username}</div>
                  </div>

                  <hr />
                </>
              ))}
            </div>
          </div>
        </div>
      )}
    </>
  );
};

export default Comments;
