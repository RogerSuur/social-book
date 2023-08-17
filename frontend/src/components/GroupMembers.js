import { useState, useEffect } from "react";
import { makeRequest } from "../services/makeRequest";
import GenericModal from "../components/GenericModal";
import GenericUserList from "../components/GenericUserList";
import { Link } from "react-router-dom";
import ImageHandler from "../utils/imageHandler";
import { GROUP_MEMBERS_URL } from "../utils/routes";
import { ListGroup } from "react-bootstrap";

const GroupMembers = ({ groupId }) => {
  const [groupMembers, setGroupMembers] = useState([]);
  const [error, setError] = useState(null);

  useEffect(() => {
    const loadMembers = async () => {
      try {
        const response = await makeRequest(`/groupmembers/${groupId}`);
        if (response !== null) {
          setGroupMembers(response);
        }
      } catch (error) {
        setError(error);
      }
    };
    loadMembers();
  }, [groupId]);

  return (
    <GenericModal
      buttonText={`${groupMembers?.length} members`}
      headerText={"Group members"}
    >
      <ListGroup>
        <GenericUserList variant="flush" url={GROUP_MEMBERS_URL + groupId} />
      </ListGroup>
    </GenericModal>
  );
};

export default GroupMembers;
