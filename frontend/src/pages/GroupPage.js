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
  const [followers, setFollowers] = useState([]);
  const [formData, setFormData] = useState([]);

  const [open, setOpen] = useState(false);

  const handleOpen = () => {
    setOpen(!open);
  };

  useEffect(() => {
    setOpen(false);
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

  useEffect(() => {
    const fetchFollowers = async () => {
      try {
        const response = await axios.get("http://localhost:8000/followers", {
          withCredentials: true,
        });
        setFollowers(response.data);
      } catch (err) {
        console.error(err);
      }
    };
    fetchFollowers();
  }, []);

  const handleChange = (event) => {
    const { value, checked } = event.target;

    setFormData((prevFormData) => {
      const newFormData = [...prevFormData];

      if (!checked) {
        const index = newFormData.indexOf(value);
        if (index !== -1) {
          newFormData.splice(index, 1);
        }
      } else {
        newFormData.push(value);
      }
      return newFormData;
    });
  };

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

          <div>
            <button onClick={handleOpen}>Add members</button>
            {open
              ? followers.map((follower) => (
                  <div key={follower.id}>
                    <label htmlFor={`receiver_${follower.id}`}>
                      <input
                        type="checkbox"
                        name="newMembers"
                        onChange={handleChange}
                        value={follower.id}
                      />
                      {follower.firstName} {follower.lastName}
                    </label>
                  </div>
                ))
              : null}
          </div>
          <GroupMembers groupId={groupId} />
        </div>
      )}
    </>
  );
};

export default GroupPage;
