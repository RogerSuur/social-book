import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import Comments from "./Comments";
import axios from "axios";

const Post = () => {
  const [post, setPost] = useState({ categories: [] });
  const { id } = useParams();

  useEffect(() => {
    const loadPost = async () => {
      console.log("loadPost ComponentPost");
      await axios
        .get(`http://localhost:8000/posts/${id}`, {
          withCredentials: true,
        })
        .then((response) => setPost(response.data));
    };

    loadPost();
  }, []);

  return (
    <>
      <div className="content-area">
        <div className="column-title">
          <h2>{post.body}</h2>
          <sub>{post.username}</sub>
          <sub>{new Date(post.post_datetime).toLocaleString("et-EE")}</sub>
        </div>
      </div>
      <Comments />
    </>
  );
};

export default Post;
