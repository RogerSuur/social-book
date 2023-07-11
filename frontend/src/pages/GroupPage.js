import React, { useEffect, useState } from "react";
import GroupSidebar from "../components/GroupSidebar";
import { GROUP_PAGE_URL } from "../utils/routes";
import axios from "axios";
import { useParams } from "react-router-dom";
import GroupMembers from "../components/GroupMembers";
import ImageHandler from "../utils/imageHandler";

const GroupPage = () => {
  const [group, setGroup] = useState({});
  const { groupId } = useParams();
  const [error, setError] = useState(null);

  useEffect(() => {
    const loadGroup = async () => {
      try {
        await axios
          .get(GROUP_PAGE_URL + groupId, {
            withCredentials: true,
          })
          .then((response) => {
            setGroup(response.data);
          });
      } catch (err) {
        setError(err.message);
      }
    };
    loadGroup();
  }, [groupId]);

  const image = () =>
    ImageHandler(group.imagePath, "defaultgroup.png", "group-image");

  return (
    <>
      {<GroupSidebar />}
      {error ? (
        <div className="error">{error}</div>
      ) : (
        <div className="group-page">
          {image()}
          <h1>{group.title}</h1>
          <p>{group.description}</p>
          <GroupMembers groupId={groupId} />
        </div>
      )}
    </>
  );
};

export default GroupPage;
