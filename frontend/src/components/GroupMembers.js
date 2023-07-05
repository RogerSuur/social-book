import { useState, useEffect } from "react";
import { makeRequest } from "../services/makeRequest";
import Modal from "./Modal";

const GroupMembers = ({ groupId }) => {
  const [groupMembers, setGroupMembers] = useState([]);
  const [error, setError] = useState(null);
  const [modalOpen, setModalOpen] = useState(false);

  useEffect(() => {
    const loadMembers = async () => {
      try {
        const response = await makeRequest(`/groupmembers/${groupId}`);
        setGroupMembers(response);
      } catch (error) {
        setError(error.message);
      }
    };
    loadMembers();
  }, []);

  const openModal = () => {
    setModalOpen(true);
  };

  const closeModal = () => {
    setModalOpen(false);
  };

  return (
    <>
      <p>
        This group with id: {groupId} has {groupMembers.length} members
      </p>
      <button onClick={openModal}>Show Members</button>
      <Modal open={modalOpen} onClose={closeModal}>
        {groupMembers.map((member) => (
          <div key={member.userId}>
            <p>User ID: {member.userId}</p>
            <p>Username: {member.userName}</p>
            <p>Image Path: {member.imagePath}</p>
          </div>
        ))}
      </Modal>
    </>
  );
};

export default GroupMembers;
