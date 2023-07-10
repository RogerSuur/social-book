import React, { useEffect, useState } from "react";
import GroupSidebar from "../components/GroupSidebar";
import { GROUP_PAGE_URL } from "../utils/routes";
import axios from "axios";
import { useParams } from "react-router-dom";
import GroupMembers from "../components/GroupMembers";

const GroupPage = () => {
  const [group, setGroup] = useState({});
  const { groupId } = useParams();

  useEffect(() => {
    console.log("group page request", GROUP_PAGE_URL, groupId);
    const loadGroup = async () => {
      await axios
        .get(GROUP_PAGE_URL + groupId, {
          withCredentials: true,
        })
        .then((response) => {
          console.log("response.data:", response.data);
          setGroup(response.data);
        });
    };
    loadGroup();
  }, [groupId]);

  return (
    <>
      <img
        style={{
          width: "20vw",
          height: "20vw",
          objectFit: "cover",
          objectPosition: "0% 100%",
        }}
        src={`../images/${group.imagePath}`}
        alt={`groupId: ${groupId}`}
      ></img>
      <h1>{group.title}</h1>
      <p>{group.description}</p>
      <GroupMembers groupId={groupId} />
      <GroupSidebar />
    </>
  );
};

export default GroupPage;
