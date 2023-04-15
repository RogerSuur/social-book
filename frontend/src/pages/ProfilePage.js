import { useState, useEffect } from "react";
import axios from "axios";

const Profile = () => {
  const [user, setUser] = useState({});

  useEffect(() => {
    const loadUser = async () => {
      await axios
        .get(`http://localhost:8000/profile`, {
          withCredentials: true,
        })
        .then((response) => setUser(response.data));
    };

    loadUser();
  }, []);

  return (
    <div className="content-area">
      <div className="row">
        <h1 className="column-title">{user.username}'s profile</h1>
      </div>
      <div className="row">
        <div className="column-title">Username</div>
        <div className="column">{user.username}</div>
      </div>
      <div className="row">
        <div className="column-title">First Name</div>
        <div className="column">{user.first_name}</div>
      </div>
      <div className="row">
        <div className="column-title">Last Name</div>
        <div className="column">{user.last_name}</div>
      </div>
      <div className="row">
        <div className="column-title">Sex</div>
        <div className="column">{user.sex}</div>
      </div>
      <div className="row">
        <div className="column-title">Age</div>
        <div className="column">{user.age}</div>
      </div>
      <div className="row">
        <div className="column-title">Email Address</div>
        <div className="column">{user.email}</div>
      </div>
    </div>
  );
};

export default Profile;
