import React from "react";
import { useState, useEffect } from "react";
import FeedPosts from "../components/FeedPosts";
import CreatePost from "../components/CreatePost";
import axios from "axios";

const Posts = () => {
  const [posts, setPosts] = useState([]);
  const [loading, setLoading] = useState(false);

  const loader = () => {
    setLoading((prevLoading) => !prevLoading);
  };

  useEffect(() => {
    const loadPosts = async () => {
      await axios
        .get("http://localhost:8000/feedposts", {
          withCredentials: true,
        })
        .then((response) =>  {
         setPosts(response.data)}
        );
    };

    loadPosts();
  }, [loading]);

  return (
    <>
      <CreatePost handler={loader} />
      <div className="content-area">
       <FeedPosts/>
      </div>
    </>
  );
};

export default Posts;
