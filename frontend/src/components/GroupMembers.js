import { useState, useEffect } from "react";
import { makeRequest } from "../services/makeRequest";
import Modal from "./Modal";
import { Link } from "react-router-dom";
import ImageHandler from "../utils/imageHandler";

const GroupMembers = ({ groupId }) => {
  const [groupMembers, setGroupMembers] = useState([]);
  const [error, setError] = useState(null);
  const [modalOpen, setModalOpen] = useState(false);

  console.log(groupMembers, "MEMBERS");

  useEffect(() => {
    const loadMembers = async () => {
      try {
        const response = await makeRequest(`/groupmembers/${groupId}`);
        if (response !== null) {
          setGroupMembers(response);
        }
      } catch (error) {
        setError(error);
        console.log(error.response?.status);
      }
    };
    loadMembers();
  }, [groupId]);

  const groupMembersMap = groupMembers.map((member, index) => (
    <div key={index}>
      <Link to={`/profile/${member.id}`}>
        {ImageHandler(member.imagePath, "defaultuser.jpg", "profile-image")}
        <p>{member.nickname}</p>
      </Link>
    </div>
  ));

  const openModal = () => {
    setModalOpen(true);
  };

  const closeModal = () => {
    setModalOpen(false);
  };

  return (
    <>
      {groupMembers.length > 0 ? (
        <button onClick={openModal}>
          <p>{groupMembers.length} members</p>

          <Modal className="modal" open={modalOpen} onClose={closeModal}>
            {groupMembersMap}
          </Modal>
        </button>
      ) : (
        <p>Apply to become a member of this group</p>
      )}
    </>
  );
};

export default GroupMembers;
