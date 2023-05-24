import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import axios from "axios";
import CreateComment from "./CreateComment";

const Comments = (props) => {
  const [comments, setComments] = useState([]);
  const { postid: id } = props;
  const [first, setFirst] = useState(false);
  const [loading, setLoading] = useState(false);

  const loader = () => {
    setLoading((prevLoading) => !prevLoading);
  };

  // localhost:8000/comments/postid/offset

  useEffect(() => {
    console.log("POST ID IN COMMENTS", props.id);
    const loadComments = async () => {
      try {
        await axios
          .get(`http://localhost:8000/comments/${props.id}/0`, {
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
  }, [id, loading]);

  return (
    <>
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
