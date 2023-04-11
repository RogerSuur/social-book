import React from "react";
import { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import CreatePost from "./CreatePost";
import axios from "axios";

const Posts = () => {
  const [posts, setPosts] = useState([]);
  const [categories, setCategories] = useState([]);
  const [loading, setLoading] = useState(false);

  const loader = () => {
    setLoading((prevLoading) => !prevLoading);
  };

  useEffect(() => {
    const loadPosts = async () => {
      await axios
        .get("http://localhost:8000/posts", {
          withCredentials: true,
        })
        .then((response) => setPosts(response.data));
    };

    loadPosts();
  }, [loading]);

  useEffect(() => {
    const loadCategories = async () => {
      await axios
        .get("http://localhost:8000/categories", {
          withCredentials: true,
        })
        .then((response) => setCategories(response.data));
    };

    loadCategories();
  }, []);

  return (
    <>
      <CreatePost handler={loader} />
      <div className="content-area">
        {categories.map((category) => (
          <div key={category.category_id} title={category.description}>
            <Link to={`/categories/${category.category_id}`}>
              {category.title}
              <br />
            </Link>
          </div>
        ))}
      </div>
      ;
      <div>
        {posts.map((post) => (
          <div className="content-area" key={post.post_id}>
            <Link to={`/posts/${post.post_id}`}>
              {post.title}
              <br />
            </Link>
            <div className="row">{post.body}</div>
            <div className="row">
              {new Date(post.post_datetime).toLocaleString("et-EE")}
            </div>
          </div>
        ))}
      </div>
    </>
  );
};

export default Posts;
