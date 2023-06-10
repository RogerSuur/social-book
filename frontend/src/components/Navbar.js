import { useEffect, useState } from "react";
import { Outlet, Link, useNavigate } from "react-router-dom";
import useAuth from "../hooks/useAuth";
import NotificationList from "../components/NotificationList.js";

const Navbar = () => {
  const { auth } = useAuth();
  const navigate = useNavigate();
  const [toggle, setToggle] = useState(false);

  useEffect(() => {
    navigate("/profile", { replace: true });
  }, []);

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
          <li onClick={handleToggle}>Notifications</li>
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
      {toggle && <NotificationList />}
      <Outlet />
    </>
  );
};

export default Navbar;
