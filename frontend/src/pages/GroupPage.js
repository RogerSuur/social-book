import React, { useEffect, useState } from "react";
import GroupSidebar from "../components/GroupSidebar";
import { GROUP_PAGE_URL, ADD_GROUP_MEMBERS_URL } from "../utils/routes";
import axios from "axios";
import { useParams } from "react-router-dom";
import GroupMembers from "../components/GroupMembers";
import ImageHandler from "../utils/imageHandler";
import Posts from "../pages/PostsPage.js";
import Select from "react-select";
import Modal from "../components/Modal";
import AvatarUpdater from "../components/AvatarUpdater";
import Events from "../components/Events";

const GroupPage = () => {
  const [group, setGroup] = useState({});
  const { groupId } = useParams();
  const [error, setError] = useState(null);
  const [followers, setFollowers] = useState([]);
  const [formData, setFormData] = useState([]);
  const [submitting, setSubmitting] = useState(false);
  const [modalOpen, setModalOpen] = useState(false);

  const [open, setOpen] = useState(false);

  const handleOpen = () => {
    setOpen(!open);
  };

  const handleModalClose = () => {
    setModalOpen(false);
  };

  const handleModalClick = () => {
    setModalOpen(true);
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
  }, [groupId, modalOpen]);

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

  const handleSelectChange = (selectedOptions) => {
    const selectedValues = selectedOptions.map((option) => option.value);
    setFormData(selectedValues);
  };

  const followersOptions = followers.map((follower) => ({
    value: follower.id,
    label: `${follower.firstName} ${follower.lastName}`,
  }));

  const image = () =>
    ImageHandler(group.imagePath, "defaultgroup.png", "group-image");

  const handleSubmit = async () => {
    try {
      console.log(formData);
      await axios.post(
        ADD_GROUP_MEMBERS_URL,
        JSON.stringify({ groupId: Number(groupId), userIds: formData }),
        {
          withCredentials: true,
        }
      );
      setFormData([]);
    } catch (err) {
      console.error(err);
    }
    setOpen(false);
  };

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
            {open && (
              <>
                <Select
                  options={followersOptions}
                  isMulti
                  onChange={handleSelectChange}
                />
                <button onClick={handleSubmit} disabled={submitting}>
                  Submit
                </button>
              </>
            )}
          </div>
          <GroupMembers groupId={groupId} />
          <div className="profile-actions">
            <Modal open={modalOpen} onClose={handleModalClose}>
              <AvatarUpdater
                url={`http://localhost:8000/groups/${groupId}/avatar`}
                onUploadSuccess={handleModalClose}
              />
            </Modal>
            <button onClick={handleModalClick}>Upload New Image</button>
          </div>
          <Posts
            forGroupPage={true}
            showGroupSidebar={false}
            showCreatePost={true}
            url={`/groupfeed/${groupId}`}
          />
          <Events groupId={groupId} />
        </div>
      )}
    </>
  );
};

export default GroupPage;
