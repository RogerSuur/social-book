import { useState, useEffect } from "react";
import { makeRequest } from "../services/makeRequest";
import GenericModal from "../components/GenericModal";
import GenericUserList from "../components/GenericUserList";
import { GROUP_MEMBERS_URL } from "../utils/routes";
import { ListGroup } from "react-bootstrap";

const GroupMembers = ({ groupId }) => {
  const [groupMembers, setGroupMembers] = useState([]);
  const [errMsg, setErrMsg] = useState(null);

  useEffect(() => {
    const loadMembers = async () => {
      try {
        const response = await makeRequest(`/groupmembers/${groupId}`);
        if (response !== null) {
          setGroupMembers(response);
        }
      } catch (err) {
        setErrMsg(err);
      }
    };
    loadMembers();
  }, [groupId]);

  return (
    <GenericModal
      linkText={`${groupMembers?.length} members`}
      headerText={"Group members"}
      variant="flush"
    >
      <ListGroup>
        <GenericUserList variant="flush" url={GROUP_MEMBERS_URL + groupId} />
      </ListGroup>
    </GenericModal>
  );
};

export default GroupMembers;
