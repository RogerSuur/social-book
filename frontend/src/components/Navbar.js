import { useEffect, useState } from "react";
import { Outlet, Link, useLocation } from "react-router-dom";
import useAuth from "../hooks/useAuth";
import NotificationList from "../components/NotificationList.js";

const Navbar = () => {
  const { auth } = useAuth();
  const location = useLocation();
  const [toggle, setToggle] = useState(false);

  useEffect(() => {
    setToggle(false);
  }, [location]);

  const handleToggle = () => {
    setToggle(!toggle);
  };

  return (
    <>
      <nav>
        <ul>
          {!auth && (
            <>
              <li>
                <Link className="text-link" to="/login">
                  Sign In
                </Link>
              </li>
              <li>
                <Link className="text-link" to="/signup">
                  Sign Up
                </Link>
              </li>
            </>
          )}
          {auth && <li onClick={handleToggle}>Notifications</li>}
          <li>
            <Link className="text-link" to="/profile">
              Profile
            </Link>
          </li>
          <li>
            <Link className="text-link" to="/posts">
              Posts
            </Link>
          </li>
          <li>
            <Link className="text-link" to="/chat">
              Inbox
            </Link>
          </li>
          <li>
            <Link className="text-link" id="logout" to="/logout">
              Logout
            </Link>
          </li>
        </ul>
      </nav>
      {auth && toggle && <NotificationList />}
      <Outlet />
    </>
  );
};

export default Navbar;
