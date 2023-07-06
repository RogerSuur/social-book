import React, { useEffect } from "react";
import GroupSidebar from "../components/GroupSidebar";
import { GROUP_PAGE_URL } from "../utils/routes";
import axios from "axios";

const GroupPage = () => {
  const [group, setGroup] = useState({});
  const { id } = useParams();

  useEffect(() => {
    const loadGroup = async () => {
      await axios
        .get(GROUP_PAGE_URL + id, {
          withCredentials: true,
        })
        .then((response) => {
          setGroup(response.data);
        });
    };
    loadGroup();
  }, [id]);

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
      <h1>Group</h1>
      <GroupSidebar />
    </>
  );
};

export default GroupPage;
