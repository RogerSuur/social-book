import React, { useState, useEffect } from "react";

const Users = ({ sender_id, data, users }) => {
  const [selected, setSelected] = useState(0);

  const passData = (event) => {
    data(event.target.value);
    setSelected(event.target.value);
  };

  useEffect(() => {
    sortUsers();
  }, [users]);

  const sortUsers = () => {
    users.sort((a, b) =>
      a.datetime < b.datetime ? 1 : b.datetime < a.datetime ? -1 : 0
    );
  };

  const styles = {
    fontWeight: selected === data ? "bold" : "normal",
  };

  const changeStyle = (id) => {
    if (id === selected) {
      return { border: "3px solid black" };
    }
    return { border: "none" };
  };

  return (
    <>
      <div style={{ flexBasis: "20%", border: "1px solid black" }}>
        <ul>
          {users.map((user) => (
            <div key={user.user_id}>
              {user.user_id !== sender_id && (
                <div style={changeStyle(user.user_id)}>
                  {user.online === true ? (
                    <span style={{ color: "green", fontWeight: "bold" }}>
                      Online{" "}
                    </span>
                  ) : (
                    <span style={{ color: "red" }}>Offline </span>
                  )}
                  <li
                    key={user.user_id}
                    onClick={passData}
                    value={user.user_id}
                  >
                    {user.username}
                  </li>
                </div>
              )}
            </div>
          ))}
        </ul>
      </div>
    </>
  );
};

export default Users;
