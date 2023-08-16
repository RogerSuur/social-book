import List from "./List.js";
import { ListGroup } from "react-bootstrap";
import { LinkContainer } from "react-router-bootstrap";

const GenericGroupList = ({ url, loadNewGroups }) => {
  const mapGenericGroupList = (group, index) => (
    <LinkContainer action active={false} to={`/groups/${group.groupId}`}>
      <ListGroup.Item key={index}>
        <>{group.groupName}</>
      </ListGroup.Item>
    </LinkContainer>
  );

  return (
    <List
      url={url}
      mapFunction={mapGenericGroupList}
      loadNewGroups={loadNewGroups}
    />
  );
};

export default GenericGroupList;
