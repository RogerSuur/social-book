import { useEffect } from "react";
import { Outlet, Link, useNavigate } from "react-router-dom";
import useAuth from "../hooks/useAuth";

const Navbar = () => {
  const { auth } = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    navigate("/profile", { replace: true });
  }, []);

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
      <Outlet />
    </>
  );
};

export default Navbar;
