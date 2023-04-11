import React, { useState, useEffect } from "react";
import { useParams, Link } from "react-router-dom";
import Comments from "./Comments";
import axios from "axios";

const Post = () => {
  const [post, setPost] = useState({ categories: [] });
  const { id } = useParams();

  useEffect(() => {
    const loadPost = async () => {
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
          <h1>{post.title}</h1>
          <h2>{post.body}</h2>
          <sub>{post.username}</sub>
          <sub>{new Date(post.post_datetime).toLocaleString("et-EE")}</sub>
          <div>
            Posted in:
            {post.categories.map((category) => (
              <div className="content-area" key={category.value}>
                <Link to={`/categories/${category.value}`}>
                  {category.label}
                  <br />
                </Link>
              </div>
            ))}
          </div>
        </div>
      </div>
      <Comments />
    </>
  );
};

export default Post;
