import React, { useEffect, useState } from "react";
import { GROUPFEED_URL, GROUP_PAGE_URL } from "../utils/routes";
import axios from "axios";
import { useParams } from "react-router-dom";
import GroupMembers from "../components/GroupMembers";
import ImageHandler from "../utils/imageHandler";
import FeedPosts from "../components/FeedPosts.js";
import AvatarUpdater from "../components/AvatarUpdater";
import Events from "../components/Events";
import GroupRequestButton from "../components/GroupRequestButton.js";
import CreateGroupPosts from "../components/CreateGroupPosts.js";
import GenericModal from "../components/GenericModal";
import AddGroupMembers from "../components/AddGroupMembers";
import { Alert, Container, Stack, Col, Row, Badge } from "react-bootstrap";
import { PlusCircle } from "react-bootstrap-icons";

const GroupPage = () => {
  const [group, setGroup] = useState({});
  const { id } = useParams();
  const [errMsg, setErrMsg] = useState(null);
  const [reload, setReload] = useState(false);

  const handlePostUpdate = () => {
    setReload(!reload);
  };

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
      setErrMsg(err.message);
    }
  };

  const handleAvatarUpdate = () => {
    loadGroup();
  };

  useEffect(() => {
    loadGroup();
  }, [id]);

  const image = ImageHandler(group.imagePath, "defaultgroup.png", "group-img");

  return (
    <>
      {/* <div style={{ "max-width": `100px` }}>
        <Events groupId={+id} />
      </div> */}
      {errMsg ? (
        <Alert variant="danger" className="text-center">
          {errMsg}
        </Alert>
      ) : (
        <Container fluid>
          {image}
          <Stack direction="horizontal" gap="2">
            <Col xs="auto">
              <h1>{group.title}</h1>
            </Col>
            <Col xs="auto">
              {group?.isMember ? (
                <Stack direction="horizontal">
                  <div>
                    <GroupMembers groupId={+id} />
                  </div>
                  <div>
                    <GenericModal
                      img={<PlusCircle />}
                      variant="flush"
                      headerText="Add members"
                    >
                      <AddGroupMembers id={+id} />
                    </GenericModal>
                  </div>
                </Stack>
              ) : (
                <GroupRequestButton groupid={+id} />
              )}
            </Col>
          </Stack>

          <p>{group.description}</p>
          {group.isMember && (
            <>
              {/* <GenericModal buttonText="Upload new image">
                <AvatarUpdater
                  url={`${GROUP_PAGE_URL}${id}/avatar`}
                  onUploadSuccess={handleAvatarUpdate}
                />
              </GenericModal> */}
              <CreateGroupPosts groupId={id} onPostsUpdate={handlePostUpdate} />
              <FeedPosts url={GROUPFEED_URL + id} key={id} reload={reload} />
            </>
          )}
        </Container>
      )}
    </>
  );
};

export default GroupPage;
