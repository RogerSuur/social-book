import { useState, useEffect } from "react";
import { makeRequest } from "../services/makeRequest";
import Modal from "./Modal";

const GroupMembers = ({ groupId }) => {
  const [groupMembers, setGroupMembers] = useState([]);
  const [error, setError] = useState(null);
  const [modalOpen, setModalOpen] = useState(false);

  useEffect(() => {
    console.log("groupMembers", `/groupmembers/${groupId}`);
    const loadMembers = async () => {
      try {
        const response = await makeRequest(`/groupmembers/${groupId}`);
        if (response !== null) {
          setGroupMembers(response);
        }
        console.log("groupmembers response", response);
      } catch (error) {
        setError(error.message);
      }
    };
    loadMembers();
  }, []);

  const groupMembersMap = groupMembers.map((member, index) => (
    <div key={index}>
      <p>User ID: {member.userId}</p>
      <p>Username: {member.userName}</p>
      <p>Image Path: {member.imagePath}</p>
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
          <p>
            This group with id: {groupId} has {groupMembers.length} members
          </p>
          <button onClick={openModal}>Show Members</button>
          <Modal open={modalOpen} onClose={closeModal}>
            {groupMembersMap}
          </Modal>
        </>
      ) : (
        <p>Apply to become a member of this group</p>
      )}
      ;
    </>
  );
};

export default GroupMembers;
