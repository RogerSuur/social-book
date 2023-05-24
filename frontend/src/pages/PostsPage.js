import React from "react";
import { useState, useEffect } from "react";
import FeedPosts from "../components/FeedPosts";
import CreatePost from "../components/CreatePost";
import axios from "axios";
import { makeRequest } from "../services/makeRequest";

const Posts = () => {
  const [posts, setPosts] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  let offset = 0;

  const loader = () => {
    setLoading((prevLoading) => !prevLoading);
  };

  useEffect(() => {
    const loadPosts = async () => {
      return makeRequest(`feedposts/${offset}`);

      // try {
      //   const response = await axios.get(`http://localhost:8000/feedposts/${offset}`, {
      //     withCredentials: true,
      //   });
      //   setPosts(response.data);
      // } catch (error) {
      //   setError(error.message);
      // }
    };

    loadPosts();
  }, [loading, offset]);

  return (
    <>
      {/* kutsub v'lja CreatePost componendi handler */}
      <CreatePost handler={loader} />
      {error ? (
        <div className="error">{error}</div>
      ) : (
        <div className="content-area">
          <FeedPosts offset={offset} />
        </div>
      )}
    </>
  );
};

export default Posts;
