import { useState, useEffect } from "react";
import axios from "axios";
import ProfileSidebar from "../components/ProfileSidebar.js";
import ProfileContent from "../components/ProfileContent.js";

const Profile = () => {
  const [user, setUser] = useState({});
  const [selected, setSelected] = useState(null);

  const handleSelect = (item) => {
    setSelected(item);
  };

  useEffect(() => {
    const loadUser = async () => {
      await axios
        .get(`http://localhost:8000/profile`, {
          withCredentials: true,
        })
        .then((response) => {
          console.log(response);
          setUser(response.data);
        });
    };

    loadUser();
  }, []);

  console.log(user, "USER");

  return (
    <>
      <ProfileSidebar onSelect={handleSelect} />
      <ProfileContent selected={selected} />
    </>
  );
};

export default Profile;
