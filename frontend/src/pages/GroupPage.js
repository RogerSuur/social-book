import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import GroupSidebar from "../components/GroupSidebar";
import { GROUP_PAGE_URL } from "../utils/routes";
import axios from "axios";
import { useParams } from "react-router-dom";
import { useState } from "react";

const GroupPage = () => {
  const [group, setGroup] = useState({});
  const { groupId } = useParams();

  useEffect(() => {
    console.log(groupId);
    const loadGroup = async () => {
      console.log("LOAD GROUP", GROUP_PAGE_URL + groupId);
      await axios
        .get(GROUP_PAGE_URL + groupId, {
          withCredentials: true,
        })
        .then((response) => {
          setGroup(response.data);
          console.log(response.data);
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
        alt={`groupId: ${groupId}`}
      ></img>
      <h1>{group.title}</h1>
      <p>{group.description}</p>
      <GroupSidebar />
    </>
  );
};

export default GroupPage;
