import React, { useEffect } from "react";
import { Navigate, useLocation } from "react-router-dom";
import useAuth from "../hooks/useAuth";
import axios from "axios";

const LOGOUT_URL = "http://localhost:8000/logout";

const Logout = () => {
  const { setAuth } = useAuth();
  const location = useLocation();

  useEffect(() => {
    const authorisation = async () => {
      const response = await axios.get(LOGOUT_URL, {
        withCredentials: true,
      });

      console.log(response);

      setAuth(false);
    };

    authorisation();
  }, []);

  return <Navigate to="/signin" state={{ from: location }} replace />;
};

export default Logout;
