import React from "react";
import { useState, useEffect } from "react";
import FeedPosts from "../components/FeedPosts";
import CreatePost from "../components/CreatePost";
import { makeRequest } from "../services/makeRequest";
import GroupSidebar from "../components/GroupSidebar";
import CreateGroupPosts from "../components/CreateGroupPosts";

const Posts = ({ groupId, showGroupSidebar, showCreatePost, url, reload }) => {
  const [posts, setPosts] = useState([]);
  const [error, setError] = useState(null);
  const [offset, setOffset] = useState(0);
  const [loadMore, setLoadMore] = useState(true);
  const [hasMore, setHasMore] = useState(false);

  const handlePostUpdate = () => {
    // console.log("handlePostUpdate", posts, offset);
    setOffset(0);
    setLoadMore(!loadMore);
    setPosts([]);
  };

  const handlePageChange = (postId) => {
    setOffset(postId);
  };

  useEffect(() => {
    const abortController = new AbortController();
    const loadPosts = async () => {
      try {
        const response = await makeRequest(`${url}/${offset}`, {
          signal: abortController.signal,
        });
        setPosts((prevPosts) => {
          return [...prevPosts, ...response];
        });
        setHasMore(response.length > 0);
      } catch (error) {
        setError(error.message);
      }
    };
    loadPosts();

    return () => {
      abortController.abort();
    };
  }, [offset, loadMore, reload]);

  console.log("GROUP ID: ", groupId);

  return (
    <>
      {showGroupSidebar && <GroupSidebar />}
      <div className="content-as">
        {showCreatePost && !groupId && (
          <CreatePost onPostsUpdate={handlePostUpdate} />
        )}

        {showCreatePost && groupId && (
          <CreateGroupPosts
            onPostsUpdate={handlePostUpdate}
            groupId={groupId}
          />
        )}

        {error ? (
          <div className="error">{error}</div>
        ) : (
          <FeedPosts
            posts={posts}
            hasMore={hasMore}
            onLoadMore={handlePageChange}
          />
        )}
      </div>
    </>
  );
};

export default Posts;
