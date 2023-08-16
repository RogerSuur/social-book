import React, { useEffect, useState } from "react";
import { GROUP_PAGE_URL, ADD_GROUP_MEMBERS_URL } from "../utils/routes";
import axios from "axios";
import { useParams } from "react-router-dom";
import GroupMembers from "../components/GroupMembers";
import ImageHandler from "../utils/imageHandler";
import FeedPosts from "../components/FeedPosts.js";
import Select from "react-select";
import Modal from "../components/Modal";
import AvatarUpdater from "../components/AvatarUpdater";
import Events from "../components/Events";
import GroupRequestButton from "../components/GroupRequestButton.js";
import CreateGroupPosts from "../components/CreateGroupPosts.js";

const GroupPage = () => {
  const [group, setGroup] = useState({});
  const { id } = useParams();
  const [error, setError] = useState(null);
  const [followers, setFollowers] = useState([]);
  const [formData, setFormData] = useState([]);
  const [submitting, setSubmitting] = useState(false);
  const [modalOpen, setModalOpen] = useState(false);
  const [open, setOpen] = useState(false);
  const [reload, setReload] = useState(false);

  const handlePostUpdate = () => {
    setReload(!reload);
  };

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
          .get(GROUP_PAGE_URL + id, {
            withCredentials: true,
          })
          .then((response) => {
            console.log("RESP: ", response.data);
            setGroup(response.data);
          });
      } catch (err) {
        setError(err.message);
      }
    };
    loadGroup();
  }, [id, modalOpen]);

  useEffect(() => {
    const fetchFollowers = async () => {
      try {
        const response = await axios.get(ADD_GROUP_MEMBERS_URL + `/${id}`, {
          withCredentials: true,
        });
        console.log("ADD GROUP MEMBERS: ", response?.data);
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
      await axios.post(
        ADD_GROUP_MEMBERS_URL,
        JSON.stringify({ groupId: +id, userIds: formData }),
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
      {/* <div style={{ "max-width": `100px` }}>
        <Events groupId={+id} />
      </div> */}
      {error ? (
        <div className="error">{error}</div>
      ) : (
        <div className="group-page">
          {image()}
          <h1>{group.title}</h1>
          <p>{group.description}</p>
          {group.isMember ? (
            <>
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
              <GroupMembers groupId={+id} />
              <div className="profile-actions">
                <Modal open={modalOpen} onClose={handleModalClose}>
                  <AvatarUpdater
                    url={`http://localhost:8000/groups/${id}/avatar`}
                    onUploadSuccess={handleModalClose}
                  />
                </Modal>
                <button onClick={handleModalClick}>Upload New Image</button>
              </div>
              <CreateGroupPosts groupId={id} onPostsUpdate={handlePostUpdate} />
              <FeedPosts url={`/groupfeed/${id}`} key={id} reload={reload} />
            </>
          ) : (
            <GroupRequestButton groupid={+id} />
          )}
        </div>
      )}
    </>
  );
};

export default GroupPage;
