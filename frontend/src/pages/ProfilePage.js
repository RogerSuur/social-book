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
        .then((response) => setUser(response.data.user));
    };

    loadUser();
  }, []);

  return (
    <div className="content-area">
      <div className="row">
        <h1 className="column-title">{user.firstName}'s profile</h1>
      </div>
      <div className="row">
        <div className="column-title">First Name</div>
        <div className="column">{user.firstName}</div>
      </div>
      <div className="row">
        <div className="column-title">Last Name</div>
        <div className="column">{user.lastName}</div>
      </div>
      <div className="row">
        <div className="column-title">Email</div>
        <div className="column">{user.email}</div>
      </div>
      <div className="row">
        <div className="column-title">Birthday</div>
        <div className="column">{user.birthday}</div>
      </div>
      <div className="row">
        <div className="column-title">Nickname</div>
        <div className="column">{user.nickname}</div>
      </div>
      <div className="row">
        <div className="column-title">About Me</div>
        <div className="column">{user.about}</div>
      </div>
      <div className="row">
        <div className="column-title">Avatar</div>
        <div className="column"><img src={user.avatarImage}></img></div>
      </div>
      <div className="row">
        <div className="column-title">User Joined at</div>
        <div className="column">{user.createdAt}</div>
      </div>
      <div className="row">
        <div className="column-title">User Profile is public</div>
        <div className="column">{user.isPublic ? 'true' : 'false'}</div>
      </div>
    </div>
  );
};

export default Profile;
