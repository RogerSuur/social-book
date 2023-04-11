import { useLocation, Navigate, Outlet } from "react-router-dom";
import { useEffect } from "react";
import useAuth from "../hooks/useAuth";
import axios from "axios";

const AUTH_URL = "http://localhost:8000/auth";

const RequireGuest = () => {
  const { auth, setAuth } = useAuth();
  const location = useLocation();

  console.log(location, "LOC GUEST");

  useEffect(() => {
    const authorisation = async () => {
      try {
        const response = await axios.get(AUTH_URL, {
          withCredentials: true,
        });

        console.log(JSON.stringify(response));
        setAuth(true);
      } catch (err) {
        if (!err?.response) {
          setAuth(false);
        }
      }
    };

    authorisation();
  }, [location]);

  return !auth ? (
    <Outlet />
  ) : (
    <Navigate to="/posts" state={{ from: location }} replace />
  );
};

export default RequireGuest;
