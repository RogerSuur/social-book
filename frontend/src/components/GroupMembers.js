import { useState, useEffect } from "react";
import { makeRequest } from "../services/makeRequest";
import Modal from "./Modal";
import { Link } from "react-router-dom";

const GroupMembers = ({ groupId }) => {
  const [groupMembers, setGroupMembers] = useState([]);
  const [error, setError] = useState(null);
  const [modalOpen, setModalOpen] = useState(false);

  useEffect(() => {
    const loadMembers = async () => {
      try {
        const response = await makeRequest(`/groupmembers/${groupId}`);
        if (response !== null) {
          setGroupMembers(response);
          console.log(response);
        }
      } catch (error) {
        setError(error.message);
      }
    };
    loadMembers();
  }, []);

  const groupMembersMap = groupMembers.map((member, index) => (
    <div key={index}>
      <p>User ID: {member.userId}</p>
      <Link to={`/profile/${member.userId}`}>
        <p>{member.userName}</p>
      </Link>
      {/* <p>Username: {member.userName}</p> */}
      <p>Image Path: {member.imagePath}</p>
      <img
        className="profile-image"
        src={
          member.imagePath
            ? `${process.env.PUBLIC_URL}/images/${member.imagePath}`
            : `${process.env.PUBLIC_URL}/images/defaultuser.jpg`
        }
        alt={`${member.userName}`}
      />
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
        <>
          <button onClick={openModal}>
            <p>{groupMembers.length} members</p>

            <Modal open={modalOpen} onClose={closeModal}>
              {groupMembersMap}
            </Modal>
          </button>
        </>
      ) : (
        <p>Apply to become a member of this group</p>
      )}
      ;
    </>
  );
};

export default GroupMembers;
