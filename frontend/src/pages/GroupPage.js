import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import GroupSidebar from "../components/GroupSidebar";
import { GROUP_PAGE_URL } from "../utils/routes";
import axios from "axios";

const GroupPage = () => {
  const [group, setGroup] = useState({});
  const { groupId } = useParams();

  useEffect(() => {
    console.log(groupId);
    const loadGroup = async () => {
      await axios
        .get(GROUP_PAGE_URL + groupId, {
          withCredentials: true,
        })
        .then((response) => {
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
        src={`images/${group.avatarImage}`}
        alt={`${group.name}`}
      ></img>
      <h1>GroupPage</h1>
      <p>{group.name}</p>
      <GroupSidebar />
    </>
  );
};

export default GroupPage;
