import React, { useState, useEffect } from "react";
import Select from "react-select";
import axios from "axios";
import { ADD_GROUP_MEMBERS_URL } from "../utils/routes";
import { Button, Stack } from "react-bootstrap";

const AddGroupMembers = ({ id }) => {
  const [open, setOpen] = useState(false);
  const [followers, setFollowers] = useState([]);
  const [formData, setFormData] = useState([]);

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

  const userOptions = followers.map((follower) => ({
    value: follower.id,
    label: `${follower.firstName} ${follower.lastName}`,
  }));

  const handleOpen = () => {
    setOpen(!open);
  };

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
    <Stack direction="horizontal">
      <div className="add-members">
        <Select options={userOptions} isMulti onChange={handleSelectChange} />
      </div>
      <Button onClick={handleSubmit}>Invite</Button>
    </Stack>
  );
};

export default AddGroupMembers;
